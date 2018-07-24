package models

import (
	"fmt"
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

// DefaultNameForKind constructs the default name for an object of the given kind.
func DefaultNameForKind(kind string) string {
	return fmt.Sprintf("default-%s", kind)
}

// ChildFQName constructs fqName for child.
func ChildFQName(parentFQName []string, childName string) []string {
	result := make([]string, 0, len(parentFQName)+1)
	result = append(result, parentFQName...)
	if childName != "" {
		result = append(result, childName)
	}
	return result
}

// FQNameEquals checks if fqName slices have the same length and values
func FQNameEquals(fqNameA, fqNameB []string) bool {
	if len(fqNameA) != len(fqNameB) {
		return false
	}
	size := len(fqNameA)
	for i := 0; i < size; i++ {
		if fqNameA[i] != fqNameB[i] {
			return false
		}
	}
	return true
}
