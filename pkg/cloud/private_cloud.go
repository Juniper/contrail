package cloud

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/models"

	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/services"
)

type onPremCloud struct {
	Provider  string        `yaml:"provider"`
	Instances []interface{} `yaml:"instances,omitempty"`
}

type privateInstance struct {
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

func (c *onPremCloud) fill(ctx context.Context, cli *client.HTTP, log *logrus.Entry, cloudUUID string) error {
	c.Provider = "onprem"
	c.Instances = []interface{}{}

	tags, err := getTags(ctx, cli, log, cloudUUID)
	if err != nil {
		return err
	}

	if err = c.createInstances(ctx, cli, log, cloudUUID, tags); err != nil {
		return err
	}

	if err = c.createTors(ctx, cli, log, cloudUUID, tags); err != nil {
		return err
	}

	return nil
}

func validateOnPremResourceUniqueness(resourceName string, count int, log *logrus.Entry) error {
	if count == 0 {
		return errors.Errorf("onPrem cloud has no %s", resourceName)
	}
	if count > 1 {
		log.Warnf("onPrem cloud has wrong number of %s: %d, only first will be considered", resourceName, count)
	}
	return nil
}

func getTags(ctx context.Context, cli *client.HTTP, log *logrus.Entry, cloudUUID string) ([]*models.Tag, error) {
	cloudObj, err := cli.GetCloud(ctx, &services.GetCloudRequest{
		ID: cloudUUID,
	})
	if err != nil {
		return nil, err
	}

	if err = validateOnPremResourceUniqueness("CloudProviders", len(cloudObj.Cloud.CloudProviders), log); err != nil {
		return nil, err
	}

	provObj, err := cli.GetCloudProvider(ctx, &services.GetCloudProviderRequest{
		ID: cloudObj.Cloud.CloudProviders[0].UUID,
	})
	if err != nil {
		return nil, err
	}

	if err = validateOnPremResourceUniqueness("CloudRegions", len(provObj.CloudProvider.CloudRegions), log); err != nil {
		return nil, err
	}

	regionObj, err := cli.GetCloudRegion(ctx, &services.GetCloudRegionRequest{
		ID: provObj.CloudProvider.CloudRegions[0].UUID,
	})
	if err != nil {
		return nil, err
	}

	if err = validateOnPremResourceUniqueness("VirtualClouds", len(regionObj.CloudRegion.VirtualClouds), log); err != nil {
		return nil, err
	}

	vCloudObj, err := cli.GetVirtualCloud(ctx, &services.GetVirtualCloudRequest{
		ID: regionObj.CloudRegion.VirtualClouds[0].UUID,
	})
	if err != nil {
		return nil, err
	}

	tags := []*models.Tag{}
	for _, tagref := range vCloudObj.VirtualCloud.TagRefs {
		tagObj, err := cli.GetTag(ctx, &services.GetTagRequest{
			ID: tagref.UUID,
		})
		if err != nil {
			return nil, err
		}
		tags = append(tags, tagObj.Tag)
	}

	return tags, nil
}

func (c *onPremCloud) createInstances(
	ctx context.Context, cli *client.HTTP, log *logrus.Entry, cloudUUID string, tags []*models.Tag,
) error {
	nodes, err := getNodes(ctx, cli, tags)
	if err != nil {
		return err
	}

	for _, i := range nodes {
		instance := &privateInstance{
			Name:      i.Hostname,
			PublicIP:  i.IPAddress,
			Provision: true,
		}
		if err = instance.fill(ctx, cli, log, i.UUID); err != nil {
			return err
		}
		c.Instances = append(c.Instances, instance)
	}
	return nil
}

func getNodes(ctx context.Context, cli *client.HTTP, tags []*models.Tag) ([]*models.Node, error) {
	nodes := []*models.Node{}
	for _, t := range tags {
		nodes = append(nodes, t.NodeBackRefs...)
	}

	for i, n := range nodes {
		node, err := cli.GetNode(ctx, &services.GetNodeRequest{
			ID: n.UUID,
		})
		if err != nil {
			return nil, err
		}
		nodes[i] = node.Node
	}

	return nodes, nil
}

func (i *privateInstance) fill(ctx context.Context, cli *client.HTTP, log *logrus.Entry, instanceUUID string) error {
	ii, err := cli.GetNode(ctx, &services.GetNodeRequest{
		ID: instanceUUID,
	})
	if err != nil {
		return err
	}

	if err = i.fillPortDetails(log, ii.Node.Ports); err != nil {
		return err
	}
	if err = i.fillCredentialDetails(ctx, cli, ii.Node.CredentialRefs, log); err != nil {
		return err
	}
	if err = i.fillGWDetails(
		ctx, cli, log, ii.Node.ContrailMulticloudGWNodeBackRefs, ii.Node.CloudPrivateSubnetRefs); err != nil {
		return err
	}
	if err = i.fillVrouterDetails(ctx, cli, log, ii.Node.ContrailVrouterNodeBackRefs); err != nil {
		return err
	}
	i.fillRoles(ii.Node)

	return nil
}

func (i *privateInstance) fillPortDetails(log *logrus.Entry, ports []*models.Port) error {
	if len(ports) == 0 {
		return errors.New("onPrem node has no private interface")
	}
	if len(ports) > 1 {
		log.Warnf("onPrem node has invalid number of ports: %d, only first will be considered", len(ports))
	}

	i.PrivateIP = ports[0].IPAddress
	i.Interface = ports[0].Name
	return nil
}

func (i *privateInstance) fillCredentialDetails(
	ctx context.Context, cli *client.HTTP, credentialRefs []*models.NodeCredentialRef, log *logrus.Entry,
) error {
	if len(credentialRefs) == 0 {
		return errors.New("onPrem node has no credentials")
	}
	if len(credentialRefs) > 1 {
		log.Warnf("onPrem node has invalid number of credentials: %d, only first will be considered", len(credentialRefs))
	}

	credObj, err := getCredObject(ctx, cli, credentialRefs[0].UUID)
	if err != nil {
		return err
	}

	i.Username = credObj.SSHUser
	i.Password = credObj.SSHPassword
	return nil
}

func (i *privateInstance) fillGWDetails(
	ctx context.Context,
	cli *client.HTTP,
	log *logrus.Entry,
	gwNodes []*models.ContrailMulticloudGWNode,
	subnets []*models.NodeCloudPrivateSubnetRef,
) error {
	if len(gwNodes) == 0 {
		return nil
	}

	if len(gwNodes) > 1 {
		log.Warnf("onPrem node has invalid number of GWNodes: %d, only first will be considered", len(gwNodes))
	}

	gwNodeRef := gwNodes[0]
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

	i.Gateway = gwNodeResp.ContrailMulticloudGWNode.DefaultGateway

	if i.Gateway == "" {
		return fmt.Errorf(
			"default gateway is not set for contrail_multicloud_gw_node uuid: %s",
			gwNodeRef.UUID)
	}

	if err = i.fillSubNetDetails(ctx, cli, log, subnets); err != nil {
		return err
	}

	return nil
}

func (i *privateInstance) fillVrouterDetails(
	ctx context.Context,
	cli *client.HTTP,
	log *logrus.Entry,
	vrouterNodes []*models.ContrailVrouterNode,
) error {
	if len(vrouterNodes) == 0 {
		return nil
	}
	if len(vrouterNodes) > 1 {
		log.Warnf("onPrem node has invalid number of VrouterNodes: %d, only first will be considered", len(vrouterNodes))
	}

	vrouterNodeRef := vrouterNodes[0]
	vrouterNodeResponse, err := cli.GetContrailVrouterNode(ctx,
		&services.GetContrailVrouterNodeRequest{
			ID: vrouterNodeRef.UUID,
		},
	)
	if err != nil {
		return err
	}

	vrouterNode := vrouterNodeResponse.ContrailVrouterNode
	if vrouterNode.DefaultGateway != "" {
		i.Gateway = vrouterNode.DefaultGateway
		return nil
	}
	contrailClusterResponse, err := cli.GetContrailCluster(ctx,
		&services.GetContrailClusterRequest{
			ID: vrouterNode.ParentUUID,
		},
	)
	if err != nil {
		return err
	}
	i.Gateway = contrailClusterResponse.ContrailCluster.DefaultGateway

	if i.Gateway == "" {
		return fmt.Errorf(
			`default gateway is neither set for vrouter_node uuid: %s
		nor for contrail_cluster uuid: %s`,
			vrouterNodeRef.UUID, vrouterNode.ParentUUID)
	}
	return nil
}

func (i *privateInstance) fillSubNetDetails(
	ctx context.Context,
	cli *client.HTTP,
	log *logrus.Entry,
	subnets []*models.NodeCloudPrivateSubnetRef,
) error {

	if len(subnets) == 0 {
		return errors.New("onPrem node has no subnets")
	}
	if len(subnets) > 1 {
		log.Warnf("onPrem node has invalid number of subnets: %d, only first will be considered", len(subnets))
	}

	subnetRef := subnets[0]
	privateSubnetResp, err := cli.GetCloudPrivateSubnet(ctx,
		&services.GetCloudPrivateSubnetRequest{
			ID: subnetRef.UUID,
		})
	if err != nil {
		return err
	}

	i.PrivateSubnets = []string{privateSubnetResp.CloudPrivateSubnet.CidrBlock}
	return nil
}

func (i *privateInstance) fillRoles(n *models.Node) {
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

func (c *onPremCloud) createTors(
	ctx context.Context,
	cli *client.HTTP,
	log *logrus.Entry,
	cloudUUID string,
	tags []*models.Tag,
) error {
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

		for _, keyValuePair := range t.GetAnnotations().GetKeyValuePair() {
			switch keyValuePair.Key {
			case "autonomous_system":
				if tt.AS, err = strconv.Atoi(keyValuePair.Value); err != nil {
					return errors.Wrap(err, "fail to parse autonomous_system annotation")
				}
			case "interface":
				tt.Interface = strings.Split(keyValuePair.Value, ",")
			case "private_subnet":
				tt.PrivateSubnets = strings.Split(keyValuePair.Value, ",")
			}
		}
		c.Instances = append(c.Instances, tt)
	}
	return nil
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
