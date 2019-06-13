package contrailcli

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/services"
	"github.com/Juniper/contrail/pkg/services/baseservices"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"
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

const listHelpTemplate = `List command possible usages:
{% for schema in schemas %}contrail list {{ schema.ID }}
{% endfor %}`

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
		o, err := listResources(schemaID)
		if err != nil {
			logutil.FatalWithStackTrace(err)
		}
		fmt.Println(o)
	},
}

func listResources(schemaID string) (string, error) {
	if schemaID == "" {
		return showHelp("", listHelpTemplate)
	}
	client, err := newHTTPClient()
	if err != nil {
		return "", err
	}

	var response map[string][]interface{}
	if _, err = client.ReadWithQuery(
		context.Background(),
		pluralPath(schemaID),
		queryParameters(),
		&response,
	); err != nil {
		return "", err
	}

	var events services.EventList
	for _, list := range response {
		for _, d := range list {
			m, ok := d.(map[string]interface{})
			if !ok {
				return "", errors.New("type assertion on response value failed")
			}

			event, neErr := services.NewEvent(services.EventOption{
				Kind: schemaID,
				Data: m,
			})
			if neErr != nil {
				logrus.Errorf("failed to create event - skipping: %v", neErr)
				continue
			}

			events.Events = append(events.Events, event)
		}
	}

	output, err := yaml.Marshal(&events)
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func queryParameters() url.Values {
	m := map[string]interface{}{
		baseservices.FiltersKey:      filters,
		baseservices.PageMarkerKey:   pageMarker,
		baseservices.PageLimitKey:    pageLimit,
		baseservices.DetailKey:       detail,
		baseservices.CountKey:        count,
		baseservices.SharedKey:       shared,
		baseservices.ExcludeHRefsKey: excludeHRefs,
		baseservices.ParentTypeKey:   parentType,
		baseservices.ParentFQNameKey: parentFQName,
		baseservices.ParentUUIDsKey:  parentUUIDs,
		baseservices.BackrefUUIDsKey: objectUUIDs,
		baseservices.FieldsKey:       fields,
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
