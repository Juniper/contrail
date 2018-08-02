package db

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func parseFQName(fqNameStr string) (fqName []string) {
	json.Unmarshal([]byte(fqNameStr), &fqName)
	return fqName
}

func fqNameToString(fqName []string) (string, error) {
	fqNameStr, err := json.Marshal(fqName)
	if err != nil {
		return "", errors.Errorf("failed to parse fq name to string: %v", err)
	}
	return string(fqNameStr), nil
}
