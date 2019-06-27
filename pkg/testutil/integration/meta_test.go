package integration_test

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/Juniper/contrail/pkg/fileutil"
	"github.com/Juniper/contrail/pkg/testutil/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCasesDoNotContainDuplicateUUIDsAcrossTests(t *testing.T) {
	uuidsToScenarios := duplicatedUUIDsAcrossTestScenarios(t)

	for uuid, scenarios := range uuidsToScenarios {
		if len(scenarios) > 1 {
			assert.Fail(t, fmt.Sprintf("UUID %q duplicated in multiple test scenarios: %v", uuid, scenarios))
		}
	}
}

func duplicatedUUIDsAcrossTestScenarios(t *testing.T) map[string][]string {
	uuidsToScenarios := duplicatedUUIDs(t)
	for uuid, scenarioNames := range uuidsToScenarios {
		uuidsToScenarios[uuid] = deduplicateSlice(scenarioNames)
	}

	return uuidsToScenarios
}

func duplicatedUUIDs(t *testing.T) map[string][]string {
	uuidsToScenarios := map[string][]string{}
	for _, ts := range testScenarios(t) {
		for _, task := range ts.Workflow {
			task.Request.Data = fileutil.YAMLtoJSONCompat(task.Request.Data)

			if task.Request.Method == http.MethodDelete || task.Request.Method == http.MethodGet {
				continue
			}

			resources, ok := task.Request.Data.(map[string]interface{})
			if !ok {
				continue
			}

			for _, resource := range resources {
				resourceMap, ok := resource.(map[string]interface{})
				if !ok {
					continue
				}

				var uuid interface{}
				if uuid, ok = resourceMap["uuid"]; !ok {
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
	for _, d := range lookupTestDataDirectories(t, []string{}, "../..") {
		files, err := fileutil.ListFiles(d)
		require.NoError(t, err)

		var fileNames []string
		for _, f := range files {
			if isTestFile(f) {
				fileNames = append(fileNames, f.Name())
			}
		}

		tfPaths = append(tfPaths, filePaths(d, fileNames)...)
	}

	return tfPaths
}

func lookupTestDataDirectories(t *testing.T, tdDirectories []string, path string) []string {
	files, err := fileutil.ListFiles(path)
	require.NoError(t, err)

	for _, f := range files {
		if isTestDataDirectory(f) {
			tdDirectories = append(tdDirectories, filepath.Join(path, f.Name()))
		} else if f.IsDir() {
			tdDirectories = lookupTestDataDirectories(t, tdDirectories, filepath.Join(path, f.Name()))
		}
	}

	return tdDirectories
}

func isTestDataDirectory(f os.FileInfo) bool {
	if _, ok := map[string]struct{}{
		"testdata":  {},
		"test_data": {},
		"tests":     {},
	}[f.Name()]; ok && f.IsDir() {
		return true
	}

	return false
}

func isTestFile(f os.FileInfo) bool {
	return !f.IsDir() &&
		strings.HasPrefix(f.Name(), "test_") &&
		(strings.HasSuffix(f.Name(), ".yml") || strings.HasSuffix(f.Name(), ".yaml")) &&
		f.Name() != "test_with_refs.yaml"
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
