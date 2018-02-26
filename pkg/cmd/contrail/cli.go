package contrail

import "github.com/Juniper/contrail/pkg/cmd/contrailcli"

func init() {
	Contrail.AddCommand(contrailcli.SchemaCmd)
	Contrail.AddCommand(contrailcli.ShowCmd)
	Contrail.AddCommand(contrailcli.ListCmd)
	Contrail.AddCommand(contrailcli.CreateCmd)
	Contrail.AddCommand(contrailcli.SetCmd)
	Contrail.AddCommand(contrailcli.UpdateCmd)
	Contrail.AddCommand(contrailcli.SyncCmd)
	Contrail.AddCommand(contrailcli.RmCmd)
	Contrail.AddCommand(contrailcli.DeleteCmd)
}
