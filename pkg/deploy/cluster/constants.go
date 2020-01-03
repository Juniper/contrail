package cluster

// TODO:(ijohnson) generate constant resources
const (
	defaultResource                       = "contrail-cluster"
	defaultResourcePath                   = "/" + defaultResource
	defaultFilePermRWOnly                 = 0600
	defaultWorkRoot                       = "/var/tmp/contrail_cluster"
	defaultInstanceTemplate               = "instances.tmpl"
	defaultInstanceFile                   = "instances.yml"
	defaultVcenterTemplate                = "vcenter_vars.tmpl"
	defaultVcenterFile                    = "vcenter_vars.yml"
	mCProvisioner                         = "multi-cloud"
	defaultDeployer                       = "ansible"
	defaultAnsibleRepo                    = "contrail-ansible-deployer"
	defaultAnsibleRepoDir                 = "/usr/share/contrail/"
	defaultvCenterProvPlay                = "playbooks/vcenter.yml"
	defaultContrailProvPlay               = "playbooks/install_contrail.yml"
	defaultContrailDestoryPlay            = "playbooks/contrail_destroy.yml"
	defaultInstanceProvPlay               = "playbooks/provision_instances.yml"
	defaultInstanceConfPlay               = "playbooks/configure_instances.yml"
	defaultOpenstackProvPlay              = "playbooks/install_openstack.yml"
	defaultOpenstackDestoryPlay           = "playbooks/openstack_destroy.yml"
	defaultKubernetesProvPlay             = "playbooks/install_k8s.yml"
	defaultAnsibleDatapathEncryptionRepo  = "contrail-datapath-encryption/ansible"
	defaultContrailDatapathEncryptionPlay = "playbooks/deploy_and_run_all.yml"
	defaultInventoryTemplate              = "inventory.tmpl"
	defaultInventoryFile                  = "inventory.yml"

	defaultAppformixAnsibleRepoDir = "/usr/share/contrail/"
	defaultAppformixImageDir       = "/opt/software/appformix/"
	defaultAppformixAnsibleRepo    = "appformix-ansible-deployer"
	defaultAppformixProvPlay       = "playbooks/install_appformix.yml"
	defaultAppformixDir            = "appformix/"
	defaultXflowDir                = "xflow/"

	// TODO (ijohnson): Fix LP#1756958 and remove the status constants
	statusField          = "provisioning_state"
	statusNoState        = "NOSTATE"
	statusCreated        = "CREATED"
	statusCreateProgress = "CREATE_IN_PROGRESS"
	statusCreateFailed   = "CREATE_FAILED"

	statusUpdated        = "UPDATED"
	statusUpdateProgress = "UPDATE_IN_PROGRESS"
	statusUpdateFailed   = "UPDATE_FAILED"

	orchestratorOpenstack  = "openstack"
	orchestratorKubernetes = "kubernetes"
	orchestratorVcenter    = "vcenter"
)
