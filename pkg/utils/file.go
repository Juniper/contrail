package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
)

//SaveFile saves object to file. suffix of filepath will be
// used as file type. currently, json and yaml is supported
func SaveFile(file string, data interface{}) error {
	var bytes []byte
	var err error
	if strings.HasSuffix(file, ".json") {
		bytes, err = json.MarshalIndent(data, "", "    ")
	} else if strings.HasSuffix(file, ".yaml") || strings.HasSuffix(file, ".yml") {
		bytes, err = yaml.Marshal(data)

	}
	if err != nil {
		return err
	}
	return ioutil.WriteFile(file, bytes, os.ModePerm)
}

//LoadFile loads object from file. suffix of filepath will be
// used as file type. currently, json and yaml is supported
func LoadFile(filePath string, data interface{}) error {
	bodyBuff, err := GetContent(filePath)
	if err != nil {
		return err
	}
	if strings.HasSuffix(filePath, ".json") {
		return json.Unmarshal(bodyBuff, data)
	} else if strings.HasSuffix(filePath, ".yaml") || strings.HasSuffix(filePath, ".yml") {
		return yaml.Unmarshal(bodyBuff, data)
	}
	return fmt.Errorf("format isn't supported")
}

//GetContent loads file from remote or local
func GetContent(url string) ([]byte, error) {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return content, err
	}
	if strings.HasPrefix(url, "file://") {
		url = strings.TrimPrefix(url, "file://")
	}
	content, err := ioutil.ReadFile(url)
	return content, err
}

// TempFile creates a temporary file with the specified prefix and suffix
func TempFile(dir string, prefix string, suffix string) (*os.File, error) {
	if dir == "" {
		dir = os.TempDir()
	}

	name := filepath.Join(dir, fmt.Sprint(prefix, time.Now().UnixNano(), suffix))
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
}
