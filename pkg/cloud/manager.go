package cloud

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/auth"
	"github.com/Juniper/contrail/pkg/keystone"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/logutil/report"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

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

// Cloud represents cloud service.
type Cloud struct {
	config       *Config
	APIServer    *client.HTTP
	log          *logrus.Entry
	reporter     *report.Reporter
	streamServer *logutil.StreamServer
	ctx          context.Context
}

// NewCloudManager creates cloud fields by reading config from given configPath
func NewCloudManager(configPath string) (*Cloud, error) {
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.UnmarshalStrict(data, &c)
	if err != nil {
		return nil, err
	}

	return NewCloud(&c)
}

// NewCloud returns a new Cloud instance
func NewCloud(c *Config) (*Cloud, error) {
	if err := logutil.Configure(c.LogLevel); err != nil {
		return nil, err
	}

	s := client.NewHTTP(&client.HTTPConfig{
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
		var authKey interface{} = "auth"
		ctx = context.WithValue(
			context.Background(),
			authKey,
			auth.NewContext(
				c.DomainID,
				c.ProjectID,
				c.ID,
				[]string{c.ProjectName},
				"", auth.NewObjPerms(nil),
			),
		)
	}

	if c.CloudID != "" && c.Action == "" {
		return nil, fmt.Errorf("action not specified in the config")
	} else if c.CloudID == "" && c.Action != "" {
		return nil, fmt.Errorf("cloudID not specified in the config")
	}

	return &Cloud{
		APIServer: s,
		config:    c,
		log:       logutil.NewFileLogger("cloud", c.LogFile),
		reporter: report.NewReporter(
			s,
			fmt.Sprintf("%s/%s", defaultCloudResourcePath, c.CloudID),
			logutil.NewFileLogger("reporter", c.LogFile),
		),
		streamServer: logutil.NewStreamServer(c.LogFile),
		ctx:          ctx,
	}, nil
}

// Manage starts managing the cloud.
func (c *Cloud) Manage() error {
	c.streamServer.Serve()
	defer c.streamServer.Close()

	isDeleteReq, err := c.isCloudDeleteRequest()
	if err != nil {
		return err
	} else if isDeleteReq {
		if err = c.delete(); err != nil {
			return errors.Wrapf(err, "failed to delete cloud with CloudID %v", c.config.CloudID)
		}
		return nil
	}

	switch c.config.Action {
	case createAction:
		if err = c.create(); err != nil {
			return errors.Wrapf(err, "failed to create cloud with CloudID %v", c.config.CloudID)
		}
	case updateAction:
		if err = c.update(); err != nil {
			return errors.Wrapf(err, "failed to update cloud with CloudID %v", c.config.CloudID)
		}
	default:
		c.log.WithFields(logrus.Fields{
			"cloud-id": c.config.CloudID,
			"action":   c.config.Action,
		}).Info("Invalid action - ignoring")
	}
	return nil
}

func (c *Cloud) isCloudDeleteRequest() (bool, error) {
	cloudObj, err := GetCloud(c.ctx, c.APIServer, c.config.CloudID)
	if err != nil {
		return false, err
	}

	if c.config.Action == updateAction &&
		cloudObj.ProvisioningAction == deleteCloudAction &&
		cloudObj.ProvisioningState == statusNoState {
		return true, nil
	}
	return false, nil
}

func (c *Cloud) create() error {
	// Run pre-install steps
	topo, secret, data, err := c.initialize()
	if err != nil {
		c.reporter.ReportStatus(c.ctx, statusCreateFailed, defaultCloudResource)
		return err
	}

	if data.isCloudCreated() {
		return nil
	}

	c.log.Infof("Starting %s of cloud: %s", c.config.Action, data.info.FQName)

	c.reporter.ReportStatus(c.ctx, statusCreateProgress, defaultCloudResource)

	err = topo.createTopologyFile(GetTopoFile(c.config.CloudID))
	if err != nil {
		c.reporter.ReportStatus(c.ctx, statusCreateFailed, defaultCloudResource)
		return err
	}

	if data.isCloudPublic() {
		err = secret.createSecretFile()
		if err != nil {
			c.reporter.ReportStatus(c.ctx, statusCreateFailed, defaultCloudResource)
			return err
		}
		// depending upon the config action, it takes respective terraform action
		err = manageTerraform(c, c.config.Action)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, statusCreateFailed, defaultCloudResource)
			return err
		}
	}

	// update IP details only when cloud is public
	// basically when instances created by terraform
	if data.isCloudPublic() && (!c.config.Test) {
		err = updateIPDetails(c.ctx, c.config.CloudID, data)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, statusCreateFailed, defaultCloudResource)
			return err
		}
	}

	c.reporter.ReportStatus(c.ctx, statusCreated, defaultCloudResource)
	return nil
}

// nolint: gocyclo
func (c *Cloud) update() error {
	// Run pre-install steps
	topo, secret, data, err := c.initialize()
	if err != nil {
		c.reporter.ReportStatus(c.ctx, statusUpdateFailed, defaultCloudResource)
		return err
	}

	if !data.isCloudUpdateRequest() {
		var topoIsAlreadyUpdated bool
		topoIsAlreadyUpdated, err = topo.isUpdated(defaultCloudResource)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, statusUpdateFailed, defaultCloudResource)
			return err
		}
		if topoIsAlreadyUpdated {
			return nil
		}
	}

	c.log.Infof("Starting %s of cloud: %s", c.config.Action, data.info.FQName)
	c.reporter.ReportStatus(c.ctx, statusUpdateProgress, defaultCloudResource)

	err = topo.createTopologyFile(GetTopoFile(topo.cloud.config.CloudID))
	if err != nil {
		c.reporter.ReportStatus(c.ctx, statusUpdateFailed, defaultCloudResource)
		return err
	}

	//TODO(madhukar) handle if key-pair changes or aws-key

	if data.isCloudPublic() {
		err = secret.createSecretFile()
		if err != nil {
			c.reporter.ReportStatus(c.ctx, statusUpdateFailed, defaultCloudResource)
			return err
		}

		// depending upon the config action, it takes respective terraform action
		err = manageTerraform(c, c.config.Action)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, statusUpdateFailed, defaultCloudResource)
			return err
		}
	}

	//update IP address
	if data.isCloudPublic() && (!c.config.Test) {
		err = updateIPDetails(c.ctx, c.config.CloudID, data)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, statusUpdateFailed, defaultCloudResource)
			return err
		}
	}

	c.reporter.ReportStatus(c.ctx, statusUpdated, defaultCloudResource)
	return nil
}

func (c *Cloud) initialize() (*topology, *secret, *Data, error) {
	data, err := c.getCloudData(false)
	if err != nil {
		return nil, nil, nil, err
	}
	topo := newTopology(c, data)

	if data.isCloudPrivate() {
		return topo, nil, data, nil
	}

	// initialize secret struct
	secret, err := newSecret(c)
	if data.isCloudPublic() {
		if err != nil {
			return nil, nil, nil, err
		}
		err = secret.updateFileConfig(data)
		if err != nil {
			return nil, nil, nil, err
		}
	}

	return topo, secret, data, nil
}
func (c *Cloud) delete() error {
	// get cloud data
	data, err := c.getCloudData(true)
	if err != nil {
		return err
	}

	if data.isCloudPrivate() {
		err = c.verifyContrailClusterStatus(data)
		if err != nil {
			c.reporter.ReportStatus(c.ctx, statusUpdateFailed, defaultCloudResource)
			return err
		}
	}

	if data.isCloudPublic() {
		if tfStateOutputExists(c.config.CloudID) {
			err = manageTerraform(c, deleteAction)
			if err != nil {
				c.reporter.ReportStatus(c.ctx, statusUpdateFailed, defaultCloudResource)
				return err
			}
		}
	}

	// delete all the objects referred/in-tree of this cloud object
	err = c.deleteAPIObjects(data)
	if err != nil {
		c.reporter.ReportStatus(c.ctx, statusUpdateFailed, defaultCloudResource)
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
	if err != nil {
		return nil, err
	}

	return cloudData, nil
}

func (c *Cloud) newCloudData() (*Data, error) {
	data := Data{}
	data.cloud = c

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
		err := removePvtSubnetRefFromNodes(c.ctx, c.APIServer, d.getGatewayNodes())
		if err != nil {
			return err
		}
	}

	var errList, warnList []string

	retErrList := deleteContrailMCGWRole(c.ctx,
		c.APIServer, d.getGatewayNodes())

	if retErrList != nil {
		errList = append(errList, retErrList...)
	}

	if d.isCloudPublic() {
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

	if d.isCloudPublic() {
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
	templateRoot := c.config.TemplateRoot
	if templateRoot == "" {
		templateRoot = defaultTemplateRoot
	}
	return templateRoot
}
