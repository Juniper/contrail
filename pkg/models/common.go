package models

import (
	fmt "fmt"
	"strings"
)

const (
	//PermsNone for no permission
	PermsNone = iota
	//PermsX for exec permission
	PermsX
	//PermsW for write permission
	PermsW
	//PermsR if read permission
	PermsR
	//PermsWX for exec and write permission
	PermsWX
	//PermsRX for exec and read permission
	PermsRX
	//PermsRW for read and write permission
	PermsRW
	//PermsRWX for all permission
	PermsRWX
)

// ParseFQName parse string representation of FQName.
func ParseFQName(fqNameString string) []string {
	if fqNameString == "" {
		return nil
	}
	return strings.Split(fqNameString, ":")
}

// FQNameToString returns string representation of FQName.
func FQNameToString(fqName []string) string {
	return strings.Join(fqName, ":")
}

// DefaultNameForKind constructs default name for object of given kind.
func DefaultNameForKind(kind string) string {
	return fmt.Sprintf("default-%s", kind)
}
