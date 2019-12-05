package cloud

import (
	"context"
	"strings"

	"github.com/pkg/errors"

	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

type publicCloud struct {
	Providers []*publicProvider
}

type publicProvider struct {
	Type         string            `yaml:"provider,omitempty"`
	Organization string            `yaml:"organization,omitempty"`
	Project      string            `yaml:"project,omitempty"`
	Prebuild     string            `yaml:"prebuild,omitempty"`
	Tags         map[string]string `yaml:"tags,omitempty"`
	Regions      []region          `yaml:"regions,omitempty"`
}

type region interface {
	addVPC(v *virtualCloud)
	fill(ctx context.Context, cli *client.HTTP, regionUUID string, providerType string) error
}

type azureRegion struct {
	Name          string
	ResourceGroup string          `yaml:"resource_group,omitempty"`
	VNET          []*virtualCloud `yaml:"vnet,omitempty"`
}

type awsGCPRegion struct {
	Name string
	VPC  []*virtualCloud `yaml:"vpc,omitempty"`
}

type virtualCloud struct {
	Name      string            `yaml:"name,omitempty"`
	CIDRBlock string            `yaml:"cidr_block,omitempty"`
	Subnets   []*subnet         `yaml:"subnets,omitempty"`
	GcpFWs    gcpFirewalls      `yaml:",inline,omitempty"`
	SG        []sgRule          `yaml:"security_groups,omitempty"`
	Instances []*publicInstance `yaml:"instances,omitempty"`
}

type subnet struct {
	Name             string `yaml:"name"`
	CIDRBlock        string `yaml:"cidr_block"`
	AvailabilityZone string `yaml:"availability_zone,omitempty"`
	SecurityGroup    string `yaml:"security_group,omitempty"`
}

type sgRule interface {
	allow(*models.CloudSecurityGroupRule)
}

type awsSecurityGroup struct {
	Name      string
	direction string
	Rule      map[string]awsSecurityGroupRule `yaml:",inline,omitempty"`
}

type awsSecurityGroupRule struct {
	From       int64    `yaml:"from_port"`
	To         int64    `yaml:"to_port"`
	Protocol   string   `yaml:"protocol,omitempty"`
	CIDRBlocks []string `yaml:"cidr_blocks,omitempty"`
}

type azureSecurityGroup struct {
	Name  string
	Rules []*azSecurityGroupRule
}

type azSecurityGroupRule struct {
	Name      string
	Direction string
}

type gcpFirewalls struct {
	FirewallsExternal []*gcpSecurityGroup `yaml:"firewalls_external,omitempty"`
	FirewallsInternal []*gcpSecurityGroup `yaml:"firewalls_internal,omitempty"`
}

type gcpSecurityGroup struct {
	Name              string
	Allow             *gcpSecurityGroupRule
	Direction         string   `yaml:"direction,omitempty"`
	SourceRanges      []string `yaml:"source_ranges,omitempty"`
	DestinationRanges []string `yaml:"destination_ranges,omitempty"`
}

type gcpSecurityGroupRule struct {
	Protocol string  `yaml:"protocol"`
	Ports    []int64 `yaml:"ports,omitempty"`
}

type publicInstance struct {
	Name             string         `yaml:"name,omitempty"`
	Roles            []instanceRole `yaml:"roles,omitempty"`
	Provision        bool           `yaml:"provision,omitempty"`
	Username         string         `yaml:"username,omitempty"`
	OS               string         `yaml:"os,omitempty"`
	InstanceType     string         `yaml:"instance_type,omitempty"`
	Subnets          string         `yaml:"subnets,omitempty"`
	AvailabilityZone string         `yaml:"availability_zone,omitempty"`
	ProtocolsMode    []string       `yaml:"protocols_mode,omitempty"`
	SecurityGroups   []string       `yaml:"security_groups,omitempty"`
}

func (c *publicCloud) fill(ctx context.Context, cli *client.HTTP, cloudUUID string) error {
	cloudObj, err := cli.GetCloud(ctx, &services.GetCloudRequest{
		ID: cloudUUID,
	})
	if err != nil {
		return err
	}
	c.Providers = []*publicProvider{}
	for _, p := range cloudObj.Cloud.CloudProviders {
		pp := &publicProvider{
			Type:         p.Type,
			Organization: cloudObj.Cloud.Organization,
			Project:      p.Project,
			Prebuild:     cloudObj.Cloud.PrebuildImageID,
		}
		if err = pp.fill(ctx, cli, p.UUID, pp.Type); err != nil {
			return err
		}
		if pp.Type == gcp {
			pp.Type = google
		}
		c.Providers = append(c.Providers, pp)
	}
	return nil
}

func (p *publicProvider) fill(ctx context.Context, cli *client.HTTP, providerUUID string, providerType string) error {
	providerObj, err := cli.GetCloudProvider(ctx, &services.GetCloudProviderRequest{
		ID: providerUUID,
	})
	if err != nil {
		return err
	}
	p.Regions = []region{}
	var rr region
	for _, r := range providerObj.CloudProvider.CloudRegions {
		if providerType == azure {
			rr = &azureRegion{
				Name: r.DisplayName,
			}
		} else {
			rr = &awsGCPRegion{
				Name: r.DisplayName,
			}
		}
		if err = rr.fill(ctx, cli, r.UUID, providerType); err != nil {
			return err
		}
		p.Regions = append(p.Regions, rr)
	}

	if providerType == gcp {
		p.fillDefaultTags()
	}

	return nil
}

func (p *publicProvider) fillDefaultTags() {
	p.Tags = map[string]string{
		"owner":    "juniper",
		"project":  "contrail_multicloud",
		"build_id": "latest",
	}
}

func (r *azureRegion) fill(ctx context.Context, cli *client.HTTP, regionUUID string, providerType string) error {
	regionObj, err := cli.GetCloudRegion(ctx, &services.GetCloudRegionRequest{
		ID: regionUUID,
	})
	if err != nil {
		return err
	}

	r.ResourceGroup = regionObj.CloudRegion.ResourceGroup

	r.VNET = []*virtualCloud{}
	for _, v := range regionObj.CloudRegion.VirtualClouds {
		vv := &virtualCloud{
			Name:      v.Name,
			CIDRBlock: v.CidrBlock,
		}
		if err = vv.fill(ctx, cli, v.UUID, providerType); err != nil {
			return err
		}
		r.addVPC(vv)
	}
	return nil
}

func (r *azureRegion) addVPC(v *virtualCloud) {
	r.VNET = append(r.VNET, v)
}

func (r *awsGCPRegion) fill(ctx context.Context, cli *client.HTTP, regionUUID string, providerType string) error {
	regionObj, err := cli.GetCloudRegion(ctx, &services.GetCloudRegionRequest{
		ID: regionUUID,
	})
	if err != nil {
		return err
	}

	r.VPC = []*virtualCloud{}
	for _, v := range regionObj.CloudRegion.VirtualClouds {
		vv := &virtualCloud{
			Name:      v.Name,
			CIDRBlock: v.CidrBlock,
		}
		if err = vv.fill(ctx, cli, v.UUID, providerType); err != nil {
			return err
		}
		r.addVPC(vv)
	}
	return nil
}

func (r *awsGCPRegion) addVPC(v *virtualCloud) {
	r.VPC = append(r.VPC, v)
}

func (v *virtualCloud) fill(ctx context.Context, cli *client.HTTP, vCloudUUID string, providerType string) error {
	vCloudObj, err := cli.GetVirtualCloud(ctx, &services.GetVirtualCloudRequest{
		ID: vCloudUUID,
	})
	if err != nil {
		return err
	}
	v.Subnets = []*subnet{}
	for _, s := range vCloudObj.VirtualCloud.CloudPrivateSubnets {
		ss := &subnet{
			Name:             s.Name,
			CIDRBlock:        s.CidrBlock,
			AvailabilityZone: s.AvailabilityZone,
		}

		if providerType == azure && len(vCloudObj.VirtualCloud.CloudSecurityGroups) > 0 {
			ss.SecurityGroup = vCloudObj.VirtualCloud.CloudSecurityGroups[0].Name
		}
		v.Subnets = append(v.Subnets, ss)
	}

	v.Instances = []*publicInstance{}
	instances, err := getInstancesObjects(ctx, cli, vCloudObj.VirtualCloud.TagRefs)
	if err != nil {
		return err
	}
	for _, i := range instances {
		ii := &publicInstance{
			Name:         i.Hostname,
			Provision:    true,
			OS:           i.CloudInfo.OperatingSystem,
			InstanceType: i.CloudInfo.InstanceType,
		}
		if err = ii.fill(ctx, cli, i.UUID, providerType); err != nil {
			return err
		}
		v.Instances = append(v.Instances, ii)
	}

	return v.fillSecurityGroups(ctx, cli, vCloudObj.VirtualCloud.CloudSecurityGroups, vCloudUUID, providerType)
}

func (v *virtualCloud) fillSecurityGroups(
	ctx context.Context, cli *client.HTTP, sg []*models.CloudSecurityGroup, vCloudUUID, providerType string,
) error {
	v.SG = []sgRule{}
	for _, ssg := range sg {
		sgResp, err := cli.GetCloudSecurityGroup(ctx, &services.GetCloudSecurityGroupRequest{ID: ssg.UUID})
		if err != nil {
			return err
		}
		switch providerType {
		case aws:
			for _, sgr := range sgResp.CloudSecurityGroup.CloudSecurityGroupRules {
				r := &awsSecurityGroup{}
				r.allow(sgr)
				v.SG = append(v.SG, r)
			}
		case azure:
			s := &azureSecurityGroup{Name: sgResp.CloudSecurityGroup.Name}
			for _, sgr := range sgResp.CloudSecurityGroup.CloudSecurityGroupRules {
				s.allow(sgr)
				v.SG = append(v.SG, s)
			}
		case gcp:
			r := gcpFirewalls{}
			r.fillDefault(vCloudUUID)
			for _, sgr := range sgResp.CloudSecurityGroup.CloudSecurityGroupRules {
				r.allow(sgr)
			}
			v.GcpFWs = r
		default:
			return errors.Errorf("provider type %s is not a valid one", providerType)
		}
	}
	return nil
}

func (s *awsSecurityGroup) allow(r *models.CloudSecurityGroupRule) {
	s.Name = r.Name
	s.direction = r.Direction
	protocol := r.Protocol
	if r.Protocol == "ANY" {
		protocol = "-1"
	}
	s.Rule = map[string]awsSecurityGroupRule{
		r.Direction: {
			From:       r.FromPort,
			To:         r.ToPort,
			Protocol:   protocol,
			CIDRBlocks: []string{r.CidrBlock},
		},
	}
}

func (s *azureSecurityGroup) allow(r *models.CloudSecurityGroupRule) {
	var direction string
	switch r.Direction {
	case "ingress":
		direction = "inbound"
	case "egress":
		direction = "outbound"
	}

	s.Rules = append(s.Rules, &azSecurityGroupRule{Name: r.Name, Direction: direction})
}

func (fw *gcpFirewalls) allow(r *models.CloudSecurityGroupRule) {
	sg := &gcpSecurityGroup{
		Name: r.Name,
		Allow: &gcpSecurityGroupRule{
			Protocol: r.Protocol,
			Ports:    r.Ports,
		},
	}

	sg.Direction = strings.ToUpper(r.Direction)

	if r.Protocol == "" || r.Protocol == "ANY" {
		sg.Allow.Protocol = "all"
	}

	cidr := r.CidrBlock
	if cidr == "" {
		cidr = "0.0.0.0/0"
	}

	if sg.Direction == "INGRESS" {
		sg.SourceRanges = []string{cidr}
	} else {
		sg.DestinationRanges = []string{cidr}
	}

	fw.FirewallsInternal = append(fw.FirewallsInternal, sg)
}

func (fw *gcpFirewalls) fillDefault(vCloudUUID string) {
	fw.FirewallsExternal = []*gcpSecurityGroup{
		{
			Name: "default-wan-tcp-" + vCloudUUID,
			Allow: &gcpSecurityGroupRule{
				Protocol: "tcp",
				Ports:    []int64{22, 443},
			},
		}, {
			Name: "default-wan-udp-" + vCloudUUID,
			Allow: &gcpSecurityGroupRule{
				Protocol: "udp",
				Ports:    []int64{4500},
			},
		}, {
			Name: "default-wan-vrrp-" + vCloudUUID,
			Allow: &gcpSecurityGroupRule{
				Protocol: "112",
			},
		},
	}
}

func getInstancesObjects(
	ctx context.Context, cli *client.HTTP, tagRefs []*models.VirtualCloudTagRef,
) ([]*models.Node, error) {
	nn := []*models.Node{}
	for _, t := range tagRefs {
		tag, err := cli.GetTag(ctx, &services.GetTagRequest{ID: t.UUID})
		if err != nil {
			return nil, err
		}
		nn = append(nn, tag.Tag.NodeBackRefs...)
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

func (i *publicInstance) fill(ctx context.Context, cli *client.HTTP, instanceUUID string, providerType string) error {
	node, err := cli.GetNode(ctx, &services.GetNodeRequest{
		ID: instanceUUID,
	})
	if err != nil {
		return err
	}

	username, err := getOSSpecificUsername(i.OS, providerType)
	if err != nil {
		return err
	}
	i.Username = username

	i.fillRoles(node.Node.CloudInfo.Roles)

	for _, gw := range node.Node.ContrailMulticloudGWNodeBackRefs {
		if err = i.fillGWDetails(ctx, cli, gw.UUID); err != nil {
			return err
		}
	}

	if len(node.Node.CloudPrivateSubnetRefs) == 1 {
		if err = i.fillSubnetDetails(ctx, cli, node.Node.CloudPrivateSubnetRefs[0].UUID); err != nil {
			return err
		}
	}

	if providerType == aws {
		return i.fillSecurityGroups(ctx, cli, node.Node.CloudSecurityGroupRefs)
	}

	return nil
}

func (i *publicInstance) fillRoles(roles []string) {
	for _, r := range roles {
		switch r {
		case "compute":
			i.Roles = append(i.Roles, computeNodeInstanceRole)
		case string(bareInstanceRole):
			i.Roles = []instanceRole{bareInstanceRole}
			break
		default:
			i.Roles = append(i.Roles, instanceRole(r))
		}
	}
}

func (i *publicInstance) fillSubnetDetails(ctx context.Context, cli *client.HTTP, subnetUUID string) error {
	s, err := cli.GetCloudPrivateSubnet(ctx, &services.GetCloudPrivateSubnetRequest{
		ID: subnetUUID,
	})
	if err != nil {
		return err
	}
	i.Subnets = s.CloudPrivateSubnet.Name
	i.AvailabilityZone = s.CloudPrivateSubnet.AvailabilityZone
	return nil
}

func (i *publicInstance) fillGWDetails(ctx context.Context, cli *client.HTTP, gwUUID string) error {
	gwNode, err := cli.GetContrailMulticloudGWNode(ctx, &services.GetContrailMulticloudGWNodeRequest{
		ID: gwUUID,
	})
	if err != nil {
		return err
	}

	i.ProtocolsMode = gwNode.ContrailMulticloudGWNode.ProtocolsMode
	return nil
}

func (i *publicInstance) fillSecurityGroups(
	ctx context.Context, cli *client.HTTP, sgRefs []*models.NodeCloudSecurityGroupRef,
) error {
	for _, sg := range sgRefs {
		ssg, err := cli.GetCloudSecurityGroup(ctx, &services.GetCloudSecurityGroupRequest{ID: sg.UUID})
		if err != nil {
			return err
		}
		for _, sgr := range ssg.CloudSecurityGroup.CloudSecurityGroupRules {
			i.SecurityGroups = append(i.SecurityGroups, sgr.Name)
		}
	}
	return nil
}
