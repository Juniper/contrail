package cluster

// TODO:(ijohnson) generate constant resources
const (
	DefaultResource                       = "contrail-cluster"
	DefaultResourcePath                   = "/" + DefaultResource
	DefaultFilePermRWOnly                 = 0600
	DefaultWorkRoot                       = "/var/tmp/contrail_cluster"
	DefaultInstanceTemplate               = "instances.tmpl"
	DefaultInstanceFile                   = "instances.yml"
	DefaultVcenterTemplate                = "vcenter_vars.tmpl"
	DefaultVcenterFile                    = "vcenter_vars.yml"
	MCProvisioner                         = "multi-cloud"
	DefaultDeployer                       = "ansible"
	DefaultAnsibleRepo                    = "contrail-ansible-deployer"
	DefaultAnsibleRepoDir                 = "/usr/share/contrail/"
	DefaultAnsibleRepoInContainer         = "/root/contrail-ansible-deployer"
	DefaultvCenterProvPlay                = "playbooks/vcenter.yml"
	DefaultContrailProvPlay               = "playbooks/install_contrail.yml"
	DefaultContrailDestoryPlay            = "playbooks/contrail_destroy.yml"
	DefaultInstanceProvPlay               = "playbooks/provision_instances.yml"
	DefaultInstanceConfPlay               = "playbooks/configure_instances.yml"
	DefaultOpenstackProvPlay              = "playbooks/install_openstack.yml"
	DefaultOpenstackDestoryPlay           = "playbooks/openstack_destroy.yml"
	DefaultKubernetesProvPlay             = "playbooks/install_k8s.yml"
	DefaultAnsibleDatapathEncryptionRepo  = "contrail-datapath-encryption/ansible"
	DefaultContrailDatapathEncryptionPlay = "playbooks/deploy_and_run_all.yml"
	DefaultInventoryTemplate              = "inventory.tmpl"
	DefaultInventoryFile                  = "inventory.yml"

	defaultAppformixAnsibleRepoDir = "/usr/share/contrail/"
	defaultAppformixImageDir       = "/opt/software/appformix/"
	defaultAppformixAnsibleRepo    = "appformix-ansible-deployer"
	defaultAppformixProvPlay       = "playbooks/install_appformix.yml"
	defaultAppformixDir            = "appformix/"
	defaultXflowDir                = "xflow/"

	// TODO (ijohnson): Fix LP#1756958 and remove the status constants
	StatusField          = "provisioning_state"
	StatusNoState        = "NOSTATE"
	StatusCreated        = "CREATED"
	StatusCreateProgress = "CREATE_IN_PROGRESS"
	StatusCreateFailed   = "CREATE_FAILED"
	StatusUpdated        = "UPDATED"
	StatusUpdateProgress = "UPDATE_IN_PROGRESS"
	StatusUpdateFailed   = "UPDATE_FAILED"

	orchestratorOpenstack  = "openstack"
	orchestratorKubernetes = "kubernetes"
	orchestratorVcenter    = "vcenter"

	multicloudCLI = "deployer"
)
