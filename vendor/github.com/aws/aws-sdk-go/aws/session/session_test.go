package session

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/awstesting"
	"github.com/aws/aws-sdk-go/internal/shareddefaults"
	"github.com/aws/aws-sdk-go/service/s3"
)

func TestNewDefaultSession(t *testing.T) {
	oldEnv := initSessionTestEnv()
	defer awstesting.PopEnv(oldEnv)

	s := New(&aws.Config{Region: aws.String("region")})

	assert.Equal(t, "region", *s.Config.Region)
	assert.Equal(t, http.DefaultClient, s.Config.HTTPClient)
	assert.NotNil(t, s.Config.Logger)
	assert.Equal(t, aws.LogOff, *s.Config.LogLevel)
}

func TestNew_WithCustomCreds(t *testing.T) {
	oldEnv := initSessionTestEnv()
	defer awstesting.PopEnv(oldEnv)

	customCreds := credentials.NewStaticCredentials("AKID", "SECRET", "TOKEN")
	s := New(&aws.Config{Credentials: customCreds})

	assert.Equal(t, customCreds, s.Config.Credentials)
}

type mockLogger struct {
	*bytes.Buffer
}

func (w mockLogger) Log(args ...interface{}) {
	fmt.Fprintln(w, args...)
}

func TestNew_WithSessionLoadError(t *testing.T) {
	oldEnv := initSessionTestEnv()
	defer awstesting.PopEnv(oldEnv)

	os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
	os.Setenv("AWS_CONFIG_FILE", testConfigFilename)
	os.Setenv("AWS_PROFILE", "assume_role_invalid_source_profile")

	logger := bytes.Buffer{}
	s := New(&aws.Config{Logger: &mockLogger{&logger}})

	assert.NotNil(t, s)

	svc := s3.New(s)
	_, err := svc.ListBuckets(&s3.ListBucketsInput{})

	assert.Error(t, err)
	assert.Contains(t, logger.String(), "ERROR: failed to create session with AWS_SDK_LOAD_CONFIG enabled")
	assert.Contains(t, err.Error(), SharedConfigAssumeRoleError{
		RoleARN: "assume_role_invalid_source_profile_role_arn",
	}.Error())
}

func TestSessionCopy(t *testing.T) {
	oldEnv := initSessionTestEnv()
	defer awstesting.PopEnv(oldEnv)

	os.Setenv("AWS_REGION", "orig_region")

	s := Session{
		Config:   defaults.Config(),
		Handlers: defaults.Handlers(),
	}

	newSess := s.Copy(&aws.Config{Region: aws.String("new_region")})

	assert.Equal(t, "orig_region", *s.Config.Region)
	assert.Equal(t, "new_region", *newSess.Config.Region)
}

func TestSessionClientConfig(t *testing.T) {
	s, err := NewSession(&aws.Config{
		Credentials: credentials.AnonymousCredentials,
		Region:      aws.String("orig_region"),
		EndpointResolver: endpoints.ResolverFunc(
			func(service, region string, opts ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
				if e, a := "mock-service", service; e != a {
					t.Errorf("expect %q service, got %q", e, a)
				}
				if e, a := "other-region", region; e != a {
					t.Errorf("expect %q region, got %q", e, a)
				}
				return endpoints.ResolvedEndpoint{
					URL:           "https://" + service + "." + region + ".amazonaws.com",
					SigningRegion: region,
				}, nil
			},
		),
	})
	assert.NoError(t, err)

	cfg := s.ClientConfig("mock-service", &aws.Config{Region: aws.String("other-region")})

	assert.Equal(t, "https://mock-service.other-region.amazonaws.com", cfg.Endpoint)
	assert.Equal(t, "other-region", cfg.SigningRegion)
	assert.Equal(t, "other-region", *cfg.Config.Region)
}

func TestNewSession_NoCredentials(t *testing.T) {
	oldEnv := initSessionTestEnv()
	defer awstesting.PopEnv(oldEnv)

	s, err := NewSession()
	assert.NoError(t, err)

	assert.NotNil(t, s.Config.Credentials)
	assert.NotEqual(t, credentials.AnonymousCredentials, s.Config.Credentials)
}

func TestNewSessionWithOptions_OverrideProfile(t *testing.T) {
	oldEnv := initSessionTestEnv()
	defer awstesting.PopEnv(oldEnv)

	os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testConfigFilename)
	os.Setenv("AWS_PROFILE", "other_profile")

	s, err := NewSessionWithOptions(Options{
		Profile: "full_profile",
	})
	assert.NoError(t, err)

	assert.Equal(t, "full_profile_region", *s.Config.Region)

	creds, err := s.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "full_profile_akid", creds.AccessKeyID)
	assert.Equal(t, "full_profile_secret", creds.SecretAccessKey)
	assert.Empty(t, creds.SessionToken)
	assert.Contains(t, creds.ProviderName, "SharedConfigCredentials")
}

func TestNewSessionWithOptions_OverrideSharedConfigEnable(t *testing.T) {
	oldEnv := initSessionTestEnv()
	defer awstesting.PopEnv(oldEnv)

	os.Setenv("AWS_SDK_LOAD_CONFIG", "0")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testConfigFilename)
	os.Setenv("AWS_PROFILE", "full_profile")

	s, err := NewSessionWithOptions(Options{
		SharedConfigState: SharedConfigEnable,
	})
	assert.NoError(t, err)

	assert.Equal(t, "full_profile_region", *s.Config.Region)

	creds, err := s.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "full_profile_akid", creds.AccessKeyID)
	assert.Equal(t, "full_profile_secret", creds.SecretAccessKey)
	assert.Empty(t, creds.SessionToken)
	assert.Contains(t, creds.ProviderName, "SharedConfigCredentials")
}

func TestNewSessionWithOptions_OverrideSharedConfigDisable(t *testing.T) {
	oldEnv := initSessionTestEnv()
	defer awstesting.PopEnv(oldEnv)

	os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testConfigFilename)
	os.Setenv("AWS_PROFILE", "full_profile")

	s, err := NewSessionWithOptions(Options{
		SharedConfigState: SharedConfigDisable,
	})
	assert.NoError(t, err)

	assert.Empty(t, *s.Config.Region)

	creds, err := s.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "full_profile_akid", creds.AccessKeyID)
	assert.Equal(t, "full_profile_secret", creds.SecretAccessKey)
	assert.Empty(t, creds.SessionToken)
	assert.Contains(t, creds.ProviderName, "SharedConfigCredentials")
}

func TestNewSessionWithOptions_OverrideSharedConfigFiles(t *testing.T) {
	oldEnv := initSessionTestEnv()
	defer awstesting.PopEnv(oldEnv)

	os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testConfigFilename)
	os.Setenv("AWS_PROFILE", "config_file_load_order")

	s, err := NewSessionWithOptions(Options{
		SharedConfigFiles: []string{testConfigOtherFilename},
	})
	assert.NoError(t, err)

	assert.Equal(t, "shared_config_other_region", *s.Config.Region)

	creds, err := s.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "shared_config_other_akid", creds.AccessKeyID)
	assert.Equal(t, "shared_config_other_secret", creds.SecretAccessKey)
	assert.Empty(t, creds.SessionToken)
	assert.Contains(t, creds.ProviderName, "SharedConfigCredentials")
}

func TestNewSessionWithOptions_Overrides(t *testing.T) {
	cases := []struct {
		InEnvs    map[string]string
		InProfile string
		OutRegion string
		OutCreds  credentials.Value
	}{
		{
			InEnvs: map[string]string{
				"AWS_SDK_LOAD_CONFIG":         "0",
				"AWS_SHARED_CREDENTIALS_FILE": testConfigFilename,
				"AWS_PROFILE":                 "other_profile",
			},
			InProfile: "full_profile",
			OutRegion: "full_profile_region",
			OutCreds: credentials.Value{
				AccessKeyID:     "full_profile_akid",
				SecretAccessKey: "full_profile_secret",
				ProviderName:    "SharedConfigCredentials",
			},
		},
		{
			InEnvs: map[string]string{
				"AWS_SDK_LOAD_CONFIG":         "0",
				"AWS_SHARED_CREDENTIALS_FILE": testConfigFilename,
				"AWS_REGION":                  "env_region",
				"AWS_ACCESS_KEY":              "env_akid",
				"AWS_SECRET_ACCESS_KEY":       "env_secret",
				"AWS_PROFILE":                 "other_profile",
			},
			InProfile: "full_profile",
			OutRegion: "env_region",
			OutCreds: credentials.Value{
				AccessKeyID:     "env_akid",
				SecretAccessKey: "env_secret",
				ProviderName:    "EnvConfigCredentials",
			},
		},
		{
			InEnvs: map[string]string{
				"AWS_SDK_LOAD_CONFIG":         "0",
				"AWS_SHARED_CREDENTIALS_FILE": testConfigFilename,
				"AWS_CONFIG_FILE":             testConfigOtherFilename,
				"AWS_PROFILE":                 "shared_profile",
			},
			InProfile: "config_file_load_order",
			OutRegion: "shared_config_region",
			OutCreds: credentials.Value{
				AccessKeyID:     "shared_config_akid",
				SecretAccessKey: "shared_config_secret",
				ProviderName:    "SharedConfigCredentials",
			},
		},
	}

	for _, c := range cases {
		oldEnv := initSessionTestEnv()
		defer awstesting.PopEnv(oldEnv)

		for k, v := range c.InEnvs {
			os.Setenv(k, v)
		}

		s, err := NewSessionWithOptions(Options{
			Profile:           c.InProfile,
			SharedConfigState: SharedConfigEnable,
		})
		assert.NoError(t, err)

		creds, err := s.Config.Credentials.Get()
		assert.NoError(t, err)
		assert.Equal(t, c.OutRegion, *s.Config.Region)
		assert.Equal(t, c.OutCreds.AccessKeyID, creds.AccessKeyID)
		assert.Equal(t, c.OutCreds.SecretAccessKey, creds.SecretAccessKey)
		assert.Equal(t, c.OutCreds.SessionToken, creds.SessionToken)
		assert.Contains(t, creds.ProviderName, c.OutCreds.ProviderName)
	}
}

const assumeRoleRespMsg = `
<AssumeRoleResponse xmlns="https://sts.amazonaws.com/doc/2011-06-15/">
  <AssumeRoleResult>
    <AssumedRoleUser>
      <Arn>arn:aws:sts::account_id:assumed-role/role/session_name</Arn>
      <AssumedRoleId>AKID:session_name</AssumedRoleId>
    </AssumedRoleUser>
    <Credentials>
      <AccessKeyId>AKID</AccessKeyId>
      <SecretAccessKey>SECRET</SecretAccessKey>
      <SessionToken>SESSION_TOKEN</SessionToken>
      <Expiration>%s</Expiration>
    </Credentials>
  </AssumeRoleResult>
  <ResponseMetadata>
    <RequestId>request-id</RequestId>
  </ResponseMetadata>
</AssumeRoleResponse>
`

func TestSesisonAssumeRole(t *testing.T) {
	oldEnv := initSessionTestEnv()
	defer awstesting.PopEnv(oldEnv)

	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testConfigFilename)
	os.Setenv("AWS_PROFILE", "assume_role_w_creds")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf(assumeRoleRespMsg, time.Now().Add(15*time.Minute).Format("2006-01-02T15:04:05Z"))))
	}))

	s, err := NewSession(&aws.Config{Endpoint: aws.String(server.URL), DisableSSL: aws.Bool(true)})

	creds, err := s.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "AKID", creds.AccessKeyID)
	assert.Equal(t, "SECRET", creds.SecretAccessKey)
	assert.Equal(t, "SESSION_TOKEN", creds.SessionToken)
	assert.Contains(t, creds.ProviderName, "AssumeRoleProvider")
}

func TestSessionAssumeRole_WithMFA(t *testing.T) {
	oldEnv := initSessionTestEnv()
	defer awstesting.PopEnv(oldEnv)

	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testConfigFilename)
	os.Setenv("AWS_PROFILE", "assume_role_w_creds")

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.FormValue("SerialNumber"), "0123456789")
		assert.Equal(t, r.FormValue("TokenCode"), "tokencode")

		w.Write([]byte(fmt.Sprintf(assumeRoleRespMsg, time.Now().Add(15*time.Minute).Format("2006-01-02T15:04:05Z"))))
	}))

	customProviderCalled := false
	sess, err := NewSessionWithOptions(Options{
		Profile: "assume_role_w_mfa",
		Config: aws.Config{
			Region:     aws.String("us-east-1"),
			Endpoint:   aws.String(server.URL),
			DisableSSL: aws.Bool(true),
		},
		SharedConfigState: SharedConfigEnable,
		AssumeRoleTokenProvider: func() (string, error) {
			customProviderCalled = true

			return "tokencode", nil
		},
	})
	assert.NoError(t, err)

	creds, err := sess.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.True(t, customProviderCalled)

	assert.Equal(t, "AKID", creds.AccessKeyID)
	assert.Equal(t, "SECRET", creds.SecretAccessKey)
	assert.Equal(t, "SESSION_TOKEN", creds.SessionToken)
	assert.Contains(t, creds.ProviderName, "AssumeRoleProvider")
}

func TestSessionAssumeRole_WithMFA_NoTokenProvider(t *testing.T) {
	oldEnv := initSessionTestEnv()
	defer awstesting.PopEnv(oldEnv)

	os.Setenv("AWS_REGION", "us-east-1")
	os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testConfigFilename)
	os.Setenv("AWS_PROFILE", "assume_role_w_creds")

	_, err := NewSessionWithOptions(Options{
		Profile:           "assume_role_w_mfa",
		SharedConfigState: SharedConfigEnable,
	})
	assert.Equal(t, err, AssumeRoleTokenProviderNotSetError{})
}

func TestSessionAssumeRole_DisableSharedConfig(t *testing.T) {
	// Backwards compatibility with Shared config disabled
	// assume role should not be built into the config.
	oldEnv := initSessionTestEnv()
	defer awstesting.PopEnv(oldEnv)

	os.Setenv("AWS_SDK_LOAD_CONFIG", "0")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testConfigFilename)
	os.Setenv("AWS_PROFILE", "assume_role_w_creds")

	s, err := NewSession()
	assert.NoError(t, err)

	creds, err := s.Config.Credentials.Get()
	assert.NoError(t, err)
	assert.Equal(t, "assume_role_w_creds_akid", creds.AccessKeyID)
	assert.Equal(t, "assume_role_w_creds_secret", creds.SecretAccessKey)
	assert.Contains(t, creds.ProviderName, "SharedConfigCredentials")
}

func TestSessionAssumeRole_InvalidSourceProfile(t *testing.T) {
	// Backwards compatibility with Shared config disabled
	// assume role should not be built into the config.
	oldEnv := initSessionTestEnv()
	defer awstesting.PopEnv(oldEnv)

	os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", testConfigFilename)
	os.Setenv("AWS_PROFILE", "assume_role_invalid_source_profile")

	s, err := NewSession()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "SharedConfigAssumeRoleError: failed to load assume role")
	assert.Nil(t, s)
}

func TestSharedConfigCredentialSource(t *testing.T) {
	cases := []struct {
		name              string
		profile           string
		expectedError     error
		expectedAccessKey string
		expectedSecretKey string
		init              func(*aws.Config, string) func() error
	}{
		{
			name:              "env var credential source",
			profile:           "env_var_credential_source",
			expectedAccessKey: "access_key",
			expectedSecretKey: "secret_key",
			init: func(cfg *aws.Config, profile string) func() error {
				os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
				os.Setenv("AWS_CONFIG_FILE", "testdata/credential_source_config")
				os.Setenv("AWS_PROFILE", profile)
				os.Setenv("AWS_ACCESS_KEY", "access_key")
				os.Setenv("AWS_SECRET_KEY", "secret_key")

				return func() error {
					os.Unsetenv("AWS_SDK_LOAD_CONFIG")
					os.Unsetenv("AWS_CONFIG_FILE")
					os.Unsetenv("AWS_PROFILE")
					os.Unsetenv("AWS_ACCESS_KEY")
					os.Unsetenv("AWS_SECRET_KEY")

					return nil
				}
			},
		},
		{
			name:          "credential source and source profile",
			profile:       "invalid_source_and_credential_source",
			expectedError: ErrSharedConfigSourceCollision,
			init: func(cfg *aws.Config, profile string) func() error {
				os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
				os.Setenv("AWS_CONFIG_FILE", "testdata/credential_source_config")
				os.Setenv("AWS_PROFILE", profile)
				os.Setenv("AWS_ACCESS_KEY", "access_key")
				os.Setenv("AWS_SECRET_KEY", "secret_key")

				return func() error {
					os.Unsetenv("AWS_SDK_LOAD_CONFIG")
					os.Unsetenv("AWS_CONFIG_FILE")
					os.Unsetenv("AWS_PROFILE")
					os.Unsetenv("AWS_ACCESS_KEY")
					os.Unsetenv("AWS_SECRET_KEY")

					return nil
				}
			},
		},
		{
			name:              "ec2metadata credential source",
			profile:           "ec2metadata",
			expectedAccessKey: "AKID",
			expectedSecretKey: "SECRET",
			init: func(cfg *aws.Config, profile string) func() error {
				os.Setenv("AWS_REGION", "us-east-1")
				os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
				os.Setenv("AWS_CONFIG_FILE", "testdata/credential_source_config")
				os.Setenv("AWS_PROFILE", "ec2metadata")

				const ec2MetadataResponse = `{
	  "Code": "Success",
	  "Type": "AWS-HMAC",
	  "AccessKeyId" : "access-key",
	  "SecretAccessKey" : "secret-key",
	  "Token" : "token",
	  "Expiration" : "2100-01-01T00:00:00Z",
	  "LastUpdated" : "2009-11-23T0:00:00Z"
	}`

				ec2MetadataCalled := false
				ec2MetadataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Path == "/meta-data/iam/security-credentials/RoleName" {
						ec2MetadataCalled = true
						w.Write([]byte(ec2MetadataResponse))
					} else if r.URL.Path == "/meta-data/iam/security-credentials/" {
						w.Write([]byte("RoleName"))
					} else {
						w.Write([]byte(""))
					}
				}))

				stsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte(fmt.Sprintf(assumeRoleRespMsg, time.Now().Add(15*time.Minute).Format("2006-01-02T15:04:05Z"))))
				}))

				cfg.EndpointResolver = endpoints.ResolverFunc(
					func(service, region string, opts ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
						if service == "ec2metadata" {
							return endpoints.ResolvedEndpoint{
								URL: ec2MetadataServer.URL,
							}, nil
						}

						return endpoints.ResolvedEndpoint{
							URL: stsServer.URL,
						}, nil
					},
				)

				return func() error {
					os.Unsetenv("AWS_SDK_LOAD_CONFIG")
					os.Unsetenv("AWS_CONFIG_FILE")
					os.Unsetenv("AWS_PROFILE")
					os.Unsetenv("AWS_REGION")

					ec2MetadataServer.Close()
					stsServer.Close()

					if !ec2MetadataCalled {
						return fmt.Errorf("expected ec2metadata to be called")
					}

					return nil
				}
			},
		},
		{
			name:              "ecs container credential source",
			profile:           "ecscontainer",
			expectedAccessKey: "access-key",
			expectedSecretKey: "secret-key",
			init: func(cfg *aws.Config, profile string) func() error {
				os.Setenv("AWS_REGION", "us-east-1")
				os.Setenv("AWS_SDK_LOAD_CONFIG", "1")
				os.Setenv("AWS_CONFIG_FILE", "testdata/credential_source_config")
				os.Setenv("AWS_PROFILE", "ecscontainer")
				os.Setenv("AWS_CONTAINER_CREDENTIALS_RELATIVE_URI", "/ECS")

				const ecsResponse = `{
	  "Code": "Success",
	  "Type": "AWS-HMAC",
	  "AccessKeyId" : "access-key",
	  "SecretAccessKey" : "secret-key",
	  "Token" : "token",
	  "Expiration" : "2100-01-01T00:00:00Z",
	  "LastUpdated" : "2009-11-23T0:00:00Z"
	}`

				ecsCredsCalled := false
				ecsMetadataServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Path == "/ECS" {
						ecsCredsCalled = true
						w.Write([]byte(ecsResponse))
					} else {
						w.Write([]byte(""))
					}
				}))

				shareddefaults.ECSContainerCredentialsURI = ecsMetadataServer.URL

				stsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte(fmt.Sprintf(assumeRoleRespMsg, time.Now().Add(15*time.Minute).Format("2006-01-02T15:04:05Z"))))
				}))

				cfg.Endpoint = aws.String(stsServer.URL)

				cfg.EndpointResolver = endpoints.ResolverFunc(
					func(service, region string, opts ...func(*endpoints.Options)) (endpoints.ResolvedEndpoint, error) {
						fmt.Println("SERVICE", service)
						return endpoints.ResolvedEndpoint{
							URL: stsServer.URL,
						}, nil
					},
				)

				return func() error {
					os.Unsetenv("AWS_SDK_LOAD_CONFIG")
					os.Unsetenv("AWS_CONFIG_FILE")
					os.Unsetenv("AWS_PROFILE")
					os.Unsetenv("AWS_REGION")
					os.Unsetenv("AWS_CONTAINER_CREDENTIALS_RELATIVE_URI")

					ecsMetadataServer.Close()
					stsServer.Close()

					if !ecsCredsCalled {
						return fmt.Errorf("expected ec2metadata to be called")
					}

					return nil
				}
			},
		},
	}

	for _, c := range cases {
		cfg := &aws.Config{}
		clean := c.init(cfg, c.profile)
		sess, err := NewSession(cfg)
		if e, a := c.expectedError, err; e != a {
			t.Errorf("expected %v, but received %v", e, a)
		}

		if c.expectedError != nil {
			continue
		}

		creds, err := sess.Config.Credentials.Get()
		if err != nil {
			t.Errorf("expected no error, but received %v", err)
		}

		if e, a := c.expectedAccessKey, creds.AccessKeyID; e != a {
			t.Errorf("expected %v, but received %v", e, a)
		}

		if e, a := c.expectedSecretKey, creds.SecretAccessKey; e != a {
			t.Errorf("expected %v, but received %v", e, a)
		}

		if err := clean(); err != nil {
			t.Errorf("expected no error, but received %v", err)
		}
	}
}

func initSessionTestEnv() (oldEnv []string) {
	oldEnv = awstesting.StashEnv()
	os.Setenv("AWS_CONFIG_FILE", "file_not_exists")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", "file_not_exists")

	return oldEnv
}
