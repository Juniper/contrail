package basemodels

import (
	"fmt"
	"strings"
)

// CommonFieldPerms2 is a resource field that stores PermType2 data.
const CommonFieldPerms2 = "perms2"

const (
	//PermsNone for no permission
	PermsNone = iota
	//PermsX for exec permission
	PermsX
	//PermsW for write permission
	PermsW
	//PermsWX for exec and write permission
	PermsWX
	//PermsR if read permission
	PermsR
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

// FQNameToName gets object's name from it's fqName.
func FQNameToName(fqName []string) string {
	pos := len(fqName) - 1
	if pos < 0 {
		return ""
	}
	return fqName[pos]
}

// FQNameEquals checks if fqName slices have the same length and values
func FQNameEquals(fqNameA, fqNameB []string) bool {
	if len(fqNameA) != len(fqNameB) {
		return false
	}
	for i, v := range fqNameA {
		if v != fqNameB[i] {
			return false
		}
	}
	return true
}

// KindToSchemaID makes a snake_case schema ID from a kebab-case kind.
func KindToSchemaID(kind string) string {
	return strings.Replace(kind, "-", "_", -1)
}

// SchemaIDToKind makes a kebab-case kind from a snake_case schema ID.
func SchemaIDToKind(kind string) string {
	return strings.Replace(kind, "_", "-", -1)
}

// ReferenceKind constructs reference kind for given from and to kinds.
func ReferenceKind(fromKind, toKind string) string {
	return fmt.Sprintf("%s-%s", fromKind, toKind)
}
