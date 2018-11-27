package baseservices

// SimpleListSpec is a list spec for a common case where List is filtered only by uuids.
// Whole resources are fetched if no fields are specified
func SimpleListSpec(uuids []string, fields ...string) *ListSpec {
	return &ListSpec{
		Filters: []*Filter{{
			Key:    "uuid",
			Values: uuids,
		}},
		Detail: true,
		Fields: fields,
	}
}

// ToGetSpec returns corresponding GetSpec associated with ListSpec.
func (m *ListSpec) ToGetSpec() *GetSpec {
	return &GetSpec{
		Fields:          m.Fields,
		ExcludeChildren: m.ExcludeChildren,
		ExcludeBackRefs: m.ExcludeBackRefs,
	}
}
