package cluster

import (
	"path/filepath"
	"strings"

	"github.com/Juniper/asf/pkg/fileutil"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

const (
	defaultContrailCommonFile = "ansible-multicloud/contrail/common.yml"
	defaultTORCommonFile      = "ansible-multicloud/tor/common.yml"
	defaultGatewayCommonFile  = "ansible-multicloud/gateway/common.yml"
	// Contrail
	defaultContrailUser          = "admin"
	defaultContrailPassword      = "c0ntrail123"
	defaultContrailTenant        = "default-project"
	defaultContrailConfigPort    = 8082
	defaultEncapsulationPriority = "MPLSoGRE,MPLSoUDP,VXLAN"
	defaultOrchestrator          = "kubernetes"
	defaultProvider              = "bms"
	// Gateway
	defaultConfigPath          = "/etc/multicloud"
	defaultLocalSSLConfigPath  = "~/.multicloud/ssl"
	defaultBGPSecret           = "bgp_secret"
	defaultKernelVersion       = "3.10.0-862.11.6.el7.x86_64"
	defaultInstanceConfigPath  = "/etc/network/interfaces.d"
	defaultUpgradeKernelPolicy = false
	// TOR
	defaultTORBGPSecret      = "contrail_secret"
	defaultTOROSPFSecret     = "contrail_secret"
	bgpMCRoutesForController = false
	debugMCRoutes            = false
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

func (m *multiCloudProvisioner) createContrailCommonFile() error {
	m.Log.Info("Creating contrail/common.yml input file for multi-cloud deployer")

	common := contrailCommon{
		User:     defaultContrailUser,
		Password: m.getContrailPassword(),
		Tenant:   defaultContrailTenant,
		Port:     defaultContrailConfigPort,
		Configuration: contrailConfiguration{
			CloudOrchestrator:     defaultOrchestrator,
			VrouterEncryption:     m.clusterData.ClusterInfo.ContrailConfiguration.GetValue("VROUTER_ENCRYPTION"),
			EncapsulationPriority: m.getEncapsulationPriority(),
			Version:               m.getContrailVersion(),
		},
		ProviderConfigs: map[string]providerConfig{
			defaultProvider: {
				SSHUser:      m.clusterData.DefaultSSHUser,
				SSHPassword:  m.clusterData.DefaultSSHPassword,
				SSHPublicKey: m.clusterData.DefaultSSHKey,
				NTPServer:    m.clusterData.ClusterInfo.NTPServer,
				DomainSuffix: m.clusterData.ClusterInfo.DomainSuffix,
			},
		},
	}

	if err := m.marshalAndSave(common, m.getContrailCommonFile()); err != nil {
		return errors.Wrap(err, "cannot create contrail common yml file")
	}
	m.Log.Info("Created contrail/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) getContrailVersion() string {
	if m.clusterData.ClusterInfo.ContrailVersion == "" {
		return m.clusterData.ClusterInfo.ContrailConfiguration.GetValue("CONTRAIL_CONTAINER_TAG")
	}
	return m.clusterData.ClusterInfo.ContrailVersion
}

func (m *multiCloudProvisioner) getEncapsulationPriority() string {
	if m.clusterData.ClusterInfo.EncapPriority == "" {
		return defaultEncapsulationPriority
	}
	return m.clusterData.ClusterInfo.EncapPriority
}

func (m *multiCloudProvisioner) marshalAndSave(data interface{}, dst string) error {
	d, err := yaml.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "cannot marshal file")
	}
	if err = fileutil.WriteToFile(dst, d, defaultFilePermRWOnly); err != nil {
		return err
	}
	return nil
}

func (m *multiCloudProvisioner) getContrailPassword() string {
	if strings.ToLower(m.clusterData.ClusterInfo.Orchestrator) == openstack {
		openStackClusterInfo := m.clusterData.GetOpenstackClusterInfo()
		if o := openStackClusterInfo.KollaPasswords.GetValue("keystone_admin_password"); o != "" {
			return o
		}
	}
	return defaultContrailPassword
}

func (m *multiCloudProvisioner) getContrailCommonFile() string {
	return filepath.Join(m.workDir, defaultContrailCommonFile)
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
	UpgradeKernel             bool     `yaml:"UPGRADE_KERNEL"`
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

func (m *multiCloudProvisioner) createGatewayCommonFile() error {
	m.Log.Info("Creating gateway/common.yml input file for multi-cloud deployer")
	pathConfigVar := "{{ PATH_CONFIG }}"

	common := gatewayCommon{
		ConfigPath:                defaultConfigPath,
		SSLConfigLocalPath:        defaultLocalSSLConfigPath,
		SSLConfigPath:             filepath.Join(pathConfigVar, "ssl"),
		OpenVPNConfigPath:         filepath.Join(pathConfigVar, "openvpn"),
		BirdConfigPath:            filepath.Join(pathConfigVar, "bird"),
		StrongSwanConfigPath:      filepath.Join(pathConfigVar, "strongswan"),
		VRRPConfigPath:            filepath.Join(pathConfigVar, "vrrp"),
		AWSConfigPath:             filepath.Join(pathConfigVar, "aws"),
		InterfaceConfigPath:       defaultInstanceConfigPath,
		FwConfigPath:              filepath.Join(pathConfigVar, "firewall"),
		GCPConfigPath:             filepath.Join(pathConfigVar, "gcp"),
		SecretConfigPath:          filepath.Join(pathConfigVar, "secret"),
		VxLanConfigPath:           filepath.Join(pathConfigVar, "vxlan"),
		ContainerRegistry:         m.clusterData.ClusterInfo.ContainerRegistry,
		UpgradeKernel:             defaultUpgradeKernelPolicy,
		KernelVersion:             defaultKernelVersion,
		AutonomousSystem:          m.clusterData.ClusterInfo.MCGWInfo.AS,
		VPNLoNetwork:              m.clusterData.ClusterInfo.MCGWInfo.VPNLoNetwork,
		VPNNetwork:                m.clusterData.ClusterInfo.MCGWInfo.VPNNetwork,
		OpenVPNPort:               m.clusterData.ClusterInfo.MCGWInfo.OpenvpnPort,
		BfdInterval:               m.clusterData.ClusterInfo.MCGWInfo.BFDInterval,
		BfdMultiplier:             m.clusterData.ClusterInfo.MCGWInfo.BFDMultiplier,
		BfdIntervalMultihop:       m.clusterData.ClusterInfo.MCGWInfo.BFDIntervalMultihop,
		BfdMultiplierMultihop:     m.clusterData.ClusterInfo.MCGWInfo.BFDMultiplierMultihop,
		CoreBgpSecret:             defaultBGPSecret,
		ContrailMulticloudVersion: m.getMultiCloudVersion(),
		RegistryPrivateInsecure:   m.getRegistryPrivateInsecure(),
	}

	if err := m.marshalAndSave(common, m.getGatewayCommonFile()); err != nil {
		return errors.Wrap(err, "cannot create gateway common yml file")
	}
	m.Log.Info("Created gateway/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) getMultiCloudVersion() string {
	if m.clusterData.ClusterInfo.ContrailVersion == "" {
		return m.clusterData.ClusterInfo.ContrailConfiguration.GetValue("CONTRAIL_CONTAINER_TAG")
	}
	return m.clusterData.ClusterInfo.ContrailVersion
}

func (m *multiCloudProvisioner) getRegistryPrivateInsecure() []string {
	if m.clusterData.ClusterInfo.RegistryPrivateInsecure {
		return []string{m.clusterData.ClusterInfo.ContainerRegistry}
	}
	return nil
}

func (m *multiCloudProvisioner) getGatewayCommonFile() string {
	return filepath.Join(m.workDir, defaultGatewayCommonFile)
}

type torCommon struct {
	BGPSecret                        string `yaml:"tor_bgp_secret"`
	OspfSecret                       string `yaml:"tor_ospf_secret"`
	DebugMulticloudRoutes            bool   `yaml:"DEBUG_MULTI_CLOUD_ROUTES"`
	BGPMulticloudRoutesForController bool   `yaml:"BGP_MULTI_CLOUD_ROUTES_FOR_CONTROLLER"`
}

func (m *multiCloudProvisioner) createTORCommonFile() error {
	m.Log.Info("Creating tor/common.yml input file for multi-cloud deployer")

	common := torCommon{
		BGPMulticloudRoutesForController: bgpMCRoutesForController,
		DebugMulticloudRoutes:            debugMCRoutes,
		BGPSecret:                        defaultTORBGPSecret,
		OspfSecret:                       defaultTOROSPFSecret,
	}
	if err := m.marshalAndSave(common, m.getTORCommonFile()); err != nil {
		return errors.Wrap(err, "cannot create tor common yml file")
	}
	m.Log.Info("Created tor/common.yml input file for multi-cloud deployer")
	return nil
}

func (m *multiCloudProvisioner) getTORCommonFile() string {
	return filepath.Join(m.workDir, defaultTORCommonFile)
}
