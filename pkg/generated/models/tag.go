package models

// Tag

import "encoding/json"

// Tag
type Tag struct {
	DisplayName string         `json:"display_name"`
	UUID        string         `json:"uuid"`
	FQName      []string       `json:"fq_name"`
	TagID       U32BitHexInt   `json:"tag_id"`
	TagValue    string         `json:"tag_value"`
	Annotations *KeyValuePairs `json:"annotations"`
	Perms2      *PermType2     `json:"perms2"`
	ParentUUID  string         `json:"parent_uuid"`
	ParentType  string         `json:"parent_type"`
	IDPerms     *IdPermsType   `json:"id_perms"`
	TagTypeName string         `json:"tag_type_name"`

	TagTypeRefs []*TagTagTypeRef `json:"tag_type_refs"`
}

// TagTagTypeRef references each other
type TagTagTypeRef struct {
	UUID string   `json:"uuid"`
	To   []string `json:"to"` //FQDN

}

// String returns json representation of the object
func (model *Tag) String() string {
	b, _ := json.Marshal(model)
	return string(b)
}

// MakeTag makes Tag
func MakeTag() *Tag {
	return &Tag{
		//TODO(nati): Apply default
		DisplayName: "",
		UUID:        "",
		FQName:      []string{},
		TagID:       MakeU32BitHexInt(),
		TagValue:    "",
		Annotations: MakeKeyValuePairs(),
		Perms2:      MakePermType2(),
		ParentUUID:  "",
		ParentType:  "",
		IDPerms:     MakeIdPermsType(),
		TagTypeName: "",
	}
}

// InterfaceToTag makes Tag from interface
func InterfaceToTag(iData interface{}) *Tag {
	data := iData.(map[string]interface{})
	return &Tag{
		TagTypeName: data["tag_type_name"].(string),

		//{"description":"Tag type string representation","type":"string"}
		TagValue: data["tag_value"].(string),

		//{"description":"Tag value string representation","type":"string"}
		Annotations: InterfaceToKeyValuePairs(data["annotations"]),

		//{"type":"object","properties":{"key_value_pair":{"type":"array","item":{"type":"object","properties":{"key":{"type":"string"},"value":{"type":"string"}}}}}}
		Perms2: InterfaceToPermType2(data["perms2"]),

		//{"type":"object","properties":{"global_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7},"share":{"type":"array","item":{"type":"object","properties":{"tenant":{"type":"string"},"tenant_access":{"type":"integer","minimum":0,"maximum":7}}}}}}
		ParentUUID: data["parent_uuid"].(string),

		//{"type":"string"}
		ParentType: data["parent_type"].(string),

		//{"type":"string"}
		IDPerms: InterfaceToIdPermsType(data["id_perms"]),

		//{"type":"object","properties":{"created":{"type":"string"},"creator":{"type":"string"},"description":{"type":"string"},"enable":{"type":"boolean"},"last_modified":{"type":"string"},"permissions":{"type":"object","properties":{"group":{"type":"string"},"group_access":{"type":"integer","minimum":0,"maximum":7},"other_access":{"type":"integer","minimum":0,"maximum":7},"owner":{"type":"string"},"owner_access":{"type":"integer","minimum":0,"maximum":7}}},"user_visible":{"type":"boolean"}}}
		TagID: InterfaceToU32BitHexInt(data["tag_id"]),

		//{"description":"Internal Tag ID encapsulating tag type and value in                  hexadecimal fomat: 0xTTTTVVVV (T: type, V: value)","type":"string"}
		DisplayName: data["display_name"].(string),

		//{"type":"string"}
		UUID: data["uuid"].(string),

		//{"type":"string"}
		FQName: data["fq_name"].([]string),

		//{"type":"array","item":{"type":"string"}}

	}
}

// InterfaceToTagSlice makes a slice of Tag from interface
func InterfaceToTagSlice(data interface{}) []*Tag {
	list := data.([]interface{})
	result := MakeTagSlice()
	for _, item := range list {
		result = append(result, InterfaceToTag(item))
	}
	return result
}

// MakeTagSlice() makes a slice of Tag
func MakeTagSlice() []*Tag {
	return []*Tag{}
}
