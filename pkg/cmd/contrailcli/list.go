package contrailcli

import (
	"fmt"

	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/spf13/cobra"
)

var (
	filters      string
	pageMarker   string
	pageLimit    int
	detail       bool
	count        bool
	shared       bool
	excludeHRefs bool
	parentType   string
	parentFQName string
	parentUUIDs  string
	backrefUUIDs string
	objectUUIDs  string
	fields       string
)

func init() {
	ContrailCLI.AddCommand(listCmd)

	listCmd.Flags().StringVarP(&filters, baseservices.FiltersKey, "f", "",
		"Comma-separated filter parameters (e.g. check==a,check==b,name==Bob)")
	listCmd.Flags().StringVarP(&pageMarker, baseservices.PageMarkerKey, "m", "",
		"Page marker: return only the resources with UUIDs lexically greater than this value")
	listCmd.Flags().IntVarP(&pageLimit, baseservices.PageLimitKey, "l", 100,
		"Limit number of returned resources")
	listCmd.Flags().BoolVarP(&detail, baseservices.DetailKey, "d", false,
		"Detailed data in response")
	listCmd.Flags().BoolVar(&count, baseservices.CountKey, false,
		"Return only resource count in response")
	listCmd.Flags().BoolVarP(&shared, baseservices.SharedKey, "s", false,
		"Include shared object in response")
	listCmd.Flags().BoolVarP(&excludeHRefs, baseservices.ExcludeHRefsKey, "e", false,
		"Exclude hrefs from response")
	listCmd.Flags().StringVarP(&parentType, baseservices.ParentTypeKey, "t", "",
		"Parent's type")
	listCmd.Flags().StringVarP(&parentFQName, baseservices.ParentFQNameKey, "n", "",
		"Colon-separated list of parents' fully-qualified names")
	listCmd.Flags().StringVarP(&parentUUIDs, baseservices.ParentUUIDsKey, "u", "",
		"Comma-separated list of parents' UUIDs")
	listCmd.Flags().StringVar(&backrefUUIDs, baseservices.BackrefUUIDsKey, "",
		"Comma-separated list of back references' UUIDs")
	listCmd.Flags().StringVar(&objectUUIDs, baseservices.ObjectUUIDsKey, "",
		"Comma-separated list of objects' UUIDs")
	listCmd.Flags().StringVar(&fields, baseservices.FieldsKey, "",
		"Comma-separated list of object fields returned in response")
}

var listCmd = &cobra.Command{
	Use:   "list [SchemaID]",
	Short: "List data of specified resources",
	Long:  "Invoke command with empty SchemaID in order to show possible usages",
	Run: func(cmd *cobra.Command, args []string) {
		schemaID := ""
		if len(args) > 0 {
			schemaID = args[0]
		}

		cli, err := NewCLI()
		if err != nil {
			logutil.FatalWithStackTrace(err)
		}

		r, err := cli.ListResources(schemaID)
		if err != nil {
			logutil.FatalWithStackTrace(err)
		}

		fmt.Println(r)
	},
}
