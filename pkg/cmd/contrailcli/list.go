package contrailcli

import (
	"fmt"
	"net/url"

	"github.com/Juniper/contrail/pkg/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	filters      string
	pageMarker   int
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
	ContrailCLI.AddCommand(ListCmd)

	ListCmd.Flags().StringVarP(&filters, models.FiltersKey, "f", "",
		"Comma-separated filter parameters (e.g. check==a,check==b,name==Bob)")
	ListCmd.Flags().IntVarP(&pageMarker, models.PageMarkerKey, "m", 0,
		"Page marker that returned resources start from (i.e. offset)")
	ListCmd.Flags().IntVarP(&pageLimit, models.PageLimitKey, "l", 100,
		"Limit number of returned resources")
	ListCmd.Flags().BoolVarP(&detail, models.DetailKey, "d", false,
		"Detailed data in response")
	ListCmd.Flags().BoolVar(&count, models.CountKey, false,
		"Return only resource count in response")
	ListCmd.Flags().BoolVarP(&shared, models.SharedKey, "s", false,
		"Include shared object in response")
	ListCmd.Flags().BoolVarP(&excludeHRefs, models.ExcludeHRefsKey, "e", false,
		"Exclude hrefs from response")
	ListCmd.Flags().StringVarP(&parentType, models.ParentTypeKey, "t", "",
		"Parent's type")
	ListCmd.Flags().StringVarP(&parentFQName, models.ParentFQNameKey, "n", "",
		"Colon-separated list of parents' fully-qualified names")
	ListCmd.Flags().StringVarP(&parentUUIDs, models.ParentUUIDsKey, "u", "",
		"Comma-separated list of parents' UUIDs")
	ListCmd.Flags().StringVar(&backrefUUIDs, models.BackrefUUIDsKey, "",
		"Comma-separated list of back references' UUIDs")
	ListCmd.Flags().StringVar(&objectUUIDs, models.ObjectUUIDsKey, "",
		"Comma-separated list of objects' UUIDs")
	ListCmd.Flags().StringVar(&fields, models.FieldsKey, "",
		"Comma-separated list of object fields returned in response")
}

// ListCmd defines list command.
var ListCmd = &cobra.Command{
	Use:   "list [SchemaID]",
	Short: "List data of specified resources",
	Long:  "Invoke command with empty SchemaID in order to show possible usages",
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
	m := map[string]interface{}{
		models.FiltersKey:      filters,
		models.PageMarkerKey:   pageMarker,
		models.PageLimitKey:    pageLimit,
		models.DetailKey:       detail,
		models.CountKey:        count,
		models.SharedKey:       shared,
		models.ExcludeHRefsKey: excludeHRefs,
		models.ParentTypeKey:   parentType,
		models.ParentFQNameKey: parentFQName,
		models.ParentUUIDsKey:  parentUUIDs,
		models.BackrefUUIDsKey: objectUUIDs,
		models.FieldsKey:       fields,
	}

	values := url.Values{}
	for k, v := range m {
		value := fmt.Sprint(v)
		if !isZeroValue(value) {
			values.Set(k, value)
		}
	}
	return values
}

func isZeroValue(value interface{}) bool {
	return value == "" || value == 0 || value == false
}
