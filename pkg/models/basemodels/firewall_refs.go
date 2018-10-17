package basemodels

import (
	"strings"

	"github.com/Juniper/contrail/pkg/common"
)

// CheckAssociatedRefsInSameScope checks scope of firewall resources refs
// that global scoped firewall resource cannot reference scoped resources
func CheckAssociatedRefsInSameScope(
	obj Object, fqName []string, refs []Reference,
) error {

	// this method is simply based on the fact global
	// firewall resource (draft or not) have a FQ name length equal to two
	// and scoped one (draft or not) have a FQ name longer than 2 to
	// distinguish a scoped firewall resource to a global one. If that
	// assumption disappear, all that method will need to be re-worked
	if len(fqName) != 2 {
		return nil
	}

	for _, ref := range refs {
		if err := checkRefInSameScope(obj, ref, fqName); err != nil {
			return err
		}
	}

	return nil
}

func checkRefInSameScope(
	obj Object, ref Reference, fqName []string,
) error {
	if len(ref.GetTo()) == 2 {
		return nil
	}

	kind := strings.Replace(obj.Kind(), "-", " ", -1)
	refKind := strings.Replace(ref.GetReferredKind(), "-", " ", -1)
	return common.ErrorBadRequestf(
		"global %s %s (%s) cannot reference a scoped %s %s (%s)",
		strings.Title(kind),
		FQNameToString(fqName),
		obj.GetUUID(),
		strings.Title(refKind),
		FQNameToString(ref.GetTo()),
		ref.GetUUID(),
	)
}
