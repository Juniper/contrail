package db

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func parseFQName(fqNameStr string) ([]string, error) {
	var fqName []string
	err := json.Unmarshal([]byte(fqNameStr), &fqName)
	if err != nil {
		return nil, errors.Errorf("failed to parse fq name from string: %v", err)
	}
	return fqName, nil
}

func fqNameToString(fqName []string) (string, error) {
	fqNameStr, err := json.Marshal(fqName)
	if err != nil {
		return "", errors.Errorf("failed to parse fq name to string: %v", err)
	}
	return string(fqNameStr), nil
}
