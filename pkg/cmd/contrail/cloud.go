package contrail

import (
	"github.com/spf13/cobra"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/Juniper/asf/pkg/logutil/report"
	"github.com/Juniper/asf/pkg/osutil"
	"github.com/Juniper/contrail/pkg/cloud"
)

func init() {
	Contrail.AddCommand(cloudCmd)
}

var cloudCmd = &cobra.Command{
	Use:   "cloud",
	Short: "sub command cloud is used to manage public cloud infra",
	Long: `Cloud is a sub command used to manage
            public cloud infra. Currently
            supported infra are Azure`,
	Run: func(cmd *cobra.Command, args []string) {
		manageCloud()
	},
}

// TODO(apasyniuk) Export this into asf/osutils
type osCommandExecutor struct{}

func (e *osCommandExecutor) ExecCmdAndWait(
	r *report.Reporter, cmd string, args []string, dir string, envVars ...string,
) error {
	return osutil.ExecCmdAndWait(r, cmd, args, dir, envVars...)
}

func manageCloud() {
	manager, err := cloud.NewCloudManager(configFile, &osCommandExecutor{})
	if err != nil {
		logutil.FatalWithStackTrace(err)
	}
	if err = manager.Manage(); err != nil {
		logutil.FatalWithStackTrace(err)
	}
}
