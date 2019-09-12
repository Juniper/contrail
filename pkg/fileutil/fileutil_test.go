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
		srcPerm       os.FileMode
		dstFilepath   string
		dstPerm       os.FileMode
		dstFileExists bool
		overwrite     bool
		fails         bool
	}{
		{
			name:        "Simple copy test",
			srcFilepath: "/var/tmp/contrail/fileutiltest/copy_test_file",
			dstFilepath: "/var/tmp/contrail/fileutiltest/copy_test_output_file",
			srcPerm:     0700,
		},
		{
			name:          "Disallow overwriting file",
			srcFilepath:   "/var/tmp/contrail/fileutiltest/copy_test_file",
			dstFilepath:   "/var/tmp/contrail/fileutiltest/copy_test_output_file",
			srcPerm:       0777,
			dstPerm:       0644,
			dstFileExists: true,
			fails:         true,
		},
		{
			name:          "Allow overwriting file",
			srcFilepath:   "/var/tmp/contrail/fileutiltest/copy_test_file",
			dstFilepath:   "/var/tmp/contrail/fileutiltest/copy_test_output_file",
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

			var expectedPermissions os.FileMode
			if tt.dstFileExists {
				assert.NoError(t, createFile(tt.dstFilepath, dstTestData, tt.dstPerm))
				expectedPermissions = tt.dstPerm
			} else {
				expectedPermissions = resolveActualSrcPermissions(t, tt.srcFilepath)
			}

			if tt.fails {
				assert.Error(t, CopyFile(tt.srcFilepath, tt.dstFilepath, tt.overwrite))
				return
			}

			assert.NoError(t, CopyFile(tt.srcFilepath, tt.dstFilepath, tt.overwrite))
			permissionsEqual, err := filePermissionEqualTo(tt.dstFilepath, expectedPermissions)
			assert.NoError(t, err)
			assert.True(t, permissionsEqual)
			dstDataEqualToSrcData, err := fileContains(tt.dstFilepath, srcTestData)
			assert.NoError(t, err)
			assert.True(t, dstDataEqualToSrcData)
		})
	}
}

func ensureTestPathExist(file string) error {
	return os.MkdirAll(filepath.Base(file), 0644)
}

func createFile(path, data string, perm os.FileMode) error {
	return ioutil.WriteFile(path, []byte(data), os.FileMode(perm))
}

func filePermissionEqualTo(filepath string, permission os.FileMode) (bool, error) {
	stat, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}
	return stat.Mode().Perm() == permission, nil
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

func resolveActualSrcPermissions(t *testing.T, srcFile string) os.FileMode {
	info, err := os.Stat(srcFile)
	assert.NoError(t, err, "could not resolve permissions of created source file")
	return info.Mode().Perm()
}
