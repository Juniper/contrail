package cluster

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"strconv"
)

// Bms represents baremetal provider of ansible inventory
type Bms struct {
	SSHPassword   string `yaml:"ssh_pwd,omitempty"`
	SSHUser       string `yaml:"ssh_user,omitempty"`
	SSHPublicKey  string `yaml:"ssh_public_key,omitempty"`
	SSHPrivateKey string `yaml:"ssh_private_key,omitempty"`
	NtpServer     string `yaml:"ntpserver,omitempty"`
	DomainSuffix  string `yaml:"domainsuffix,omitempty"`
}

// Kvm represents qemu provider of ansible inventory
type Kvm struct {
	Image         string `yaml:"image"`
	ImageURL      string `yaml:"url"`
	SSHPassword   string `yaml:"ssh_pwd,omitempty"`
	SSHUser       string `yaml:"ssh_user,omitempty"`
	SSHPublicKey  string `yaml:"ssh_public_key,omitempty"`
	SSHPrivateKey string `yaml:"ssh_private_key,omitempty"`
	Vcpu          string `yaml:"vcpu,omitempty"`
	Vram          string `yaml:"vram,omitempty"`
	Vdisk         string `yaml:"vdisk,omitempty"`
	SubnetPrefix  string `yaml:"subnet_prefix,omitempty"`
	SubnetNetmask string `yaml:"subnet_netmask,omitempty"`
	Gateway       string `yaml:"gateway,omitempty"`
	NameServer    string `yaml:"nameserver,omitempty"`
	NtpServer     string `yaml:"ntpserver,omitempty"`
	DomainSuffix  string `yaml:"domainsuffix,omitempty"`
}

// Gce represents google cloud provider of ansible inventory
type Gce struct {
	ServiceAccountEmail string `yaml:"service_account_email"`
	CredentialsFile     string `yaml:"credentials_file"`
	ProjectID           string `yaml:"project_id"`
	SSHUser             string `yaml:"ssh_user"`
	SSHPassword         string `yaml:"ssh_pwd,omitempty"`
	SSHPrivateKey       string `yaml:"ssh_private_key,omitempty"`
	MachineType         string `yaml:"machine_type"`
	Image               string `yaml:"image"`
	Network             string `yaml:"network,omitempty"`
	SubnetWork          string `yaml:"subnetwork,omitempty"`
	Zone                string `yaml:"zone,omitempty"`
	DiskSize            string `yaml:"disk_size"`
}

// ProviderConfig represents types of providers in ansible inventory
type ProviderConfig struct {
	Bms Bms `yaml:"bms,omitempty"`
	Kvm Kvm `yaml:"kvm,omitempty"`
	Gce Gce `yaml:"gce,omitempty"`
}

// Roles represents role played by instance in ansible inventory
type Roles struct {
	ConfigDB            bool `yaml:"config_database,omitempty"`
	Config              bool `yaml:"config,omitempty"`
	Control             bool `yaml:"control,omitempty"`
	AnalyticsDB         bool `yaml:"analytics_database,omitempty"`
	Analytics           bool `yaml:"analytics,omitempty"`
	Webui               bool `yaml:"webui,omitempty"`
	Vrouter             bool `yaml:"vrouter,omitempty"`
	OpenstackControl    bool `yaml:"openstack_control,omitempty"`
	OpenstackNetwork    bool `yaml:"openstack_network,omitempty"`
	OpenstackStorage    bool `yaml:"openstack_storage,omitempty"`
	OpenstackMonitoring bool `yaml:"openstack_monitoring,omitempty"`
}

// Instance represents every node in ansible inventory
type Instance struct {
	Provider string `yaml:"provider"`
	IP       string `yaml:"ip"`
	Roles    *Roles `yaml:"roles"`
}

func apiToAnsibleInventory() {

	instancesMap := map[string]map[string]*Roles{}
	instancesMap["bms"] = map[string]*Roles{}
	roles := &Roles{}
	instancesMap["bms"]["1.1.1.1"] = roles
	roles.Config = true
	roles.ConfigDB = true
	roles.Control = true
	instancesMap["bms"]["1.1.1.2"] = &Roles{Control: true}

	instances := map[string]map[string]*Instance{}
	instances["instances"] = map[string]*Instance{}
	for provider, instanceMap := range instancesMap {
		i := 0
		for ip, roles := range instanceMap {
			i = i + 1
			stri := strconv.Itoa(i)
			instances["instances"][provider+stri] = &Instance{
				Provider: provider,
				IP:       ip,
				Roles:    roles,
			}
		}
	}

	y, _ := yaml.Marshal(&instances)
	fmt.Println(string(y))
}
