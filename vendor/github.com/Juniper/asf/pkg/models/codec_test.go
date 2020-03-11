package models

import (
	"fmt"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
