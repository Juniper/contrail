package contrailcli

import (
	"fmt"
	"net/url"

	"github.com/Juniper/contrail/pkg/common"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	filters      string
	pageMarker   string
	pageLimit    string
	detail       string
	count        string
	shared       string
	excludeHRefs string
	parentType   string
	parentFQName string
	parentUUIDs  string
	backrefUUIDs string
	objectUUIDs  string
	fields       string
)

func init() {
	ContrailCLI.AddCommand(ListCmd)

	ListCmd.Flags().StringVarP(&filters, common.FiltersKey, "f", "",
		"Comma-separated filter parameters (e.g. check==a,check==b,name==Bob)")
	ListCmd.Flags().StringVarP(&pageMarker, common.PageMarkerKey, "m", "",
		"Page marker that returned resources start from (i.e. offset) [integer]")
	ListCmd.Flags().StringVarP(&pageLimit, common.PageLimitKey, "l", "",
		"Limit number of returned resources [integer]")
	ListCmd.Flags().StringVarP(&detail, common.DetailKey, "d", "",
		"Detailed data in response [bool]")
	ListCmd.Flags().StringVar(&count, common.CountKey, "",
		"Return only resource count in response [bool]")
	ListCmd.Flags().StringVarP(&shared, common.SharedKey, "s", "",
		"Include shared object in response [bool]")
	ListCmd.Flags().StringVarP(&excludeHRefs, common.ExcludeHRefsKey, "e", "",
		"Exclude HRefs from response [bool]")
	ListCmd.Flags().StringVarP(&parentType, common.ParentTypeKey, "t", "",
		"Parent's type")
	ListCmd.Flags().StringVarP(&parentFQName, common.ParentFQNameKey, "n", "",
		"Colon-separated list of parents' fully-qualified names")
	ListCmd.Flags().StringVarP(&parentUUIDs, common.ParentUUIDsKey, "u", "",
		"Comma-separated list of parents' UUIDs")
	ListCmd.Flags().StringVar(&backrefUUIDs, common.BackrefUUIDsKey, "",
		"Comma-separated list of back references' UUIDs")
	ListCmd.Flags().StringVar(&objectUUIDs, common.ObjectUUIDsKey, "",
		"Comma-separated list of objects' UUIDs")
	ListCmd.Flags().StringVar(&fields, common.FieldsKey, "",
		"Comma-separated list of object fields returned in response")
}

// ListCmd defines list command.
var ListCmd = &cobra.Command{
	Use:   "list [SchemaID]",
	Short: "List resources data",
	Long:  "Invoke command with empty SchemaID in order to show available schemas",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		listResources(args)
	},
}

func listResources(args []string) {
	a, err := getAuthenticatedAgent(configFile)
	if err != nil {
		log.Fatal(err)
	}

	output, err := a.ListCLI(args[0], queryParameters())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(output)
}

func queryParameters() url.Values {
	m := map[string]string{
		common.FiltersKey:      filters,
		common.PageMarkerKey:   pageMarker,
		common.PageLimitKey:    pageLimit,
		common.DetailKey:       detail,
		common.CountKey:        count,
		common.SharedKey:       shared,
		common.ExcludeHRefsKey: excludeHRefs,
		common.ParentTypeKey:   parentType,
		common.ParentFQNameKey: parentFQName,
		common.ParentUUIDsKey:  parentUUIDs,
		common.BackrefUUIDsKey: objectUUIDs,
		common.FieldsKey:       fields,
	}

	values := url.Values{}
	for k, v := range m {
		if v != "" {
			values.Set(k, v)
		}
	}
	return values
}
