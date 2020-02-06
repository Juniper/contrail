package cloud

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Juniper/asf/pkg/auth"
	"github.com/Juniper/asf/pkg/keystone"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/Juniper/asf/pkg/osutil"
	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	asfclient "github.com/Juniper/asf/pkg/client"
	yaml "gopkg.in/yaml.v2"
)

// Config represents cloud configuration needed by cloudManager
type Config struct { // nolint: maligned
	// ID of cloud
	ID string `yaml:"id"`
	// Password of Cluster account.
	Password string `yaml:"password"`
	// DomainID is ID of keystone domain used for authentication.
	DomainID string `yaml:"domain_id"`
	// ProjectID is ID of keystone project used for authentication.
	ProjectID string `yaml:"project_id"`
	// DomainName is Name of keystone domain used for authentication.
	DomainName string `yaml:"domain_name"`
	// ProjectName is Name of keystone project used for authentication.
	ProjectName string `yaml:"project_name"`
	// AuthURL defines authentication URL.
	AuthURL string `yaml:"auth_url"`
	// Endpoint of API Server.
	Endpoint string `yaml:"endpoint"`
	// InSecure https connection to endpoint
	InSecure bool `yaml:"insecure"`
	// UUID of cloud to be managed.
	CloudID string `yaml:"cloud_id,omitempty"`
	// Action to the performed with the cloud (values: create, update, delete).
	Action string `yaml:"cloud_action,omitempty"`
	// Logging level
	LogLevel string `yaml:"log_level"`
	// Logging  file
	LogFile string `yaml:"log_file"`
	// Template root directory
	TemplateRoot string `yaml:"template_root"`

	// Optional Test var to run cloud in test mode
	Test bool `yaml:"test"`
}

type terraformStateReader interface {
	Read() (TerraformState, error)
}

// CommandExecutor interface provides methods to execute a command
type CommandExecutor interface {
	ExecCmdAndWait(r *report.Reporter, cmd string, args []string, dir string, envVars ...string) error
}

// Cloud represents cloud service.
type Cloud struct {
	config               *Config
	APIServer            *client.HTTP
	commandExecutor      CommandExecutor
	log                  *logrus.Entry
	reporter             *report.Reporter
	streamServer         *logutil.StreamServer
	terraformStateReader terraformStateReader
	// TODO(dfurman): Do not store Contexts inside a struct type; instead, pass a Context explicitly to each function
	// that needs it. See: https://golang.org/pkg/context/
	ctx context.Context
}

// NewCloudManager creates cloud fields by reading config from given configPath
func NewCloudManager(configPath string, commandExecutor CommandExecutor) (*Cloud, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.UnmarshalStrict(data, &c)
	if err != nil {
		return nil, err
	}

	return NewCloud(&c, cloudTfStateReader{c.CloudID}, commandExecutor)
}

// NewCloud returns a new Cloud instance
func NewCloud(c *Config, terraformStateReader terraformStateReader, e CommandExecutor) (*Cloud, error) {
	if err := logutil.Configure(c.LogLevel); err != nil {
		return nil, err
	}

	s := client.NewHTTP(&asfclient.HTTPConfig{
		ID:       c.ID,
		Password: c.Password,
		Endpoint: c.Endpoint,
		AuthURL:  c.AuthURL,
		Scope: keystone.NewScope(
			c.DomainID,
			c.DomainName,
			c.ProjectID,
			c.ProjectName,
		),
		Insecure: c.InSecure,
	})

	ctx := auth.NoAuth(context.Background())
	if c.AuthURL != "" {
		ctx = auth.WithIdentity(
			context.Background(),
			keystone.NewAuthIdentity(c.DomainID, c.ProjectID, c.ID, []string{c.ProjectName}),
		)
	}

	if c.Action == "" {
		return nil, fmt.Errorf("action not specified in the config")
	}
	if c.CloudID == "" {
		return nil, fmt.Errorf("cloudID not specified in the config")
	}

	return &Cloud{
		APIServer: s,
		config:    c,
		log:       logutil.NewFileLogger("cloud", c.LogFile).WithField("cloudID", c.CloudID),
		reporter: report.NewReporter(
			s,
			fmt.Sprintf("%s/%s", defaultCloudResourcePath, c.CloudID),
			logutil.NewFileLogger("reporter", c.LogFile).WithField("cloudID", c.CloudID),
		),
		streamServer:         logutil.NewStreamServer(c.LogFile),
		commandExecutor:      e,
		terraformStateReader: terraformStateReader,
		ctx:                  ctx,
	}, nil
}

// Manage starts managing the cloud.
func (c *Cloud) Manage() error {
	if err := c.manage(); err != nil {
		if setErr := c.fail(); setErr != nil {
			c.log.Errorf("Could not set Cloud state to failed state: %v", setErr)
		}
		c.log.Errorf("Cloud operation failed: %v", err)
		return err
	}
	return nil
}

func (c *Cloud) manage() (err error) {
	var isSet bool
	if isSet, err = c.provisioningSetToNonState(); err != nil {
		return errors.Wrap(err, "failed to resolve state of Cloud")
	} else if !isSet {
		return nil
	}
	data, err := c.getCloudData(false)
	if err != nil {
		return errors.Wrap(err, "failed to get Cloud data")
	}
	defer func() {
		if delErr := c.removeVulnerableFiles(data); delErr != nil {
			err = errors.Errorf(
				"failed to delete vulnerable files: %s; manage error (if any): %s", delErr, err,
			)
		}
	}()
	return c.HandleCloudRequest()
}

func (c *Cloud) provisioningSetToNonState() (bool, error) {
	cloudObject, err := GetCloud(c.ctx, c.APIServer, c.config.CloudID)
	if err != nil {
		return false, err
	}
	return cloudObject.GetProvisioningState() == StatusNoState, nil
}

// HandleCloudRequest handles Cloud create/update/delete request.
func (c *Cloud) HandleCloudRequest() error {
	c.streamServer.Serve()
	defer c.streamServer.Close()

	isDeleteReq, err := c.isCloudDeleteRequest()
	if err != nil {
		return err
	} else if isDeleteReq {
		c.log.Debug("Starting delete of cloud")
		if err = c.delete(); err != nil {
			return errors.Wrapf(err, "failed to delete cloud with CloudID %v", c.config.CloudID)
		}
		c.log.Debug("Cloud deleted")
		return nil
	}

	switch c.config.Action {
	case CreateAction:
		c.log.Debug("Starting create of cloud")
		c.reporter.ReportStatus(c.ctx, map[string]interface{}{StatusField: StatusCreateProgress}, defaultCloudResource)
		if err = c.create(); err != nil {
			return errors.Wrapf(err, "failed to create cloud with CloudID %v", c.config.CloudID)
		}
		c.reporter.ReportStatus(c.ctx, map[string]interface{}{StatusField: StatusCreated}, defaultCloudResource)
	case UpdateAction:
		c.log.Debug("Starting update of cloud")
		c.reporter.ReportStatus(c.ctx, map[string]interface{}{StatusField: StatusUpdateProgress}, defaultCloudResource)
		if err = c.update(); err != nil {
			return errors.Wrapf(err, "failed to update cloud with CloudID %v", c.config.CloudID)
		}
		c.reporter.ReportStatus(c.ctx, map[string]interface{}{StatusField: StatusUpdated}, defaultCloudResource)
	default:
		return errors.Errorf("Invalid action %s called for cloud %s", c.config.Action, c.config.CloudID)
	}
	return nil
}

func (c *Cloud) isCloudDeleteRequest() (bool, error) {
	cloudObj, err := GetCloud(c.ctx, c.APIServer, c.config.CloudID)
	if err != nil {
		return false, err
	}

	return c.config.Action == UpdateAction && cloudObj.ProvisioningAction == DeleteCloudAction, nil
}

func (c *Cloud) fail() error {
	cloudObject, err := GetCloud(c.ctx, c.APIServer, c.config.CloudID)
	if err != nil {
		return err
	}
	failStatus := c.getFailStatusField(cloudObject.GetProvisioningState())
	if failStatus == "" {
		return errors.New("unknown state change. Trying to set failure from state: " +
			cloudObject.GetProvisioningState())
	}
	c.reporter.ReportStatus(c.ctx, map[string]interface{}{StatusField: failStatus}, defaultCloudResource)
	return nil
}

func (c *Cloud) getFailStatusField(currentStatus string) string {
	switch currentStatus {
	case StatusCreateProgress:
		return StatusCreateFailed
	case StatusUpdateProgress:
		return StatusUpdateFailed
	case StatusNoState:
		switch c.config.Action {
		case CreateAction:
			return StatusCreateFailed
		case UpdateAction:
			return StatusUpdateFailed
		}
	}
	return ""
}

// nolint: gocyclo
func (c *Cloud) create() error {
	data, err := c.getCloudData(false)
	if err != nil {
		return err
	}

	pubData := &publicCloud{}
	onPremData := newOnPremCloud(c.ctx, c.APIServer, c.log)
	if !data.isCloudPrivate() {
		if err = pubData.Fill(c.ctx, c.APIServer, c.config.CloudID); err != nil {
			return err
		}
	} else {
		if err = onPremData.Fill(c.config.CloudID); err != nil {
			return err
		}
	}

	secret, err := c.initialize(data)
	if err != nil {
		return err
	}

	// Performing create // TODO(Daniel): extract function
	c.log.Infof("Starting %s of cloud: %s", c.config.Action, data.info.FQName)

	if data.isCloudPrivate() {
		err = onPremData.MarshalAndSave(GetTopoFile(c.config.CloudID))
	} else {
		err = pubData.MarshalAndSave(GetTopoFile(c.config.CloudID))
	}
	if err != nil {
		return err
	}

	c.log.Infof("Created topology file for cloud with uuid: %s ", c.config.CloudID)

	if !data.isCloudPrivate() {
		if err = secret.createSecretFile(data.info.GetParentClusterUUID()); err != nil {
			return err
		}
		// depending upon the config action, it takes respective terraform action
		if err = updateTopology(c, data.modifiedProviders()); err != nil {
			return err
		}
	}

	// update IP details only when cloud is public
	// basically when instances created by terraform
	if err = c.updatePublicCloudIP(data); err != nil {
		return err
	}

	if err = c.resetPublicCloudModifiedStatus(data); err != nil {
		return err
	}
	c.log.Debugf("Created cloud %s", c.config.CloudID)
	return nil
}

func (d *Data) modifiedProviders() []string {
	s := []string{}
	if d.info.AwsModified {
		s = append(s, aws)
	}
	if d.info.AzureModified {
		s = append(s, azure)
	}
	if d.info.GCPModified {
		s = append(s, google)
	}
	return s
}

func (c *Cloud) resetPublicCloudModifiedStatus(data *Data) error {
	if data.isCloudPrivate() {
		return nil
	}
	_, err := c.APIServer.UpdateCloud(c.ctx, &services.UpdateCloudRequest{
		Cloud: &models.Cloud{
			UUID: c.config.CloudID,
		},
		FieldMask: types.FieldMask{
			Paths: []string{
				models.CloudFieldAwsModified,
				models.CloudFieldAzureModified,
				models.CloudFieldGCPModified,
			},
		},
	})
	return err
}

// nolint: gocyclo
func (c *Cloud) update() error {
	data, err := c.getCloudData(false)
	if err != nil {
		return err
	}

	if !data.isCloudPrivate() && len(data.modifiedProviders()) == 0 {
		c.log.Debug("Topology is already up to date - skipping update")
		return nil
	}

	pubData := &publicCloud{}
	onPremData := newOnPremCloud(c.ctx, c.APIServer, c.log)
	if !data.isCloudPrivate() {
		if err = pubData.Fill(c.ctx, c.APIServer, c.config.CloudID); err != nil {
			return err
		}
	} else {
		if err = onPremData.Fill(c.config.CloudID); err != nil {
			return err
		}
	}

	var secret *secret
	if !data.isCloudPrivate() {
		secret, err = c.initializeSecret(data)
		if err != nil {
			return err
		}
	}

	// Performing update // TODO(Daniel): extract function
	c.log.Infof("Starting %s of cloud: %s", c.config.Action, data.info.FQName)

	if data.isCloudPrivate() {
		err = onPremData.MarshalAndSave(GetTopoFile(c.config.CloudID))
	} else {
		err = pubData.MarshalAndSave(GetTopoFile(c.config.CloudID))
	}
	if err != nil {
		return err
	}

	c.log.Infof("Created topology file for cloud with uuid: %s ", c.config.CloudID)

	//TODO(madhukar) handle if key-pair changes or aws-key
	if !data.isCloudPrivate() {
		if err = secret.createSecretFile(data.info.GetParentClusterUUID()); err != nil {
			return err
		}

		if err = updateTopology(c, data.modifiedProviders()); err != nil {
			return err
		}
	}

	//update IP address
	if err = c.updatePublicCloudIP(data); err != nil {
		return err
	}

	if err = c.resetPublicCloudModifiedStatus(data); err != nil {
		return err
	}

	c.log.Debugf("Updated cloud %s", c.config.CloudID)
	return nil
}

func (c *Cloud) updatePublicCloudIP(data *Data) error {
	if !data.isCloudPrivate() {
		tfState, err := c.terraformStateReader.Read()
		if err != nil {
			return err
		}
		if err = updateIPDetails(c.ctx, data.instances, tfState); err != nil {
			return err
		}
	}
	return nil
}

func (c *Cloud) initialize(d *Data) (*secret, error) {
	var secret *secret
	var err error
	if !d.isCloudPrivate() {
		secret, err = c.initializeSecret(d)
		if err != nil {
			return nil, err
		}
	}

	return secret, nil
}

func (c *Cloud) initializeSecret(d *Data) (*secret, error) {
	s := newSecret(c)
	if !d.isCloudPrivate() {
		kp, err := s.getKeypair(d)
		if err != nil {
			return nil, err
		}
		if err = s.sfc.Update(kp); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func (c *Cloud) delete() error {
	data, err := c.getCloudData(true)
	if err != nil {
		return err
	}

	if data.isCloudPrivate() {
		if err = c.verifyContrailClusterStatus(data); err != nil {
			return err
		}
	} else {
		var secret *secret
		secret, err = c.initializeSecret(data)
		if err != nil {
			return err
		}
		if err = secret.createSecretFile(data.info.GetParentClusterUUID()); err != nil {
			return err
		}
		if err = destroyTopology(c, data.modifiedProviders()); err != nil {
			return err
		}
	}

	if err = c.deleteAPIObjects(data); err != nil {
		return err
	}

	return os.RemoveAll(GetCloudDir(c.config.CloudID))
}

func (c *Cloud) getCloudData(isDelRequest bool) (*Data, error) {
	cloudData, err := c.newCloudData()
	if err != nil {
		return nil, err
	}

	err = cloudData.update(isDelRequest)
	return cloudData, err
}

func (c *Cloud) newCloudData() (*Data, error) {
	data := Data{cloud: c}

	cloudObject, err := GetCloud(c.ctx, c.APIServer, c.config.CloudID)
	if err != nil {
		return nil, err
	}

	data.info = cloudObject
	return &data, nil
}

// nolint: gocyclo
func (c *Cloud) deleteAPIObjects(d *Data) error {
	if d.isCloudPrivate() {
		if err := removePvtSubnetRefFromNodes(c.ctx, c.APIServer, d.getGatewayNodes()); err != nil {
			return err
		}
	}

	var errList, warnList []string

	retErrList := deleteContrailMCGWRole(c.ctx,
		c.APIServer, d.getGatewayNodes())

	if retErrList != nil {
		errList = append(errList, retErrList...)
	}

	if !d.isCloudPrivate() {
		retErrList = deleteNodeObjects(c.ctx, c.APIServer, d.instances)
		if retErrList != nil {
			errList = append(errList, retErrList...)
		}
	}

	retErrList = deleteCloudProviderAndDeps(c.ctx,
		c.APIServer, d.providers)
	if retErrList != nil {
		errList = append(errList, retErrList...)
	}

	_, err := c.APIServer.DeleteCloud(c.ctx,
		&services.DeleteCloudRequest{
			ID: d.info.UUID,
		},
	)
	if err != nil {
		errList = append(errList, fmt.Sprintf(
			"failed deleting Cloud %s err_msg: %s",
			d.info.UUID, err))
	}

	cloudUserErrList := deleteCloudUsers(c.ctx, c.APIServer, d.users)
	if cloudUserErrList != nil {
		warnList = append(warnList, cloudUserErrList...)
	}

	if !d.isCloudPrivate() {
		credErrList := deleteCredentialAndDeps(c.ctx, c.APIServer, d.credentials)
		warnList = append(warnList, credErrList...)
	}

	// log the warning messages
	if len(warnList) > 0 {
		c.log.Warnf("could not delete cloud refs deps because of errors: %s",
			strings.Join(warnList, "\n"))
	}
	// join all the errors and return it
	if len(errList) > 0 {
		return errors.New(strings.Join(errList, "\n"))
	}
	return nil
}

func (c *Cloud) verifyContrailClusterStatus(data *Data) error {
	for _, clusterRef := range data.info.ContrailClusterBackRefs {
		err := waitForClusterStatusToBeUpdated(c.ctx, c.log,
			c.APIServer, clusterRef.UUID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Cloud) getTemplateRoot() string {
	if c.config.TemplateRoot == "" {
		return DefaultTemplateRoot
	}
	return c.config.TemplateRoot
}

func (c *Cloud) removeVulnerableFiles(data *Data) error {
	if data.isCloudPrivate() {
		return nil
	}

	kfd := services.NewKeyFileDefaults()

	f := []string{
		GetTerraformAWSPlanFile(c.config.CloudID),
		GetTerraformAzurePlanFile(c.config.CloudID),
		GetTerraformGCPPlanFile(c.config.CloudID),
		GetSecretFile(c.config.CloudID),
		kfd.GetAzureSubscriptionIDPath(),
		kfd.GetAzureClientIDPath(),
		kfd.GetAzureClientSecretPath(),
		kfd.GetAzureTenantIDPath(),
	}
	// Deploy worker to provision multicloud needs AWS secret files. Cloud worker must not delete them.
	// Deploy worker needs to remove those files.
	if !data.info.IsMulticloudProvisioning {
		f = append(
			f,
			kfd.GetAWSAccessPath(),
			kfd.GetAWSSecretPath(),
			kfd.GetGoogleAccountPath(),
		)
	}
	return osutil.ForceRemoveFiles(f, c.log)
}

// Config returns copy of Cloud's config.
func (c *Cloud) Config() Config {
	return *c.config
}

// Context returns Cloud's context.
func (c *Cloud) Context() context.Context {
	return c.ctx
}
