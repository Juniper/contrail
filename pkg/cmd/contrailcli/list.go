package contrailcli

import (
	"fmt"
	"net/url"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	"github.com/Juniper/contrail/pkg/services"
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

const listHelpTemplate = `List command possible usages:
{% for schema in schemas %}contrail list {{ schema.ID }}
{% endfor %}`

func init() {
	ContrailCLI.AddCommand(ListCmd)

	ListCmd.Flags().StringVarP(&filters, services.FiltersKey, "f", "",
		"Comma-separated filter parameters (e.g. check==a,check==b,name==Bob)")
	ListCmd.Flags().IntVarP(&pageMarker, services.PageMarkerKey, "m", 0,
		"Page marker that returned resources start from (i.e. offset)")
	ListCmd.Flags().IntVarP(&pageLimit, services.PageLimitKey, "l", 100,
		"Limit number of returned resources")
	ListCmd.Flags().BoolVarP(&detail, services.DetailKey, "d", false,
		"Detailed data in response")
	ListCmd.Flags().BoolVar(&count, services.CountKey, false,
		"Return only resource count in response")
	ListCmd.Flags().BoolVarP(&shared, services.SharedKey, "s", false,
		"Include shared object in response")
	ListCmd.Flags().BoolVarP(&excludeHRefs, services.ExcludeHRefsKey, "e", false,
		"Exclude hrefs from response")
	ListCmd.Flags().StringVarP(&parentType, services.ParentTypeKey, "t", "",
		"Parent's type")
	ListCmd.Flags().StringVarP(&parentFQName, services.ParentFQNameKey, "n", "",
		"Colon-separated list of parents' fully-qualified names")
	ListCmd.Flags().StringVarP(&parentUUIDs, services.ParentUUIDsKey, "u", "",
		"Comma-separated list of parents' UUIDs")
	ListCmd.Flags().StringVar(&backrefUUIDs, services.BackrefUUIDsKey, "",
		"Comma-separated list of back references' UUIDs")
	ListCmd.Flags().StringVar(&objectUUIDs, services.ObjectUUIDsKey, "",
		"Comma-separated list of objects' UUIDs")
	ListCmd.Flags().StringVar(&fields, services.FieldsKey, "",
		"Comma-separated list of object fields returned in response")
}

// ListCmd defines list command.
var ListCmd = &cobra.Command{
	Use:   "list [SchemaID]",
	Short: "List data of specified resources",
	Long:  "Invoke command with empty SchemaID in order to show possible usages",
	Run: func(cmd *cobra.Command, args []string) {
		schemaID := ""
		if len(args) > 0 {
			schemaID = args[0]
		}
		output, err := listResources(schemaID)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(output)
	},
}

func queryParameters() url.Values {
	m := map[string]interface{}{
		services.FiltersKey:      filters,
		services.PageMarkerKey:   pageMarker,
		services.PageLimitKey:    pageLimit,
		services.DetailKey:       detail,
		services.CountKey:        count,
		services.SharedKey:       shared,
		services.ExcludeHRefsKey: excludeHRefs,
		services.ParentTypeKey:   parentType,
		services.ParentFQNameKey: parentFQName,
		services.ParentUUIDsKey:  parentUUIDs,
		services.BackrefUUIDsKey: objectUUIDs,
		services.FieldsKey:       fields,
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

func dashedCase(schemaID string) string {
	return strings.Replace(schemaID, "_", "-", -1)
}

func listResources(schemaID string) (string, error) {
	if schemaID == "" {
		return showHelp("", listHelpTemplate)
	}
	client, err := getClient()
	if err != nil {
		return "", nil
	}
	params := queryParameters()
	if schemaID == "" {
		//TODO
		return "", nil
	}
	//TODO support all schema
	events := &services.EventList{
		Events: []*services.Event{},
	}
	var response map[string][]interface{}
	_, err = client.Read(
		fmt.Sprintf("%s?%s", pluralPath(schemaID), params.Encode()), &response)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	for _, list := range response {
		for _, d := range list {
			m, _ := d.(map[string]interface{})
			events.Events = append(events.Events,
				services.NewEvent(&services.EventOption{
					Kind: schemaID,
					Data: m,
				}),
			)
		}
	}
	output, err := yaml.Marshal(events)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
