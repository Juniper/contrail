package services

import (
	"fmt"
	"testing"

	"github.com/Juniper/contrail/pkg/common"
	"github.com/stretchr/testify/assert"
)

func TestReorderResourceList(t *testing.T) {
	var syncRequest ResourceList
	err := common.LoadFile("../../tools/init_data.yaml", &syncRequest)
	assert.NoError(t, err, "no error expected")
	for _, resource := range syncRequest.Resources {
		resource.Data = common.YAMLtoJSONCompat(resource.Data)
	}
	err = syncRequest.Sort()
	assert.NoError(t, err, "no error expected")
	fmt.Println(syncRequest)
	common.SaveFile("../../tools/init_data_sorted.yaml", syncRequest)
}
