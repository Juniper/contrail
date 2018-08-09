package cluster

// TODO:(ijohnson) generate constant resources
const (
	defaultResource                             = "contrail-cluster"
	defaultResourcePath                         = "/" + defaultResource
	defaultK8sResource                          = "kubernetes-cluster"
	defaultK8sResourcePath                      = "/" + defaultK8sResource
	defaultOpenstackResource                    = "openstack-cluster"
	defaultOpenstackResourcePath                = "/" + defaultOpenstackResource
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
	defaultContrailVrouterNodeRes               = "contrail-vrouter-node"
	defaultContrailVrouterNodeResPath           = "/" + defaultContrailVrouterNodeRes
	defaultContrailServiceNodeRes               = "contrail-service-node"
	defaultContrailServiceNodeResPath           = "/" + defaultContrailServiceNodeRes
	defaultKubernetesNodeRes                    = "kubernetes-node"
	defaultKubernetesNodeResPath                = "/" + defaultKubernetesNodeRes
	defaultKubernetesMasterNodeRes              = "kubernetes-master-node"
	defaultKubernetesMasterNodeResPath          = "/" + defaultKubernetesMasterNodeRes
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

	defaultWorkRoot           = "/var/tmp/contrail_cluster"
	defaultTemplateRoot       = "./pkg/cluster/configs"
	defaultInstanceTemplate   = "instances.tmpl"
	defaultInstanceFile       = "instances.yml"
	defaultProvisioner        = "ansible"
	defaultAnsibleRepo        = "contrail-ansible-deployer"
	defaultAnsibleRepoDir     = "/usr/share/contrail/"
	defaultContrailProvPlay   = "playbooks/install_contrail.yml"
	defaultInstanceProvPlay   = "playbooks/provision_instances.yml"
	defaultInstanceConfPlay   = "playbooks/configure_instances.yml"
	defaultOpenstackProvPlay  = "playbooks/install_openstack.yml"
	defaultKubernetesProvPlay = "playbooks/install_k8s.yml"

	// TODO (ijohnson): Fix LP#1756958 and remove the status constants
	statusField          = "provisioning_state"
	statusNoState        = "NOSTATE"
	statusCreated        = "CREATED"
	statusCreateProgress = "CREATE_IN_PROGRESS"
	statusCreateFailed   = "CREATE_FAILED"

	statusUpdated        = "UPDATED"
	statusUpdateProgress = "UPDATE_IN_PROGRESS"
	statusUpdateFailed   = "UPDATE_FAILED"
)
