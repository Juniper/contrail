package cloud

const (
	defaultCloudResource     = "cloud"
	defaultCloudResourcePath = "/" + defaultCloudResource

	defaultWorkRoot                = "/var/tmp/cloud"
	defaultTemplateRoot            = "./pkg/cloud/configs"
	defaultGenTopoScript           = "transform/generate_topology.py"
	defaultGenInventoryScript      = "transform/generate_inventories.py"
	defaultTFStateFile             = "terraform.tfstate"
	defaultTopologyFile            = "topology.yml"
	defaultPublicCloudTopoTemplate = "public_cloud_topology.tmpl"
	defaultOnPremTopoTemplate      = "onprem_cloud_topology.tmpl"
	defaultSecretFile              = "secret.yml"
	defaultSecretTemplate          = "secret.tmpl"
	defaultMultiCloudDir           = "/usr/share/contrail/"
	defaultMultiCloudRepo          = "contrail-multi-cloud"

	statusField          = "provisioning_state"
	statusNoState        = "NOSTATE"
	statusCreated        = "CREATED"
	statusCreateProgress = "CREATE_IN_PROGRESS"
	statusCreateFailed   = "CREATE_FAILED"

	statusUpdated        = "UPDATED"
	statusUpdateProgress = "UPDATE_IN_PROGRESS"
	statusUpdateFailed   = "UPDATE_FAILED"

	createAction      = "create"
	updateAction      = "update"
	deleteAction      = "delete"
	deleteCloudAction = "DELETE_CLOUD"

	aws    = "aws"
	azure  = "azure"
	onPrem = "private"

	defaultRWOnlyPerm = 0600

	defaultSSHKeyRepo = "keypair"
	defaultSSHPvtKey  = "public_cloud_ssh_key"
	defaultSSHPubKey  = "public_cloud_ssh_key.pub"
)
