package cloud

// Cloud related constants
const (
	defaultCloudResource     = "cloud"
	defaultCloudResourcePath = "/" + defaultCloudResource

	defaultWorkRoot                = "/var/tmp/cloud"
	defaultTemplateRoot            = "./pkg/cloud/configs"
	defaultGenInventoryScript      = "transform/generate_inventories.py"
	defaultAWSPlanTF               = "aws.tf.json"
	defaultAzurePlanTF             = "azure.tf.json"
	defaultGCPPlanTF               = "google.tf.json"
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
	deleteCloudAction = "DELETE_CLOUD"

	aws    = "aws"
	azure  = "azure"
	gcp    = "gcp"
	onPrem = "private"

	defaultRWOnlyPerm = 0600
	defaultSSHKeyRepo = "keypair"

	pubSSHKey     = "PUBLIC_SSH_KEY"
	privateSSHKey = "PRIVATE_SSH_KEY"
)
