// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To abort a deployment
//
// The following code aborts a running application version deployment for an environment
// named my-env:
func ExampleElasticBeanstalk_AbortEnvironmentUpdate_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.AbortEnvironmentUpdateInput{
		EnvironmentName: aws.String("my-env"),
	}

	result, err := svc.AbortEnvironmentUpdate(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To check the availability of a CNAME
//
// The following operation checks the availability of the subdomain my-cname:
func ExampleElasticBeanstalk_CheckDNSAvailability_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.CheckDNSAvailabilityInput{
		CNAMEPrefix: aws.String("my-cname"),
	}

	result, err := svc.CheckDNSAvailability(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a new application
//
// The following operation creates a new application named my-app:
func ExampleElasticBeanstalk_CreateApplication_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.CreateApplicationInput{
		ApplicationName: aws.String("my-app"),
		Description:     aws.String("my application"),
	}

	result, err := svc.CreateApplication(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeTooManyApplicationsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyApplicationsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a new application
//
// The following operation creates a new version (v1) of an application named my-app:
func ExampleElasticBeanstalk_CreateApplicationVersion_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.CreateApplicationVersionInput{
		ApplicationName:       aws.String("my-app"),
		AutoCreateApplication: aws.Bool(true),
		Description:           aws.String("my-app-v1"),
		Process:               aws.Bool(true),
		SourceBundle: &elasticbeanstalk.S3Location{
			S3Bucket: aws.String("my-bucket"),
			S3Key:    aws.String("sample.war"),
		},
		VersionLabel: aws.String("v1"),
	}

	result, err := svc.CreateApplicationVersion(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeTooManyApplicationsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyApplicationsException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyApplicationVersionsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyApplicationVersionsException, aerr.Error())
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeS3LocationNotInServiceRegionException:
				fmt.Println(elasticbeanstalk.ErrCodeS3LocationNotInServiceRegionException, aerr.Error())
			case elasticbeanstalk.ErrCodeCodeBuildNotInServiceRegionException:
				fmt.Println(elasticbeanstalk.ErrCodeCodeBuildNotInServiceRegionException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a configuration template
//
// The following operation creates a configuration template named my-app-v1 from the
// settings applied to an environment with the id e-rpqsewtp2j:
func ExampleElasticBeanstalk_CreateConfigurationTemplate_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.CreateConfigurationTemplateInput{
		ApplicationName: aws.String("my-app"),
		EnvironmentId:   aws.String("e-rpqsewtp2j"),
		TemplateName:    aws.String("my-app-v1"),
	}

	result, err := svc.CreateConfigurationTemplate(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyConfigurationTemplatesException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyConfigurationTemplatesException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a new environment for an application
//
// The following operation creates a new environment for version v1 of a java application
// named my-app:
func ExampleElasticBeanstalk_CreateEnvironment_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.CreateEnvironmentInput{
		ApplicationName:   aws.String("my-app"),
		CNAMEPrefix:       aws.String("my-app"),
		EnvironmentName:   aws.String("my-env"),
		SolutionStackName: aws.String("64bit Amazon Linux 2015.03 v2.0.0 running Tomcat 8 Java 8"),
		VersionLabel:      aws.String("v1"),
	}

	result, err := svc.CreateEnvironment(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeTooManyEnvironmentsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyEnvironmentsException, aerr.Error())
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a new environment for an application
//
// The following operation creates a new environment for version v1 of a java application
// named my-app:
func ExampleElasticBeanstalk_CreateStorageLocation_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.CreateStorageLocationInput{}

	result, err := svc.CreateStorageLocation(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
			case elasticbeanstalk.ErrCodeS3SubscriptionRequiredException:
				fmt.Println(elasticbeanstalk.ErrCodeS3SubscriptionRequiredException, aerr.Error())
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete an application
//
// The following operation deletes an application named my-app:
func ExampleElasticBeanstalk_DeleteApplication_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.DeleteApplicationInput{
		ApplicationName: aws.String("my-app"),
	}

	result, err := svc.DeleteApplication(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeOperationInProgressException:
				fmt.Println(elasticbeanstalk.ErrCodeOperationInProgressException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete an application version
//
// The following operation deletes an application version named 22a0-stage-150819_182129
// for an application named my-app:
func ExampleElasticBeanstalk_DeleteApplicationVersion_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.DeleteApplicationVersionInput{
		ApplicationName:    aws.String("my-app"),
		DeleteSourceBundle: aws.Bool(true),
		VersionLabel:       aws.String("22a0-stage-150819_182129"),
	}

	result, err := svc.DeleteApplicationVersion(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeSourceBundleDeletionException:
				fmt.Println(elasticbeanstalk.ErrCodeSourceBundleDeletionException, aerr.Error())
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeOperationInProgressException:
				fmt.Println(elasticbeanstalk.ErrCodeOperationInProgressException, aerr.Error())
			case elasticbeanstalk.ErrCodeS3LocationNotInServiceRegionException:
				fmt.Println(elasticbeanstalk.ErrCodeS3LocationNotInServiceRegionException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a configuration template
//
// The following operation deletes a configuration template named my-template for an
// application named my-app:
func ExampleElasticBeanstalk_DeleteConfigurationTemplate_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.DeleteConfigurationTemplateInput{
		ApplicationName: aws.String("my-app"),
		TemplateName:    aws.String("my-template"),
	}

	result, err := svc.DeleteConfigurationTemplate(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeOperationInProgressException:
				fmt.Println(elasticbeanstalk.ErrCodeOperationInProgressException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a draft configuration
//
// The following operation deletes a draft configuration for an environment named my-env:
func ExampleElasticBeanstalk_DeleteEnvironmentConfiguration_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.DeleteEnvironmentConfigurationInput{
		ApplicationName: aws.String("my-app"),
		EnvironmentName: aws.String("my-env"),
	}

	result, err := svc.DeleteEnvironmentConfiguration(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To view information about an application version
//
// The following operation retrieves information about an application version labeled
// v2:
func ExampleElasticBeanstalk_DescribeApplicationVersions_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.DescribeApplicationVersionsInput{
		ApplicationName: aws.String("my-app"),
		VersionLabels: []*string{
			aws.String("v2"),
		},
	}

	result, err := svc.DescribeApplicationVersions(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To view a list of applications
//
// The following operation retrieves information about applications in the current region:
func ExampleElasticBeanstalk_DescribeApplications_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.DescribeApplicationsInput{}

	result, err := svc.DescribeApplications(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To view configuration options for an environment
//
// The following operation retrieves descriptions of all available configuration options
// for an environment named my-env:
func ExampleElasticBeanstalk_DescribeConfigurationOptions_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.DescribeConfigurationOptionsInput{
		ApplicationName: aws.String("my-app"),
		EnvironmentName: aws.String("my-env"),
	}

	result, err := svc.DescribeConfigurationOptions(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To view configurations settings for an environment
//
// The following operation retrieves configuration settings for an environment named
// my-env:
func ExampleElasticBeanstalk_DescribeConfigurationSettings_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.DescribeConfigurationSettingsInput{
		ApplicationName: aws.String("my-app"),
		EnvironmentName: aws.String("my-env"),
	}

	result, err := svc.DescribeConfigurationSettings(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To view environment health
//
// The following operation retrieves overall health information for an environment named
// my-env:
func ExampleElasticBeanstalk_DescribeEnvironmentHealth_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.DescribeEnvironmentHealthInput{
		AttributeNames: []*string{
			aws.String("All"),
		},
		EnvironmentName: aws.String("my-env"),
	}

	result, err := svc.DescribeEnvironmentHealth(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInvalidRequestException:
				fmt.Println(elasticbeanstalk.ErrCodeInvalidRequestException, aerr.Error())
			case elasticbeanstalk.ErrCodeServiceException:
				fmt.Println(elasticbeanstalk.ErrCodeServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To view information about the AWS resources in your environment
//
// The following operation retrieves information about resources in an environment named
// my-env:
func ExampleElasticBeanstalk_DescribeEnvironmentResources_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.DescribeEnvironmentResourcesInput{
		EnvironmentName: aws.String("my-env"),
	}

	result, err := svc.DescribeEnvironmentResources(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To view information about an environment
//
// The following operation retrieves information about an environment named my-env:
func ExampleElasticBeanstalk_DescribeEnvironments_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.DescribeEnvironmentsInput{
		EnvironmentNames: []*string{
			aws.String("my-env"),
		},
	}

	result, err := svc.DescribeEnvironments(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To view events for an environment
//
// The following operation retrieves events for an environment named my-env:
func ExampleElasticBeanstalk_DescribeEvents_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.DescribeEventsInput{
		EnvironmentName: aws.String("my-env"),
	}

	result, err := svc.DescribeEvents(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To view environment health
//
// The following operation retrieves health information for instances in an environment
// named my-env:
func ExampleElasticBeanstalk_DescribeInstancesHealth_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.DescribeInstancesHealthInput{
		AttributeNames: []*string{
			aws.String("All"),
		},
		EnvironmentName: aws.String("my-env"),
	}

	result, err := svc.DescribeInstancesHealth(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInvalidRequestException:
				fmt.Println(elasticbeanstalk.ErrCodeInvalidRequestException, aerr.Error())
			case elasticbeanstalk.ErrCodeServiceException:
				fmt.Println(elasticbeanstalk.ErrCodeServiceException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To view solution stacks
//
// The following operation lists solution stacks for all currently available platform
// configurations and any that you have used in the past:
func ExampleElasticBeanstalk_ListAvailableSolutionStacks_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.ListAvailableSolutionStacksInput{}

	result, err := svc.ListAvailableSolutionStacks(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To rebuild an environment
//
// The following operation terminates and recreates the resources in an environment
// named my-env:
func ExampleElasticBeanstalk_RebuildEnvironment_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.RebuildEnvironmentInput{
		EnvironmentName: aws.String("my-env"),
	}

	result, err := svc.RebuildEnvironment(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To request tailed logs
//
// The following operation requests logs from an environment named my-env:
func ExampleElasticBeanstalk_RequestEnvironmentInfo_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.RequestEnvironmentInfoInput{
		EnvironmentName: aws.String("my-env"),
		InfoType:        aws.String("tail"),
	}

	result, err := svc.RequestEnvironmentInfo(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To restart application servers
//
// The following operation restarts application servers on all instances in an environment
// named my-env:
func ExampleElasticBeanstalk_RestartAppServer_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.RestartAppServerInput{
		EnvironmentName: aws.String("my-env"),
	}

	result, err := svc.RestartAppServer(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To retrieve tailed logs
//
// The following operation retrieves a link to logs from an environment named my-env:
func ExampleElasticBeanstalk_RetrieveEnvironmentInfo_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.RetrieveEnvironmentInfoInput{
		EnvironmentName: aws.String("my-env"),
		InfoType:        aws.String("tail"),
	}

	result, err := svc.RetrieveEnvironmentInfo(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To swap environment CNAMES
//
// The following operation swaps the assigned subdomains of two environments:
func ExampleElasticBeanstalk_SwapEnvironmentCNAMEs_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.SwapEnvironmentCNAMEsInput{
		DestinationEnvironmentName: aws.String("my-env-green"),
		SourceEnvironmentName:      aws.String("my-env-blue"),
	}

	result, err := svc.SwapEnvironmentCNAMEs(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To terminate an environment
//
// The following operation terminates an Elastic Beanstalk environment named my-env:
func ExampleElasticBeanstalk_TerminateEnvironment_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.TerminateEnvironmentInput{
		EnvironmentName: aws.String("my-env"),
	}

	result, err := svc.TerminateEnvironment(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To change an application's description
//
// The following operation updates the description of an application named my-app:
func ExampleElasticBeanstalk_UpdateApplication_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.UpdateApplicationInput{
		ApplicationName: aws.String("my-app"),
		Description:     aws.String("my Elastic Beanstalk application"),
	}

	result, err := svc.UpdateApplication(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To change an application version's description
//
// The following operation updates the description of an application version named 22a0-stage-150819_185942:
func ExampleElasticBeanstalk_UpdateApplicationVersion_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.UpdateApplicationVersionInput{
		ApplicationName: aws.String("my-app"),
		Description:     aws.String("new description"),
		VersionLabel:    aws.String("22a0-stage-150819_185942"),
	}

	result, err := svc.UpdateApplicationVersion(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To update a configuration template
//
// The following operation removes the configured CloudWatch custom health metrics configuration
// ConfigDocument from a saved configuration template named my-template:
func ExampleElasticBeanstalk_UpdateConfigurationTemplate_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.UpdateConfigurationTemplateInput{
		ApplicationName: aws.String("my-app"),
		OptionsToRemove: []*elasticbeanstalk.OptionSpecification{
			{
				Namespace:  aws.String("aws:elasticbeanstalk:healthreporting:system"),
				OptionName: aws.String("ConfigDocument"),
			},
		},
		TemplateName: aws.String("my-template"),
	}

	result, err := svc.UpdateConfigurationTemplate(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To update an environment to a new version
//
// The following operation updates an environment named "my-env" to version "v2" of
// the application to which it belongs:
func ExampleElasticBeanstalk_UpdateEnvironment_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.UpdateEnvironmentInput{
		EnvironmentName: aws.String("my-env"),
		VersionLabel:    aws.String("v2"),
	}

	result, err := svc.UpdateEnvironment(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To configure option settings
//
// The following operation configures several options in the aws:elb:loadbalancer namespace:
func ExampleElasticBeanstalk_UpdateEnvironment_shared01() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.UpdateEnvironmentInput{
		EnvironmentName: aws.String("my-env"),
		OptionSettings: []*elasticbeanstalk.ConfigurationOptionSetting{
			{
				Namespace:  aws.String("aws:elb:healthcheck"),
				OptionName: aws.String("Interval"),
				Value:      aws.String("15"),
			},
			{
				Namespace:  aws.String("aws:elb:healthcheck"),
				OptionName: aws.String("Timeout"),
				Value:      aws.String("8"),
			},
			{
				Namespace:  aws.String("aws:elb:healthcheck"),
				OptionName: aws.String("HealthyThreshold"),
				Value:      aws.String("2"),
			},
			{
				Namespace:  aws.String("aws:elb:healthcheck"),
				OptionName: aws.String("UnhealthyThreshold"),
				Value:      aws.String("3"),
			},
		},
	}

	result, err := svc.UpdateEnvironment(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To validate configuration settings
//
// The following operation validates a CloudWatch custom metrics config document:
func ExampleElasticBeanstalk_ValidateConfigurationSettings_shared00() {
	svc := elasticbeanstalk.New(session.New())
	input := &elasticbeanstalk.ValidateConfigurationSettingsInput{
		ApplicationName: aws.String("my-app"),
		EnvironmentName: aws.String("my-env"),
		OptionSettings: []*elasticbeanstalk.ConfigurationOptionSetting{
			{
				Namespace:  aws.String("aws:elasticbeanstalk:healthreporting:system"),
				OptionName: aws.String("ConfigDocument"),
				Value:      aws.String("{\"CloudWatchMetrics\": {\"Environment\": {\"ApplicationLatencyP99.9\": null,\"InstancesSevere\": 60,\"ApplicationLatencyP90\": 60,\"ApplicationLatencyP99\": null,\"ApplicationLatencyP95\": 60,\"InstancesUnknown\": 60,\"ApplicationLatencyP85\": 60,\"InstancesInfo\": null,\"ApplicationRequests2xx\": null,\"InstancesDegraded\": null,\"InstancesWarning\": 60,\"ApplicationLatencyP50\": 60,\"ApplicationRequestsTotal\": null,\"InstancesNoData\": null,\"InstancesPending\": 60,\"ApplicationLatencyP10\": null,\"ApplicationRequests5xx\": null,\"ApplicationLatencyP75\": null,\"InstancesOk\": 60,\"ApplicationRequests3xx\": null,\"ApplicationRequests4xx\": null},\"Instance\": {\"ApplicationLatencyP99.9\": null,\"ApplicationLatencyP90\": 60,\"ApplicationLatencyP99\": null,\"ApplicationLatencyP95\": null,\"ApplicationLatencyP85\": null,\"CPUUser\": 60,\"ApplicationRequests2xx\": null,\"CPUIdle\": null,\"ApplicationLatencyP50\": null,\"ApplicationRequestsTotal\": 60,\"RootFilesystemUtil\": null,\"LoadAverage1min\": null,\"CPUIrq\": null,\"CPUNice\": 60,\"CPUIowait\": 60,\"ApplicationLatencyP10\": null,\"LoadAverage5min\": null,\"ApplicationRequests5xx\": null,\"ApplicationLatencyP75\": 60,\"CPUSystem\": 60,\"ApplicationRequests3xx\": 60,\"ApplicationRequests4xx\": null,\"InstanceHealth\": null,\"CPUSoftirq\": 60}},\"Version\": 1}"),
			},
		},
	}

	result, err := svc.ValidateConfigurationSettings(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case elasticbeanstalk.ErrCodeInsufficientPrivilegesException:
				fmt.Println(elasticbeanstalk.ErrCodeInsufficientPrivilegesException, aerr.Error())
			case elasticbeanstalk.ErrCodeTooManyBucketsException:
				fmt.Println(elasticbeanstalk.ErrCodeTooManyBucketsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
