package cluster

const (
	defaultResource               = "contrail-cluster"
	defaultResourcePath           = "/" + defaultResource
	defaultNodeRes                = "node"
	defaultNodeResPath            = "/" + defaultNodeRes
	defaultConfigNodeRes          = "contrail-config-node"
	defaultConfigNodeResPath      = "/" + defaultConfigNodeRes
	defaultConfigDBNodeRes        = "contrail-config-database-node"
	defaultConfigDBNodeResPath    = "/" + defaultConfigDBNodeRes
	defaultControlNodeRes         = "contrail-control-node"
	defaultControlNodeResPath     = "/" + defaultControlNodeRes
	defaultWebuiNodeRes           = "contrail-webui-node"
	defaultWebuiNodeResPath       = "/" + defaultWebuiNodeRes
	defaultAnalyticsNodeRes       = "contrail-analytics-node"
	defaultAnalyticsNodeResPath   = "/" + defaultAnalyticsNodeRes
	defaultAnalyticsDBNodeRes     = "contrail-analytics-database-node"
	defaultAnalyticsDBNodeResPath = "/" + defaultAnalyticsDBNodeRes
	defaultVrouterNodeRes         = "contrail-vrouter-node"
	defaultVrouterNodeResPath     = "/" + defaultVrouterNodeRes
	defaultWorkRoot               = "/var/tmp/contrail_cluster"
	defaultTemplateRoot           = "./pkg/cluster/configs"
	defaultInstanceTemplate       = "instances.tmpl"
	defaultInstanceFile           = "instances.yml"
	defaultProvisioner            = "ansible"
)
