package basemodels

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Juniper/asf/pkg/format"
	"github.com/stretchr/testify/require"

	proto "github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
)

func TestUpdateData(t *testing.T) {
	codecs := []Codec{
		JSONCodec,
		ProtoCodec,
	}

	tests := []struct {
		name        string
		old, update Object
		fm          types.FieldMask
		want        Object
		fails       bool
	}{
		{name: "empty"},
		{
			name: "empty vn",
			old:  &stubObject{},
			want: &stubObject{},
		},
		{
			name:   "empty vn with empty update",
			old:    &stubObject{},
			update: &stubObject{},
			want:   &stubObject{},
		},
		{
			name:   "empty fieldmask",
			old:    &stubObject{UUID: "old-uuid", Name: "old-name"},
			update: &stubObject{UUID: "new-uuid"},
			want:   &stubObject{UUID: "old-uuid", Name: "old-name"},
		},
		{
			name:   "set UUID",
			old:    &stubObject{UUID: "old-uuid", Name: "old-name", DisplayName: "old-dn"},
			update: &stubObject{UUID: "new-uuid", DisplayName: "new-dn"},
			fm:     types.FieldMask{Paths: []string{"uuid"}},
			want:   &stubObject{UUID: "new-uuid", Name: "old-name", DisplayName: "old-dn"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, c := range codecs {
				t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
					var oldData []byte
					var err error
					if tt.old != nil {
						oldData, err = c.Encode(tt.old)
						require.NoError(t, err)
					}

					gotData, err := UpdateData(c, oldData, tt.update, tt.fm)
					if tt.fails {
						assert.Error(t, err)
					} else {
						assert.NoError(t, err)
					}

					var got proto.Message
					if tt.old != nil {
						got = proto.Clone(tt.old)
						err = c.Decode(gotData, got)
						require.NoError(t, err)
					}
					assert.Equal(t, tt.want, got)
				})
			}
		})
	}
}

type stubObject struct {
	UUID        string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty" yaml:"uuid,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"name,omitempty"`
	ParentUUID  string   `protobuf:"bytes,3,opt,name=parent_uuid,json=parentUuid,proto3" json:"parent_uuid,omitempty" yaml:"parent_uuid,omitempty"`
	ParentType  string   `protobuf:"bytes,4,opt,name=parent_type,json=parentType,proto3" json:"parent_type,omitempty" yaml:"parent_type,omitempty"`
	FQName      []string `protobuf:"bytes,5,rep,name=fq_name,json=fqName,proto3" json:"fq_name,omitempty" yaml:"fq_name,omitempty"`
	DisplayName string   `protobuf:"bytes,7,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty" yaml:"display_name,omitempty"`
}

func (o *stubObject) Reset() {
	*o = stubObject{}
}

func (o *stubObject) String() string {
	if o == nil {
		return "nil"
	}
	s := strings.Join([]string{`&stubObject{`,
		`UUID:` + fmt.Sprintf("%v", o.UUID) + `,`,
		`Name:` + fmt.Sprintf("%v", o.Name) + `,`,
		`ParentUUID:` + fmt.Sprintf("%v", o.ParentUUID) + `,`,
		`ParentType:` + fmt.Sprintf("%v", o.ParentType) + `,`,
		`FQName:` + fmt.Sprintf("%v", o.FQName) + `,`,
		`DisplayName:` + fmt.Sprintf("%v", o.DisplayName) + `,`,
		`}`,
	}, "")
	return s
}

func (o *stubObject) ProtoMessage() {
}

func (o *stubObject) GetUUID() string {
	if o != nil {
		return o.UUID
	}
	return ""
}

func (o *stubObject) SetUUID(uuid string) {
	o.UUID = uuid
}

func (o *stubObject) GetFQName() []string {
	if o != nil {
		return o.FQName
	}
	return nil
}

func (o *stubObject) GetParentUUID() string {
	if o != nil {
		return o.ParentUUID
	}
	return ""
}

func (o *stubObject) GetParentType() string {
	if o != nil {
		return o.ParentType
	}
	return ""
}

func (o *stubObject) Kind() string {
	return "stub-object"
}

func (o *stubObject) GetReferences() References {
	return nil
}

func (o *stubObject) GetTagReferences() References {
	return nil
}

func (o *stubObject) GetBackReferences() []Object {
	return nil
}

func (o *stubObject) GetChildren() []Object {
	return nil
}

func (o *stubObject) SetHref(string) {
}

func (o *stubObject) AddReference(interface{}) {
}

func (o *stubObject) AddBackReference(interface{}) {
}

func (o *stubObject) AddChild(interface{}) {
}

func (o *stubObject) RemoveReference(interface{}) {
}

func (o *stubObject) RemoveBackReference(interface{}) {
}

func (o *stubObject) RemoveChild(interface{}) {
}

func (o *stubObject) RemoveReferences() {
}

func (o *stubObject) ToMap() map[string]interface{} {
	if o == nil {
		return nil
	}
	return map[string]interface{}{
		"uuid":         o.UUID,
		"name":         o.Name,
		"parent_uuid":  o.ParentUUID,
		"parent_type":  o.ParentType,
		"fq_name":      o.FQName,
		"display_name": o.DisplayName,
	}
}

func (o *stubObject) ApplyMap(m map[string]interface{}) error {
	var err error
	if len(m) == 0 || o == nil {
		return nil
	}

	if val, ok := m["uuid"]; ok && val != nil {
		o.UUID, err = format.InterfaceToStringE(val)
	}
	if val, ok := m["name"]; ok && val != nil {
		o.Name, err = format.InterfaceToStringE(val)
	}
	if val, ok := m["parent_uuid"]; ok && val != nil {
		o.ParentUUID, err = format.InterfaceToStringE(val)
	}
	if val, ok := m["parent_type"]; ok && val != nil {
		o.ParentType, err = format.InterfaceToStringE(val)
	}
	if val, ok := m["fq_name"]; ok && val != nil {
		o.FQName, err = format.InterfaceToStringListE(val)
	}
	if val, ok := m["display_name"]; ok && val != nil {
		o.DisplayName, err = format.InterfaceToStringE(val)
	}
	return err
}

func (o *stubObject) ApplyPropCollectionUpdate(*PropCollectionUpdate) (updated map[string]interface{}, err error) {
	return nil, nil
}
