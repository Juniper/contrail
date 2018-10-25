package basemodels

// Reference is a generic reference instance.
type Reference interface {
	SetUUID(uuid string)
	SetTo(to []string)
	GetUUID() string
	GetTo() []string
	GetReferredKind() string
}

// References is wrapper type for reference slice.
type References []Reference

// Find returns first reference that fulfils the predicate
func (r References) Find(pred func(Reference) bool) Reference {
	for _, ref := range r {
		if pred(ref) {
			return ref
		}
	}
	return nil
}
