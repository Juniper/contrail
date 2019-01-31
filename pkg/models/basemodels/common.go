package basemodels

import (
	"fmt"
	"strings"
	"time"
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

const (
	vncTimeLenghtWithoutMs = len("yyyy-mm-ddThh:mm:ss")
	vncTimeLenghtWithMs    = len("yyyy-mm-ddThh:mm:ss.mmmmmm")
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

// ReferenceKind constructs reference kind for given from and to kinds.
func ReferenceKind(fromKind, toKind string) string {
	return fmt.Sprintf("%s-%s", fromKind, toKind)
}

// ToVNCTime returns time string in VNC format.
func ToVNCTime(t time.Time) string {
	if t.Nanosecond() < 1000 {
		return t.UTC().Format(time.RFC3339)[0:vncTimeLenghtWithoutMs]
	} else {
		date := t.UTC().Format(time.RFC3339Nano)
		// RGC3339Nano contains Z letter at the end, we need to get rid of it
		if date = date[0 : len(date)-1]; len(date) >= vncTimeLenghtWithMs {
			return date[0:vncTimeLenghtWithMs]
		} else {
			return date + additionalZeros(vncTimeLenghtWithMs-len(date))
		}
	}
}

func additionalZeros(n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += "0"
	}
	return result
}
