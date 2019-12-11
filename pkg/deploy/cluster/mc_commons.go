package cluster

import (
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/pkg/errors"
)

const (
	defaultContrailCommonFile = "ansible-multicloud/contrail/common.yml"
	defaultTORCommonFile      = "ansible-multicloud/tor/common.yml"
	defaultGatewayCommonFile  = "ansible-multicloud/gateway/common.yml"
	defaultContrailConfigPort = 8082
	bgpSecret                 = "bgp_secret"
	pathConfig                = "/etc/multicloud"

	defaultContrailUser      = "admin"
	defaultContrailPassword  = "c0ntrail123"
	defaultContrailTenant    = "default-project"
	bgpMCRoutesForController = "False"
	debugMCRoutes            = "False"
	torBGPSecret             = "contrail_secret"
	torOSPFSecret            = "contrail_secret"
)

type contrailCommon struct {
	User            string                    `yaml:"contrail_user"`
	Password        string                    `yaml:"contrail_password"`
	Port            int                       `yaml:"contrail_port"`
	Tenant          string                    `yaml:"contrail_tenant"`
	Configuration   contrailConfiguration     `yaml:"contrail_configuration"`
	ProviderConfigs map[string]providerConfig `yaml:"provider_config"`
}

type providerConfig struct {
	SSHUser        string `yaml:"ssh_user"`
	SSHPassword    string `yaml:"ssh_pwd"`
	SSHPublicKey   string `yaml:"ssh_public_key,omitempty"`
	DomainSuffix   string `yaml:"domainsuffix,omitempty"`
	NTPServer      string `yaml:"ntpserver"`
	ManageEtcHosts bool   `yaml:"manage_etc_hosts"`
}

type contrailConfiguration struct {
	Version               string `yaml:"CONTRAIL_VERSION,omitempty"`
	EncapsulationPriority string `yaml:"ENCAP_PRIORITY"`
	CloudOrchestrator     string `yaml:"CLOUD_ORCHESTRATOR"`
	VrouterEncryption     string `yaml:"VROUTER_ENCRYPTION,omitempty"`
}

type gatewayCommon struct {
	ConfigPath                string   `yaml:"PATH_CONFIG"`
	SSLConfigLocalPath        string   `yaml:"PATH_SSL_CONFIG_LOCAL"`
	SSLConfigPath             string   `yaml:"PATH_SSL_CONFIG"`
	OpenVPNConfigPath         string   `yaml:"PATH_OPENVPN_CONFIG"`
	BirdConfigPath            string   `yaml:"PATH_BIRD_CONFIG"`
	StrongSwanConfigPath      string   `yaml:"PATH_STRONGSWAN_CONFIG"`
	VRRPConfigPath            string   `yaml:"PATH_VRRP_CONFIG"`
	AWSConfigPath             string   `yaml:"PATH_AWS_CONFIG"`
	InterfaceConfigPath       string   `yaml:"PATH_INTERFACE_CONFIG"`
	FwConfigPath              string   `yaml:"PATH_FW_CONFIG"`
	GCPConfigPath             string   `yaml:"PATH_GCP_CONFIG"`
	SecretConfigPath          string   `yaml:"PATH_SECRET_CONFIG"`
	VxLanConfigPath           string   `yaml:"PATH_VXLAN_CONFIG"`
	ContainerRegistry         string   `yaml:"CONTAINER_REGISTRY"`
	ContrailMulticloudVersion string   `yaml:"CONTRAIL_MULTICLOUD_VERSION,omitempty"`
	UpgradeKernel             string   `yaml:"UPGRADE_KERNEL"`
	KernelVersion             string   `yaml:"KERNEL_VERSION"`
	AutonomousSystem          int64    `yaml:"AS"`
	VPNLoNetwork              string   `yaml:"vpn_lo_network"`
	VPNNetwork                string   `yaml:"vpn_network"`
	OpenVPNPort               int64    `yaml:"openvpn_port"`
	BfdInterval               string   `yaml:"bfd_interval"`
	BfdMultiplier             int64    `yaml:"bfd_multiplier"`
	BfdIntervalMultihop       string   `yaml:"bfd_interval_multihop"`
	BfdMultiplierMultihop     int64    `yaml:"bfd_multiplier_multihop"`
	CoreBgpSecret             string   `yaml:"core_bgp_secret"`
	RegistryPrivateInsecure   []string `yaml:"registry_private_insecure,omitempty"`
}

type torCommon struct {
	BGPSecret                        string `yaml:"tor_bgp_secret"`
	OspfSecret                       string `yaml:"tor_ospf_secret"`
	DebugMulticloudRoutes            string `yaml:"DEBUG_MULTI_CLOUD_ROUTES"`
	BGPMulticloudRoutesForController string `yaml:"BGP_MULTI_CLOUD_ROUTES_FOR_CONTROLLER"`
}

// func (m *multiCloudProvisioner) createGatewayCommonFile() error {
// 	m.Log.Info("Creating gateway/common.yml input file for multi-cloud deployer")
// 	pContext := pongo2.Context{
// 		"cluster":    m.clusterData.ClusterInfo,
// 		"pathConfig": pathConfig,
// 		"bgpSecret":  bgpSecret,
// 	}

// 	data, err := yaml.Marshal(common)
// 	if err != nil {
// 		return errors.Wrap(err, "cannot marshal tor common yml file")
// 	}

// 	if err = fileutil.WriteToFile(m.getGatewayCommonFile(), data, defaultFilePermRWOnly); err != nil {
// 		return errors.Wrap(err, "cannot create gateway common yml file")
// 	}
// 	m.Log.Info("Created gateway/common.yml input file for multi-cloud deployer")
// 	return nil
// }

func (m *multiCloudProvisioner) createContrailCommonFile() error {
	m.Log.Info("Creating contrail/common.yml input file for multi-cloud deployer")

	encapPriority := m.clusterData.ClusterInfo.EncapPriority
	if encapPriority == "" {
		encapPriority = "MPLSoGRE,MPLSoUDP,VXLAN"
	}

	// TODO: Check if not nil and if is not empty value
	vrouterEncryption := m.clusterData.ClusterInfo.ContrailConfiguration.GetValue("VROUTER_ENCRYPTION")

	// TODO: Check if not nil and if is not empty value
	contrailVersion := m.clusterData.ClusterInfo.ContrailVersion
	if contrailVersion == "" {
		contrailVersion = m.clusterData.ClusterInfo.ContrailConfiguration.GetValue("CONTRAIL_CONTAINER_TAG")
	}

	common := contrailCommon{
		User:     defaultContrailUser,
		Password: m.getContrailPassword(),
		Tenant:   defaultContrailTenant,
		Port:     defaultContrailConfigPort,
		Configuration: contrailConfiguration{
			CloudOrchestrator:     "kubernetes",
			VrouterEncryption:     vrouterEncryption,
			EncapsulationPriority: encapPriority,
			Version:               contrailVersion,
		},
		ProviderConfigs: map[string]providerConfig{
			"bms": providerConfig{
				SSHUser:      m.clusterData.DefaultSSHUser,
				SSHPassword:  m.clusterData.DefaultSSHPassword,
				SSHPublicKey: m.clusterData.DefaultSSHKey,
				NTPServer:    m.clusterData.ClusterInfo.NTPServer,
				DomainSuffix: m.clusterData.ClusterInfo.DomainSuffix,
			},
		},
	}

	data, err := yaml.Marshal(common)
	if err != nil {
		return errors.Wrap(err, "cannot marshal contrail common yml file")
	}

	if err = fileutil.WriteToFile(m.getContrailCommonFile(), data, defaultFilePermRWOnly); err != nil {
		return errors.Wrap(err, "cannot create contrail common yml file")
	}

	m.Log.Info("Created contrail/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) createGatewayCommonFile() error {
	m.Log.Info("Creating gateway/common.yml input file for multi-cloud deployer")
	pathConfigVar := "{{ PATH_CONFIG }}"

	mcVersion := m.clusterData.ClusterInfo.ContrailVersion
	if mcVersion == "" {
		mcVersion = m.clusterData.ClusterInfo.ContrailConfiguration.GetValue("CONTRAIL_CONTAINER_TAG")
	}

	common := gatewayCommon{
		ConfigPath:                pathConfig,
		SSLConfigLocalPath:        "~/.multicloud/ssl",
		SSLConfigPath:             filepath.Join(pathConfigVar, "ssl"),
		OpenVPNConfigPath:         filepath.Join(pathConfigVar, "openvpn"),
		BirdConfigPath:            filepath.Join(pathConfigVar, "bird"),
		StrongSwanConfigPath:      filepath.Join(pathConfigVar, "strongswan"),
		VRRPConfigPath:            filepath.Join(pathConfigVar, "vrrp"),
		AWSConfigPath:             filepath.Join(pathConfigVar, "aws"),
		InterfaceConfigPath:       "/etc/network/interfaces.d",
		FwConfigPath:              filepath.Join(pathConfigVar, "firewall"),
		GCPConfigPath:             filepath.Join(pathConfigVar, "gcp"),
		SecretConfigPath:          filepath.Join(pathConfigVar, "secret"),
		VxLanConfigPath:           filepath.Join(pathConfigVar, "vxlan"),
		ContainerRegistry:         m.clusterData.ClusterInfo.ContainerRegistry,
		UpgradeKernel:             "False",
		KernelVersion:             "3.10.0-862.11.6.el7.x86_64",
		AutonomousSystem:          m.clusterData.ClusterInfo.MCGWInfo.AS,
		VPNLoNetwork:              m.clusterData.ClusterInfo.MCGWInfo.VPNLoNetwork,
		VPNNetwork:                m.clusterData.ClusterInfo.MCGWInfo.VPNNetwork,
		OpenVPNPort:               m.clusterData.ClusterInfo.MCGWInfo.OpenvpnPort,
		BfdInterval:               m.clusterData.ClusterInfo.MCGWInfo.BFDInterval,
		BfdMultiplier:             m.clusterData.ClusterInfo.MCGWInfo.BFDMultiplier,
		BfdIntervalMultihop:       m.clusterData.ClusterInfo.MCGWInfo.BFDIntervalMultihop,
		BfdMultiplierMultihop:     m.clusterData.ClusterInfo.MCGWInfo.BFDMultiplierMultihop,
		CoreBgpSecret:             bgpSecret,
		ContrailMulticloudVersion: mcVersion,
		RegistryPrivateInsecure:   m.getRegistryPrivateInsecure(),
	}

	data, err := yaml.Marshal(common)
	if err != nil {
		return errors.Wrap(err, "cannot marshal tor common yml file")
	}

	if err = fileutil.WriteToFile(m.getGatewayCommonFile(), data, defaultFilePermRWOnly); err != nil {
		return errors.Wrap(err, "cannot create tor common yml file")
	}
	m.Log.Info("Created gateway/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) getRegistryPrivateInsecure() []string {
	if m.clusterData.ClusterInfo.RegistryPrivateInsecure {
		return []string{m.clusterData.ClusterInfo.ContainerRegistry}
	}
	return nil
}

func (m *multiCloudProvisioner) createTORCommonFile() error {
	m.Log.Info("Creating tor/common.yml input file for multi-cloud deployer")

	common := torCommon{
		BGPMulticloudRoutesForController: bgpMCRoutesForController,
		BGPSecret:                        torBGPSecret,
		DebugMulticloudRoutes:            debugMCRoutes,
		OspfSecret:                       torOSPFSecret,
	}
	data, err := yaml.Marshal(common)
	if err != nil {
		return errors.Wrap(err, "cannot marshal tor common yml file")
	}

	if err = fileutil.WriteToFile(m.getTORCommonFile(), data, defaultFilePermRWOnly); err != nil {
		return errors.Wrap(err, "cannot create tor common yml file")
	}
	m.Log.Info("Created tor/common.yml input file for multi-cloud deployer")
	return nil
}
