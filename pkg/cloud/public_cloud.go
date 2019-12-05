package cloud

import (
	"context"
	"strings"

	"github.com/Juniper/contrail/pkg/client"
	"github.com/Juniper/contrail/pkg/models"
	"github.com/Juniper/contrail/pkg/services"
)

type pubCloud struct {
	Providers []*pubProvider
}

type pubProvider struct {
	Type         string            `yaml:"provider,omitempty"`
	Organization string            `yaml:"organization,omitempty"`
	Project      string            `yaml:"project,omitempty"`
	Prebuild     string            `yaml:"prebuild,omitempty"`
	Tags         map[string]string `yaml:"tags,omitempty"`
	Regions      []region          `yaml:"regions,omitempty"`
}

type region interface {
	addVPC()
	fill(ctx context.Context, cli *client.HTTP, regionUUID string, providerType string) error
}

type azRegion struct {
	Name          string
	ResourceGroup string    `yaml:"resource_group,omitempty"`
	VNET          []*vCloud `yaml:"vnet,omitempty"`
}

func (r *azRegion) addVPC() {}

type awsGCPRegion struct {
	Name string
	VPC  []*vCloud `yaml:"vpc,omitempty"`
}

func (r *awsGCPRegion) addVPC() {}

type sgRule interface {
	allow(*models.CloudSecurityGroupRule)
}

type awsSG struct {
	Name      string
	direction string
	Rule      map[string]awsSGRule `yaml:",inline,omitempty"`
}

type awsSGRule struct {
	From       int64    `yaml:"from_port"`
	To         int64    `yaml:"to_port"`
	Protocol   string   `yaml:"protocol,omitempty"`
	CIDRBlocks []string `yaml:"cidr_blocks,omitempty"`
}

func (s *awsSG) allow(r *models.CloudSecurityGroupRule) {
	s.Name = r.Name
	s.direction = r.Direction
	protocol := r.Protocol
	if r.Protocol == "ANY" {
		protocol = "-1"
	}
	s.Rule = map[string]awsSGRule{
		r.Direction: {
			From:       r.FromPort,
			To:         r.ToPort,
			Protocol:   protocol,
			CIDRBlocks: []string{r.CidrBlock},
		},
	}
}

type azSG struct {
	Name  string
	Rules []azSGRule
}

type azSGRule struct {
	Name      string
	Direction string
}

func (s *azSG) allow(r *models.CloudSecurityGroupRule) {
	var direction string
	switch r.Direction {
	case "ingress":
		direction = "inbound"
	case "egress":
		direction = "outbound"
	}

	s.Rules = append(s.Rules, azSGRule{Name: r.Name, Direction: direction})
}

type vCloud struct {
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

type gcpFirewalls struct {
	FirewallsExternal []*gcpSG `yaml:"firewalls_external,omitempty"`
	FirewallsInternal []*gcpSG `yaml:"firewalls_internal,omitempty"`
}

type gcpSG struct {
	Name              string
	Allow             *gcpSGRule
	Deny              *gcpSGRule
	Direction         string   `yaml:"direction,omitempty"`
	SourceRanges      []string `yaml:"source_ranges,omitempty"`
	DestinationRanges []string `yaml:"destination_ranges,omitempty"`
}

type gcpSGRule struct {
	Protocol string  `yaml:"protocol"`
	Ports    []int64 `yaml:"ports,omitempty"`
}

func (fw *gcpFirewalls) allow(r *models.CloudSecurityGroupRule) {
	sg := &gcpSG{
		Name: r.Name,
		Allow: &gcpSGRule{
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
	fw.FirewallsExternal = []*gcpSG{
		{
			Name: "default-wan-tcp-" + vCloudUUID,
			Allow: &gcpSGRule{
				Protocol: "tcp",
				Ports:    []int64{22, 443},
			},
		}, {
			Name: "default-wan-vrrp-" + vCloudUUID,
			Allow: &gcpSGRule{
				Protocol: "112",
			},
		},
	}
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

func (c *pubCloud) fill(ctx context.Context, cli *client.HTTP, cloudUUID string) error {
	cloudObj, err := cli.GetCloud(ctx, &services.GetCloudRequest{
		ID: cloudUUID,
	})
	if err != nil {
		return err
	}
	c.Providers = []*pubProvider{}
	for _, p := range cloudObj.Cloud.CloudProviders {
		pp := &pubProvider{
			Type:         p.Type,
			Organization: cloudObj.Cloud.Organization,
			Project:      p.Project,
			Prebuild:     cloudObj.Cloud.PrebuildImageID,
		}
		if err = pp.fill(ctx, cli, p.UUID, pp.Type); err != nil {
			return err
		}
		if pp.Type == gcp {
			pp.Type = "google"
		}
		c.Providers = append(c.Providers, pp)
	}
	return nil
}

func (p *pubProvider) fill(ctx context.Context, cli *client.HTTP, providerUUID string, providerType string) error {
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
			rr = &azRegion{
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
		p.Tags = map[string]string{
			"owner":    "juniper",
			"project":  "contrail_multicloud",
			"build_id": "latest",
		}
	}

	return nil
}

func (r *azRegion) fill(ctx context.Context, cli *client.HTTP, regionUUID string, providerType string) error {
	regionObj, err := cli.GetCloudRegion(ctx, &services.GetCloudRegionRequest{
		ID: regionUUID,
	})
	if err != nil {
		return err
	}

	r.ResourceGroup = regionObj.CloudRegion.ResourceGroup

	r.VNET = []*vCloud{}
	for _, v := range regionObj.CloudRegion.VirtualClouds {
		vv := &vCloud{
			Name:      v.Name,
			CIDRBlock: v.CidrBlock,
		}
		if err = vv.fill(ctx, cli, v.UUID, providerType); err != nil {
			return err
		}
		r.VNET = append(r.VNET, vv)
	}
	return nil
}

func (r *awsGCPRegion) fill(ctx context.Context, cli *client.HTTP, regionUUID string, providerType string) error {
	regionObj, err := cli.GetCloudRegion(ctx, &services.GetCloudRegionRequest{
		ID: regionUUID,
	})
	if err != nil {
		return err
	}

	r.VPC = []*vCloud{}
	for _, v := range regionObj.CloudRegion.VirtualClouds {
		vv := &vCloud{
			Name:      v.Name,
			CIDRBlock: v.CidrBlock,
		}
		if err = vv.fill(ctx, cli, v.UUID, providerType); err != nil {
			return err
		}
		r.VPC = append(r.VPC, vv)
	}
	return nil
}

func (v *vCloud) fill(ctx context.Context, cli *client.HTTP, vCloudUUID string, providerType string) error {
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

	return v.fillSG(ctx, cli, vCloudObj.VirtualCloud.CloudSecurityGroups, vCloudUUID, providerType)
}

func (v *vCloud) fillSG(
	ctx context.Context, cli *client.HTTP, sg []*models.CloudSecurityGroup, vCloudUUID, providerType string,
) error {
	v.SG = []sgRule{}
	for _, ssg := range sg {
		sgResp, err := cli.GetCloudSecurityGroup(ctx, &services.GetCloudSecurityGroupRequest{ID: ssg.UUID})
		if err != nil {
			return err
		}
		switch providerType {
		case AWS:
			for _, sgr := range sgResp.CloudSecurityGroup.CloudSecurityGroupRules {
				r := &awsSG{}
				r.allow(sgr)
				v.SG = append(v.SG, r)
			}
		case azure:
			s := &azSG{Name: sgResp.CloudSecurityGroup.Name}
			for _, sgr := range sgResp.CloudSecurityGroup.CloudSecurityGroupRules {
				s.allow(sgr)
				v.SG = append(v.SG, s)
			}
		case gcp:
			for _, sgr := range sgResp.CloudSecurityGroup.CloudSecurityGroupRules {
				r := gcpFirewalls{}
				r.fillDefault(vCloudUUID)
				r.allow(sgr)
				v.GcpFWs = r
			}
		}
	}
	return nil
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

//TODO: split to functions
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

	for _, r := range node.Node.CloudInfo.Roles {
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

	for _, gw := range node.Node.ContrailMulticloudGWNodeBackRefs {
		if err = i.fillGWDetails(ctx, cli, gw.UUID); err != nil {
			return err
		}
	}

	if len(node.Node.CloudPrivateSubnetRefs) > 0 {
		s, err := cli.GetCloudPrivateSubnet(ctx, &services.GetCloudPrivateSubnetRequest{
			ID: node.Node.CloudPrivateSubnetRefs[0].UUID,
		})
		if err != nil {
			return err
		}
		i.Subnets = s.CloudPrivateSubnet.Name
		i.AvailabilityZone = s.CloudPrivateSubnet.AvailabilityZone
	}

	if providerType == AWS {
		for _, sg := range node.Node.CloudSecurityGroupRefs {
			ssg, err := cli.GetCloudSecurityGroup(ctx, &services.GetCloudSecurityGroupRequest{ID: sg.UUID})
			if err != nil {
				return err
			}
			for _, sgr := range ssg.CloudSecurityGroup.CloudSecurityGroupRules {
				i.SecurityGroups = append(i.SecurityGroups, sgr.Name)
			}
		}
	}

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
