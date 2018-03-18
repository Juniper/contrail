package cluster

const (
	defaultResource                       = "contrail-cluster"
	defaultResourcePath                   = "/" + defaultResource
	defaultNodeRes                        = "node"
	defaultNodeResPath                    = "/" + defaultNodeRes
	defaultConfigNodeRes                  = "contrail-config-node"
	defaultConfigNodeResPath              = "/" + defaultConfigNodeRes
	defaultConfigDBNodeRes                = "contrail-config-database-node"
	defaultConfigDBNodeResPath            = "/" + defaultConfigDBNodeRes
	defaultControlNodeRes                 = "contrail-control-node"
	defaultControlNodeResPath             = "/" + defaultControlNodeRes
	defaultWebuiNodeRes                   = "contrail-webui-node"
	defaultWebuiNodeResPath               = "/" + defaultWebuiNodeRes
	defaultAnalyticsNodeRes               = "contrail-analytics-node"
	defaultAnalyticsNodeResPath           = "/" + defaultAnalyticsNodeRes
	defaultAnalyticsDBNodeRes             = "contrail-analytics-database-node"
	defaultAnalyticsDBNodeResPath         = "/" + defaultAnalyticsDBNodeRes
	defaultVrouterNodeRes                 = "contrail-vrouter-node"
	defaultVrouterNodeResPath             = "/" + defaultVrouterNodeRes
	defaultKubernetesNodeRes              = "kubernetes-node"
	defaultKubernetesNodeResPath          = "/" + defaultKubernetesNodeRes
	defaultKubernetesMasterNodeRes        = "kubernetes-master-node"
	defaultKubernetesMasterNodeResPath    = "/" + defaultKubernetesMasterNodeRes
	defaultOpenstackControlNodeRes        = "openstack-control-node"
	defaultOpenstackControlNodeResPath    = "/" + defaultOpenstackControlNodeRes
	defaultOpenstackMonitoringNodeRes     = "openstack-monitoring-node"
	defaultOpenstackMonitoringNodeResPath = "/" + defaultOpenstackMonitoringNodeRes
	defaultOpenstackNetworkNodeRes        = "openstack-network-node"
	defaultOpenstackNetworkNodeResPath    = "/" + defaultOpenstackNetworkNodeRes
	defaultOpenstackStorageNodeRes        = "openstack-storage-node"
	defaultOpenstackStorageNodeResPath    = "/" + defaultOpenstackStorageNodeRes
	defaultOpenstackComputeNodeRes        = "openstack-compute-node"
	defaultOpenstackComputeNodeResPath    = "/" + defaultOpenstackComputeNodeRes

	defaultWorkRoot         = "/var/tmp/contrail_cluster"
	defaultTemplateRoot     = "./pkg/cluster/configs"
	defaultInstanceTemplate = "instances.tmpl"
	defaultInstanceFile     = "instances.yml"
	defaultProvisioner      = "ansible"
	defaultAnsibleRepo      = "contrail-ansible-deployer"
	defaultAnsibleRepoURL   = "https://github.com/Juniper/" + defaultAnsibleRepo + ".git"
	defaultClusterProvPlay  = "playbooks/install_contrail.yml"
	defaultInstanceProvPlay = "playbooks/provision_instances.yml"
	defaultInstanceConfPlay = "playbooks/configure_instances.yml"

	// TODO (ijohnson): Fix LP#1756958 and remove the status constants
	status_created         = "CREATED"
	status_create_progress = "CREATE_IN_PROGRESS"
	status_create_failed   = "CREATE_FAILED"

	status_updated         = "UPDATED"
	status_update_progress = "UPDATE_IN_PROGRESS"
	status_update_failed   = "UPDATE_FAILED"
)
