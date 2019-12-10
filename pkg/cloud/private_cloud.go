package cloud

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/models"

	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/services"
)

type onPremCloud struct {
	Provider     string            `yaml:"provider"`
	Organization string            `yaml:"organization"`
	Project      string            `yaml:"project"`
	Instances    []privateInstance `yaml:"instances,omitempty"`
}

type privateInstance interface {
	addIPPair(priv, pub string)
	fill(ctx context.Context, cli *client.HTTP, instanceUUID string) error
}

type privInst struct {
	Name           string         `yaml:"name"`
	PublicIP       string         `yaml:"public_ip,omitempty"`
	PrivateIP      string         `yaml:"private_ip"`
	Interface      string         `yaml:"interface"`
	Provision      bool           `yaml:"provision"`
	Username       string         `yaml:"username,omitempty"`
	Password       string         `yaml:"password,omitempty"`
	Services       []string       `yaml:"services,omitempty"`
	Roles          []instanceRole `yaml:"roles,omitempty"`
	PrivateSubnets []string       `yaml:"private_subnet,omitempty"`
	ProtocolsMode  []string       `yaml:"protocols_mode,omitempty"`
	Gateway        string         `yaml:"gateway,omitempty"`
}

func (i *privInst) addIPPair(priv, pub string) {}

type tor struct {
	Name           string   `yaml:"name"`
	PublicIP       string   `yaml:"public_ip"`
	PrivateIP      string   `yaml:"private_ip"`
	PrivateSubnets []string `yaml:"private_subnet"`
	Roles          []string `yaml:"roles"`
	Provision      bool     `yaml:"provision"`
	Username       string   `yaml:"username"`
	Password       string   `yaml:"password"`
	Interface      []string `yaml:"interface"`
	AS             int      `yaml:"AS"`
	ProtocolsMode  []string `yaml:"protocols_mode"`
}

func (t *tor) addIPPair(priv, pub string) {}

func (c *onPremCloud) fill(ctx context.Context, cli *client.HTTP, cloudUUID string) error {
	cloudObj, err := cli.GetCloud(ctx, &services.GetCloudRequest{
		ID: cloudUUID,
	})
	if err != nil {
		return err
	}

	c.Provider = "onprem"
	c.Organization = cloudObj.Cloud.Organization
	//c.Project = cloudObj.Cloud.
	c.Instances = []privateInstance{}
	tags, err := getTags(ctx, cli, cloudUUID)
	if err != nil {
		return err
	}
	instances, err := getNodes(ctx, cli, tags)
	if err != nil {
		return err
	}
	for _, i := range instances {
		ii := &privInst{
			Name:      i.Hostname,
			PublicIP:  i.IPAddress,
			Provision: true,
		}
		if err = ii.fill(ctx, cli, i.UUID); err != nil {
			return err
		}
		c.Instances = append(c.Instances, ii)
	}
	tors, err := getTORs(ctx, cli, tags)
	if err != nil {
		return err
	}
	for _, t := range tors {
		tt := &tor{
			Name:          t.Name,
			PublicIP:      t.PhysicalRouterManagementIP,
			PrivateIP:     t.PhysicalRouterDataplaneIP,
			Roles:         []string{"tor"},
			Provision:     true,
			Username:      t.PhysicalRouterUserCredentials.Username,
			Password:      t.PhysicalRouterUserCredentials.Password,
			ProtocolsMode: []string{"bgp"},
		}
		if err = tt.fill(ctx, cli, t.UUID); err != nil {
			return err
		}
		c.Instances = append(c.Instances, tt)
	}

	return nil
}

func getTags(ctx context.Context, cli *client.HTTP, cloudUUID string) ([]*models.Tag, error) {
	cloudObj, err := cli.GetCloud(ctx, &services.GetCloudRequest{
		ID: cloudUUID,
	})
	if err != nil {
		return nil, err
	}
	provObj, err := cli.GetCloudProvider(ctx, &services.GetCloudProviderRequest{
		ID: cloudObj.Cloud.CloudProviders[0].UUID,
	})
	if err != nil {
		return nil, err
	}

	regionObj, err := cli.GetCloudRegion(ctx, &services.GetCloudRegionRequest{
		ID: provObj.CloudProvider.CloudRegions[0].UUID,
	})
	if err != nil {
		return nil, err
	}

	vCloudObj, err := cli.GetVirtualCloud(ctx, &services.GetVirtualCloudRequest{
		ID: regionObj.CloudRegion.VirtualClouds[0].UUID,
	})
	if err != nil {
		return nil, err
	}

	t := []*models.Tag{}
	for _, tagref := range vCloudObj.VirtualCloud.TagRefs {
		tagObj, err := cli.GetTag(ctx, &services.GetTagRequest{
			ID: tagref.UUID,
		})
		if err != nil {
			return nil, err
		}
		t = append(t, tagObj.Tag)
	}

	return t, nil
}

func getNodes(ctx context.Context, cli *client.HTTP, tags []*models.Tag) ([]*models.Node, error) {
	nn := []*models.Node{}
	for _, t := range tags {
		nn = append(nn, t.NodeBackRefs...)
	}

	for i, n := range nn {
		node, err := cli.GetNode(ctx, &services.GetNodeRequest{
			ID: n.UUID,
		})
		if err != nil {
			return nil, err
		}
		nn[i] = node.Node
	}

	return nn, nil
}

func getTORs(ctx context.Context, cli *client.HTTP, tags []*models.Tag) ([]*models.PhysicalRouter, error) {
	tors := []*models.PhysicalRouter{}
	for _, t := range tags {
		tors = append(tors, t.PhysicalRouterBackRefs...)
	}

	for i, t := range tors {
		tt, err := cli.GetPhysicalRouter(ctx, &services.GetPhysicalRouterRequest{
			ID: t.UUID,
		})
		if err != nil {
			return nil, err
		}
		tors[i] = tt.PhysicalRouter
	}
	return tors, nil
}

func (i *privInst) fill(ctx context.Context, cli *client.HTTP, instanceUUID string) error {
	ii, err := cli.GetNode(ctx, &services.GetNodeRequest{
		ID: instanceUUID,
	})
	if err != nil {
		return err
	}
	if len(ii.Node.Ports) == 0 {
		return errors.Errorf("onPrem node %s has no private interface", instanceUUID)
	}
	i.PrivateIP = ii.Node.Ports[0].IPAddress
	i.Interface = ii.Node.Ports[0].Name
	//TODO: fill SSH credentials
	i.Username = "centos"
	i.Password = "password"
	//TODO: fill services
	//TODO: fill protocols
	if err = i.fillGWDetails(ctx, cli, ii.Node.ContrailMulticloudGWNodeBackRefs); err != nil {
		return err
	}
	//TODO: fill subnets
	//TODO: fill gateway
	i.fillRoles(ii.Node)

	return nil
}

func (i *privInst) fillRoles(n *models.Node) {
	if n.ContrailVrouterNodeBackRefs != nil && n.KubernetesNodeBackRefs != nil {
		i.Roles = append(i.Roles, computeNodeInstanceRole)
	} else if n.ContrailVrouterNodeBackRefs != nil {
		i.Roles = append(i.Roles, vRouterInstanceRole)
	}
	if n.ContrailConfigNodeBackRefs != nil || n.ContrailControlNodeBackRefs != nil {
		i.Roles = append(i.Roles, controllerInstanceRole)
		i.Provision = false
	}
	if n.KubernetesMasterNodeBackRefs != nil {
		i.Roles = append(i.Roles, k8sMasterInstanceRole)
	}
	if n.ContrailMulticloudGWNodeBackRefs != nil {
		i.Roles = append(i.Roles, gatewayInstanceRole)
	}
	if n.OpenstackComputeNodeBackRefs != nil {
		i.Provision = false
	}
}

func (i *privInst) fillGWDetails(ctx context.Context, cli *client.HTTP, gwNodes []*models.ContrailMulticloudGWNode) error {
	for _, gwNodeRef := range gwNodes {
		gwNodeResp, err := cli.GetContrailMulticloudGWNode(ctx,
			&services.GetContrailMulticloudGWNodeRequest{
				ID: gwNodeRef.UUID,
			},
		)
		if err != nil {
			return err
		}

		i.ProtocolsMode = gwNodeResp.ContrailMulticloudGWNode.ProtocolsMode
		if gwNodeResp.ContrailMulticloudGWNode.Services != nil {
			for _, v := range gwNodeResp.ContrailMulticloudGWNode.Services {
				i.Services = append(i.Services, v)
			}
		}
	}
	return nil
}

func (t *tor) fill(ctx context.Context, cli *client.HTTP, cloudUUID string) error {
	t.AS = 6500
	t.Interface = []string{"irb.20", "irb.21"}
	t.PrivateSubnets = []string{"10.10.10.0/24", "30.10.10.0/24"}
	return nil
}
