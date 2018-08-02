package db

import (
	"encoding/json"
)

func parseFQName(fqNameStr string) (fqName []string) {
	json.Unmarshal([]byte(fqNameStr), &fqName)
	return fqName
}

func fqNameToString(fqName []string) string {
	fqNameStr, _ := json.Marshal(fqName)
	return string(fqNameStr)
}
