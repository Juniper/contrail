package models

import (
	"fmt"
	"strings"

	"github.com/Juniper/contrail/pkg/models/basemodels"
)

func TagTypeValueFromFQName(fqName []string) (tagType, tagValue string) {
	return TagTypeValueFromName(basemodels.FQNameToName(fqName))
}

func TagTypeValueFromName(name string) (tagType, tagValue string) {
	splits := strings.Split(name, "=")
	if len(splits) < 2 {
		return "", ""
	}
	return splits[0], splits[1]
}

func CreateTagName(tagType, tagValue string) string {
	return fmt.Sprintf("%v=%v", tagType, tagValue)
}

func GroupTagRefsByType(tagRefs []basemodels.Reference) map[string][]basemodels.Reference {
	refsPerType := make(map[string][]basemodels.Reference)
	for _, tagRef := range tagRefs {
		tagRefType, _ := TagTypeValueFromFQName(tagRef.GetTo())

		refsPerType[tagRefType] = append(refsPerType[tagRefType], tagRef)
	}
	return refsPerType
}

func GroupTagRefsByValue(tagRefs []basemodels.Reference) map[string]basemodels.Reference {
	refsPerValue := make(map[string]basemodels.Reference)

	for _, tagRef := range tagRefs {
		_, tagRefValue := TagTypeValueFromFQName(tagRef.GetTo())
		refsPerValue[tagRefValue] = tagRef
	}
	return refsPerValue
}
