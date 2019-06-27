package integration_test

import (
	"fmt"
	"net/http"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/logutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCasesDoNotContainDuplicateUUIDsAcrossTests(t *testing.T) {
	uuidsToScenarios := duplicatedUUIDInTestScenarios(t)
	for uuid, scenarioNames := range uuidsToScenarios {
		uuidsToScenarios[uuid] = deduplicateSlice(scenarioNames)

		if len(uuidsToScenarios[uuid]) > 1 {
			assert.Fail(
				t,
				fmt.Sprintf("UUID %q duplicated in multiple test scenarios: %v", uuid, uuidsToScenarios[uuid]),
			)
		}
	}
}

func duplicatedUUIDInTestScenarios(t *testing.T) map[string][]string {
	log := logutil.NewLogger("duplicate-uuids-test")
	uuidsToScenarios := map[string][]string{}
	for _, ts := range testScenarios(t) {
		for _, task := range ts.Workflow {
			task.Request.Data = fileutil.YAMLtoJSONCompat(task.Request.Data)

			if task.Request.Method == http.MethodDelete || task.Request.Method == http.MethodGet {
				continue
			}

			resources, ok := task.Request.Data.(map[string]interface{})
			if !ok {
				log.WithField("data", task.Request.Data).Debug("Task's resource data contains no uuid - ignoring")
				continue
			}

			for _, resource := range resources {
				resourceMap, ok := resource.(map[string]interface{})
				if !ok {
					log.WithField("data", task.Request.Data).Debug("Task's resource data contains no uuid - ignoring")
					continue
				}

				var uuid interface{}
				if uuid, ok = resourceMap["uuid"]; !ok {
					log.WithField("data", task.Request.Data).Debug("Task's resource data contains no uuid - ignoring")
					continue
				}

				uuidS, ok := uuid.(string)
				require.True(t, ok, "failed to cast uuid of task %q in scenario %q", task.Name, ts.Name)

				uuidsToScenarios[uuidS] = append(uuidsToScenarios[uuidS], ts.Name)
			}
		}
	}
	return uuidsToScenarios
}

func testScenarios(t *testing.T) []*integration.TestScenario {
	var ts []*integration.TestScenario
	for _, f := range testFilePaths(t) {
		test, lErr := integration.LoadTest(f, nil)
		require.NoError(t, lErr)

		ts = append(ts, test)
	}
	return ts
}

func testFilePaths(t *testing.T) []string {
	var tfPaths []string
	for _, d := range []string{"../../apisrv/test_data/"} {
		files, err := fileutil.ListFiles(d)
		require.NoError(t, err)

		for i, name := range files {
			if !strings.HasSuffix(name, ".yml") && !strings.HasSuffix(name, ".yaml") {
				files = append(files[:i], files[i+1:]...)
			}
		}

		tfPaths = append(tfPaths, filePaths(d, files)...)
	}

	return tfPaths
}

func filePaths(directoryPath string, fileNames []string) (paths []string) {
	for _, n := range fileNames {
		paths = append(paths, filepath.Join(directoryPath, n))
	}
	return paths
}

// deduplicateSlice remove duplication from given string slice.
// See: https://github.com/golang/go/wiki/SliceTricks#in-place-deduplicate-comparable
func deduplicateSlice(s []string) []string {
	sort.Strings(s)
	j := 0
	for i := 1; i < len(s); i++ {
		if s[j] == s[i] {
			continue
		}
		j++
		// preserve the original data
		// in[i], in[j] = in[j], in[i]
		// only set what is required
		s[j] = s[i]
	}
	return s[:j+1]
}
