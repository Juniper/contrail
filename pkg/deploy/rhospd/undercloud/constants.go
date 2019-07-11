package undercloud

const (
	defaultResource     = "rhospd-cloud-manager"
	defaultResourcePath = "/" + defaultResource

	defaultWorkRoot                   = "/var/tmp/rhospd_cloud_manager/"
	defaultSiteTemplate               = "site.tmpl"
	defaultSiteFile                   = "site.yml"
	defaultInventoryTemplate          = "inventory.tmpl"
	defaultInventoryFile              = "inventory.yml"
	defaultContrailCloudDir           = "/var/lib/contrail_cloud/"
	defaultContrailCloudScriptsDir    = defaultContrailCloudDir + "scripts/"
	defaultContrailCloudConfigDir     = defaultContrailCloudDir + "config/"
	defaultContrailCloudIntrospectDir = defaultContrailCloudDir + "introspection/"

	addKnownHostsCmd        = "ssh-keyscan -H localhost >> ~/.ssh/known_hosts"
	installContrailCloudCmd = "sudo " + defaultContrailCloudScriptsDir + "install_contrail_cloud_manager.sh"
	inventoryAssignCmd      = "sudo " + defaultContrailCloudScriptsDir + "inventory-assign.sh"

	statusNoState        = "NOSTATE"
	statusCreated        = "CREATED"
	statusCreateProgress = "CREATE_IN_PROGRESS"
	statusCreateFailed   = "CREATE_FAILED"

	statusUpdated        = "UPDATED"
	statusUpdateProgress = "UPDATE_IN_PROGRESS"
	statusUpdateFailed   = "UPDATE_FAILED"
)
