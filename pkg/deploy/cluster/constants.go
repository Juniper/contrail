package cluster

// TODO:(ijohnson) generate constant resources
const (
	defaultResource                             = "contrail-cluster"
	defaultResourcePath                         = "/" + defaultResource
	defaultK8sResource                          = "kubernetes-cluster"
	defaultK8sResourcePath                      = "/" + defaultK8sResource
	defaultOpenstackResource                    = "openstack-cluster"
	defaultOpenstackResourcePath                = "/" + defaultOpenstackResource
	defaultVCenterResource                      = "vCenter"
	defaultVCenterResourcePath                  = "/" + defaultVCenterResource
	defaultEndpointRes                          = "endpoint"
	defaultEndpointResPath                      = "/" + defaultEndpointRes
	defaultNodeRes                              = "node"
	defaultNodeResPath                          = "/" + defaultNodeRes
	defaultCredentialRes                        = "credential"
	defaultCredentialResPath                    = "/" + defaultCredentialRes
	defaultKeypairRes                           = "keypair"
	defaultKeypairResPath                       = "/" + defaultKeypairRes
	defaultContrailConfigNodeRes                = "contrail-config-node"
	defaultContrailConfigNodeResPath            = "/" + defaultContrailConfigNodeRes
	defaultContrailConfigDatabaseNodeRes        = "contrail-config-database-node"
	defaultContrailConfigDatabaseNodeResPath    = "/" + defaultContrailConfigDatabaseNodeRes
	defaultContrailControlNodeRes               = "contrail-control-node"
	defaultContrailControlNodeResPath           = "/" + defaultContrailControlNodeRes
	defaultContrailWebuiNodeRes                 = "contrail-webui-node"
	defaultContrailWebuiNodeResPath             = "/" + defaultContrailWebuiNodeRes
	defaultContrailAnalyticsNodeRes             = "contrail-analytics-node"
	defaultContrailAnalyticsNodeResPath         = "/" + defaultContrailAnalyticsNodeRes
	defaultContrailAnalyticsDatabaseNodeRes     = "contrail-analytics-database-node"
	defaultContrailAnalyticsDatabaseNodeResPath = "/" + defaultContrailAnalyticsDatabaseNodeRes
	defaultContrailAnalyticsAlarmNodeRes        = "contrail-analytics-alarm-node"
	defaultContrailAnalyticsAlarmNodeResPath    = "/" + defaultContrailAnalyticsAlarmNodeRes
	defaultContrailAnalyticsSNMPNodeRes         = "contrail-analytics-snmp-node"
	defaultContrailAnalyticsSNMPNodeResPath     = "/" + defaultContrailAnalyticsSNMPNodeRes
	defaultContrailVrouterNodeRes               = "contrail-vrouter-node"
	defaultContrailVrouterNodeResPath           = "/" + defaultContrailVrouterNodeRes
	defaultContrailMCGWNodeRes                  = "contrail-multicloud-gw-node"
	defaultContrailMCGWNodeResPath              = "/" + defaultContrailMCGWNodeRes
	defaultContrailServiceNodeRes               = "contrail-service-node"
	defaultContrailServiceNodeResPath           = "/" + defaultContrailServiceNodeRes
	defaultContrailZTPDHCPNodeRes               = "contrail-ztp-dhcp-node"
	defaultContrailZTPDHCPNodeResPath           = "/" + defaultContrailZTPDHCPNodeRes
	defaultContrailZTPTFTPNodeRes               = "contrail-ztp-tftp-node"
	defaultContrailZTPTFTPNodeResPath           = "/" + defaultContrailZTPTFTPNodeRes
	defaultKubernetesNodeRes                    = "kubernetes-node"
	defaultKubernetesNodeResPath                = "/" + defaultKubernetesNodeRes
	defaultKubernetesMasterNodeRes              = "kubernetes-master-node"
	defaultKubernetesMasterNodeResPath          = "/" + defaultKubernetesMasterNodeRes
	defaultKubernetesKubemanagerNodeRes         = "kubernetes-kubemanager-node"
	defaultKubernetesKubemanagerNodeResPath     = "/" + defaultKubernetesKubemanagerNodeRes
	defaultOpenstackControlNodeRes              = "openstack-control-node"
	defaultOpenstackControlNodeResPath          = "/" + defaultOpenstackControlNodeRes
	defaultOpenstackMonitoringNodeRes           = "openstack-monitoring-node"
	defaultOpenstackMonitoringNodeResPath       = "/" + defaultOpenstackMonitoringNodeRes
	defaultOpenstackNetworkNodeRes              = "openstack-network-node"
	defaultOpenstackNetworkNodeResPath          = "/" + defaultOpenstackNetworkNodeRes
	defaultOpenstackStorageNodeRes              = "openstack-storage-node"
	defaultOpenstackStorageNodeResPath          = "/" + defaultOpenstackStorageNodeRes
	defaultOpenstackComputeNodeRes              = "openstack-compute-node"
	defaultOpenstackComputeNodeResPath          = "/" + defaultOpenstackComputeNodeRes
	defaultVCenterComputeRes                    = "vCenter-compute"
	defaultVCenterComputeResPath                = "/" + defaultVCenterComputeRes
	defaultVCenterPluginNodeRes                 = "vCenter-plugin-node"
	defaultVCenterPluginNodeResPath             = "/" + defaultVCenterPluginNodeRes
	defaultVCenterManagerNodeRes                = "vCenter-manager-node"
	defaultVCenterManagerNodeResPath            = "/" + defaultVCenterManagerNodeRes

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
	defaultInstanceProvPlay               = "playbooks/provision_instances.yml"
	defaultInstanceConfPlay               = "playbooks/configure_instances.yml"
	defaultOpenstackProvPlay              = "playbooks/install_openstack.yml"
	defaultKubernetesProvPlay             = "playbooks/install_k8s.yml"
	defaultAnsibleDatapathEncryptionRepo  = "contrail-datapath-encryption/ansible"
	defaultContrailDatapathEncryptionPlay = "playbooks/deploy_and_run_all.yml"
	defaultInventoryTemplate              = "inventory.tmpl"
	defaultInventoryFile                  = "inventory.yml"

	defaultAppformixAnsibleRepoDir        = "/usr/share/contrail/"
	defaultAppformixAnsibleRepo           = "appformix-ansible-deployer"
	defaultAppformixProvPlay              = "playbooks/install_appformix.yml"
	defaultAppformixControllerNodeRes     = "appformix-controller-node"
	defaultAppformixControllerNodeResPath = "/" + defaultAppformixControllerNodeRes
	defaultAppformixBareHostNodeRes       = "appformix-bare-host-node"
	defaultAppformixBareHostNodeResPath   = "/" + defaultAppformixBareHostNodeRes
	defaultAppformixOpenstackNodeRes      = "appformix-openstack-node"
	defaultAppformixOpenstackNodeResPath  = "/" + defaultAppformixOpenstackNodeRes
	defaultAppformixComputeNodeRes        = "appformix-compute-node"
	defaultAppformixComputeNodeResPath    = "/" + defaultAppformixComputeNodeRes
	defaultXflowDir                       = "xflow/"

	defaultFilePermRWOnly = 0600

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
