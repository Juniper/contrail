package cluster

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Juniper/asf/pkg/keystone"
	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/Juniper/contrail/pkg/deploy/base"
)

const (
	defaultTemplateRoot = "./pkg/cluster/configs"
)

type deployCluster struct {
	base.Deploy
	cluster     *Cluster
	clusterID   string
	action      string
	clusterData *base.Data
}

func newDeployCluster(c *Cluster, cData *base.Data, moduleName string) *deployCluster {
	return &deployCluster{
		cluster:     c,
		clusterID:   c.config.ClusterID,
		action:      c.config.Action,
		clusterData: cData,
		Deploy: base.Deploy{
			Reporter: newReporter(c),
			Log:      logutil.NewFileLogger(moduleName, c.config.LogFile),
		},
	}
}

func newReporter(cluster *Cluster) *report.Reporter {
	return report.NewReporter(
		cluster.APIServer,
		fmt.Sprintf("%s/%s", defaultResourcePath, cluster.config.ClusterID),
		logutil.NewFileLogger("reporter", cluster.config.LogFile),
	)
}

func (p *deployCluster) isCreated() bool {
	state := p.clusterData.ClusterInfo.ProvisioningState
	if p.action == "create" && (state == statusNoState || state == "") {
		return false
	}
	p.Log.Infof("Cluster %s already deployed, STATE: %s", p.clusterID, state)
	return true
}

func (p *deployCluster) getTemplateRoot() string {
	templateRoot := p.cluster.config.TemplateRoot
	if templateRoot == "" {
		templateRoot = defaultTemplateRoot
	}
	return templateRoot
}

func (p *deployCluster) getWorkRoot() string {
	workRoot := p.cluster.config.WorkRoot
	if workRoot == "" {
		workRoot = defaultWorkRoot
	}
	return workRoot
}

func (p *deployCluster) getClusterHomeDir() string {
	return filepath.Join(p.getWorkRoot(), p.clusterID)
}

func (p *deployCluster) getWorkingDir() string {
	return filepath.Join(p.getClusterHomeDir())
}

func (p *deployCluster) createWorkingDir() error {
	return os.MkdirAll(p.getWorkingDir(), os.ModePerm)
}

func (p *deployCluster) deleteWorkingDir() error {
	return os.RemoveAll(p.getClusterHomeDir())
}

func (p *deployCluster) ensureServiceUserCreated() error {
	if p.clusterData.ClusterInfo.Orchestrator != "openstack" {
		return nil
	}
	ctx := context.Background()
	name, pass := p.clusterData.KeystoneAdminCredential()

	token, err := p.cluster.APIServer.Keystone.ObtainToken(
		ctx, name, pass, keystone.NewScope("default", "", "", keystone.AdminRoleName),
	)
	if err != nil {
		return err
	}
	ctx = keystone.WithXAuthToken(ctx, token)

	_, err = p.cluster.APIServer.Keystone.EnsureServiceUserCreated(ctx, keystone.User{
		Name:     p.cluster.config.ServiceUserID,
		Password: p.cluster.config.ServiceUserPassword,
	})
	return err
}

func (p *deployCluster) createEndpoints() error {
	e := &base.EndpointData{
		ClusterID:   p.clusterID,
		ResManager:  base.NewResourceManager(p.cluster.APIServer, p.cluster.config.LogFile),
		ClusterData: p.clusterData,
		Log:         p.Log,
	}

	return e.Create()
}

func (p *deployCluster) updateEndpoints() error {
	e := &base.EndpointData{
		ClusterID:   p.clusterID,
		ResManager:  base.NewResourceManager(p.cluster.APIServer, p.cluster.config.LogFile),
		ClusterData: p.clusterData,
		Log:         p.Log,
	}

	return e.Update()
}

func (p *deployCluster) deleteEndpoints() error {
	e := &base.EndpointData{
		ClusterID:  p.clusterID,
		ResManager: base.NewResourceManager(p.cluster.APIServer, p.cluster.config.LogFile),
		Log:        p.Log,
	}

	return e.Remove()
}
