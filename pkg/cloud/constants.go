package cloud

const (
	defaultCloudResource     = "cloud"
	defaultCloudResourcePath = "/" + defaultCloudResource

	statusField          = "provisioning_state"
	statusNoState        = "NOSTATE"
	statusCreated        = "CREATED"
	statusCreateProgress = "CREATE_IN_PROGRESS"
	statusCreateFailed   = "CREATE_FAILED"

	statusUpdated        = "UPDATED"
	statusUpdateProgress = "UPDATE_IN_PROGRESS"
	statusUpdateFailed   = "UPDATE_FAILED"

	statusDeleteProgress = "DELETE_IN_PROGRESS"
	statusDeleteFailed   = "DELETE_FAILED"

	createAction = "create"
	updateAction = "update"
	deleteAction = "delete"

	azTFJson  = "azure.tf.json"
	awsTFJson = "aws.tf.json"

	aws    = "aws"
	azure  = "azure"
	onPrem = "private"

	defaultRWOnlyPerm = 0600
)
