package basemodels

// Reference is a generic reference instance.
type Reference interface {
	SetUUID(uuid string)
	SetTo(to []string)
	GetUUID() string
	GetTo() []string
	GetReferredKind() string
	GetAttribute() RefAttribute
	toMapper
}

type RefAttribute interface {
	toMapper
}

// References is wrapper type for reference slice.
type References []Reference

// Find returns first reference that fulfils the predicate.
func (r References) Find(pred func(Reference) bool) Reference {
	for _, ref := range r {
		if pred(ref) {
			return ref
		}
	}
	return nil
}

// Filter removes all the values that doesn't match the predicate.
func (r References) Filter(pred func(Reference) bool) References {
	result := make(References, 0, len(r))
	for _, ref := range r {
		if pred(ref) {
			result = append(result, ref)
		}
	}
	return result
}

// Unique returns references without duplicates.
func (r References) Unique() References {
	set := map[string]struct{}{}

	return r.Filter(func(ref Reference) bool {
		uuid := ref.GetUUID()
		if _, ok := set[uuid]; ok {
			return false
		}
		set[uuid] = struct{}{}
		return true
	})
}

// ForEach performs f on each element of references slice.
func ForEach(slice interface{}, f func(interface{})) bool {
	switch l := slice.(type) {
	case References:
		for _, x := range l {
			f(x)
		}
	case []Reference:
		for _, x := range l {
			f(x)
		}
	case []interface{}:
		for _, x := range l {
			f(x)
		}
	default:
		return false
	}
	return true
}

// NewReference creates new generic reference that can be processed into proper reference type.
func NewReference(uuid, kind string) Reference {
	return &genericReference{uuid: uuid, kind: kind}
}

type genericReference struct {
	uuid string
	to   []string
	kind string
}

func (g *genericReference) SetUUID(uuid string) {
	g.uuid = uuid
}

func (g *genericReference) SetTo(to []string) {
	g.to = to
}

func (g *genericReference) GetUUID() string {
	return g.uuid
}

func (g *genericReference) GetTo() []string {
	return g.to
}

func (g *genericReference) GetReferredKind() string {
	return g.kind
}

func (g *genericReference) GetAttribute() RefAttribute {
	return nil
}

func (g *genericReference) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"uuid": g.uuid,
		"to":   g.to,
	}
}
