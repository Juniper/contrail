package commander

import (
	"context"

	"github.com/Juniper/contrail/pkg/cluster"
	"github.com/Juniper/contrail/pkg/compilation/intent"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/models/basemodels"
	"github.com/Juniper/contrail/pkg/services"
)

// ContrailClusterIntent intent
type ContrailClusterIntent struct {
	intent.BaseIntent
	*models.ContrailCluster
}

// GetObject returns embedded resource object
func (i *ContrailClusterIntent) GetObject() basemodels.Object {
	return i.ContrailCluster
}

// NewContrailClusterIntent returns a new contrail cluster intent.
func NewContrailClusterIntent(
	_ context.Context,
	_ services.ReadService,
	request *services.CreateContrailClusterRequest,
) *ContrailClusterIntent {
	return &ContrailClusterIntent{
		ContrailCluster: request.GetContrailCluster(),
	}
}

// CreateContrailCluster evaluates ContrailCluster dependencies.
func (s *Service) CreateContrailCluster(
	ctx context.Context,
	request *services.CreateContrailClusterRequest,
) (*services.CreateContrailClusterResponse, error) {

	i := NewContrailClusterIntent(ctx, s.ReadService, request)

	err := s.storeAndEvaluate(ctx, i)
	if err != nil {
		return nil, err
	}

	config := getClusterConfig()
	config.Action = cluster.ActionCreate
	config.ClusterID = cc.UUID
	clusterManager, err := cluster.NewCluster(config)
	if err != nil {
		return nil, err
	}
	err = clusterManager.Manage()
	if err != nil {
		return nil, err
	}

	return s.BaseService.CreateContrailCluster(ctx, request)
}

// UpdateContrailCluster evaluates ContrailCluster dependencies.
func (s *Service) UpdateContrailCluster(
	ctx context.Context,
	request *services.UpdateContrailClusterRequest,
) (*services.UpdateContrailClusterResponse, error) {

	cc := request.GetContrailCluster()
	i := LoadContrailClusterIntent(s.cache, cc.GetUUID())

	err := s.storeAndEvaluate(ctx, i)
	if err != nil {
		return nil, err
	}

	config := getClusterConfig()
	config.Action = cluster.ActionUpdate
	config.ClusterID = cc.UUID
	clusterManager, err := cluster.NewCluster(config)
	if err != nil {
		return nil, err
	}
	err = clusterManager.Manage()
	if err != nil {
		return nil, err
	}

	return s.BaseService.UpdateContrailCluster(ctx, request)
}

// LoadContrailClusterIntent loads a contrail cluster intent from cache.
func LoadContrailClusterIntent(
	c intent.Loader,
	uuid string,
) *ContrailClusterIntent {
	i := c.Load(models.KindContrailCluster, intent.ByUUID(uuid))
	actual, _ := i.(*ContrailClusterIntent) //nolint: errcheck
	return actual
}

func getClusterConfig() *cluster.Config {
	config := &cluster.Config{
		Endpoint:    viper.GetString("client.endpoint"),
		ID:          viper.GetString("client.id"),
		Password:    viper.GetString("client.password"),
		ProjectID:   viper.GetString("client.project_id"),
		ProjectName: viper.GetString("client.project_name"),
		DomainID:    viper.GetString("client.domain_id"),
		DomainName:  viper.GetString("client.domain_name"),

		AuthURL:  viper.GetString("keystone.authurl"),
		Insecure: viper.GetBool("insecure"),

		LogLevel:     viper.GetString("commander.log_level"),
		LogFile:      viper.GetString("commander.log_file"),
		TemplateRoot: viper.GetString("commander.template_root"),
	}

	return config
}
