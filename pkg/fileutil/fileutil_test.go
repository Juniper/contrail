package fileutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyFile(t *testing.T) {
	srcTestData := "SRC TEST DATA"
	dstTestData := "DST TEST DATA"

	tests := []struct {
		name          string
		srcFilepath   string
		srcPerm       uint32
		dstFilepath   string
		dstPerm       uint32
		dstFileExists bool
		overwrite     bool
		fails         bool
	}{
		{
			name:        "Simple copy test",
			srcFilepath: "/var/tmp/contrail/copy_test_file",
			dstFilepath: "/var/tmp/contrail/copy_test_output_file",
			srcPerm:     0777,
		},
		{
			name:          "Disallow overwriting file",
			srcFilepath:   "/var/tmp/contrail/copy_test_file",
			dstFilepath:   "/var/tmp/contrail/copy_test_output_file",
			srcPerm:       0777,
			dstPerm:       0644,
			dstFileExists: true,
			fails:         true,
		},
		{
			name:          "Allow overwriting file",
			srcFilepath:   "/var/tmp/contrail/copy_test_file",
			dstFilepath:   "/var/tmp/contrail/copy_test_output_file",
			srcPerm:       0777,
			dstPerm:       0644,
			dstFileExists: true,
			overwrite:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			removeFilesIfExist(t, []string{tt.srcFilepath, tt.dstFilepath})
			defer removeFilesIfExist(t, []string{tt.srcFilepath, tt.dstFilepath})
			assert.NoError(t, ensureTestPathExist(tt.srcFilepath))
			assert.NoError(t, createFile(tt.srcFilepath, srcTestData, tt.srcPerm))
			assert.NoError(t, ensureTestPathExist(tt.dstFilepath))
			if tt.dstFileExists {
				createFile(tt.dstFilepath, dstTestData, tt.dstPerm)
			}
			assert.Equal(t, tt.fails, CopyFile(tt.srcFilepath, tt.dstFilepath, tt.overwrite) != nil)
			if tt.dstFileExists {
				filePermissionEqualTo(tt.dstFilepath, tt.dstPerm)
			} else {
				filePermissionEqualTo(tt.dstFilepath, tt.srcPerm)
			}
			if !tt.fails {
				dstDataEqualToSrcData, err := fileContains(tt.dstFilepath, srcTestData)
				assert.NoError(t, err)
				assert.True(t, dstDataEqualToSrcData)
			}
		})
	}
}

func ensureTestPathExist(file string) error {
	return os.MkdirAll(filepath.Base(file), 0644)
}

func createFile(path, data string, perm uint32) error {
	return WriteToFile(path, []byte(data), os.FileMode(perm))
}

func filePermissionEqualTo(filepath string, permission uint32) (bool, error) {
	stat, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}
	return uint32(stat.Mode().Perm()) == permission, nil
}

func fileContains(filepath, expectedData string) (bool, error) {
	actualData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return false, err
	}
	return string(actualData) == expectedData, nil
}

func removeFilesIfExist(t *testing.T, files []string) {
	for _, file := range files {
		if _, err := os.Stat(file); err != nil && os.IsNotExist(err) {
			continue
		}
		assert.NoError(t, os.Remove(file))
	}
}
