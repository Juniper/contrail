package replication

import (
	"testing"

	"github.com/jackc/pgx/pgtype"
	"github.com/kyleconroy/pgoutput"
	"github.com/stretchr/testify/assert"
)

func TestRelationDecoderDecodeRelationData(t *testing.T) {
	exampleRelation := relation{
		Name: "test-resource",
		Columns: []pgoutput.Column{
			{Name: "string-property", Key: true, Type: pgtype.VarcharOID},
			{Name: "int-property", Type: pgtype.Int4OID},
			{Name: "float-property", Type: pgtype.Float8OID},
		},
	}

	exampleRow := []pgoutput.Tuple{
		{Value: []byte(`foo`)},
		{Value: []byte(`1337`)},
		{Value: []byte(`1.337`)},
	}

	exampleRowData := map[string]interface{}{
		"string-property": "foo",
		"int-property":    int32(1337),
		"float-property":  1.337,
	}

	examplePK := []string{"foo"}

	tests := []struct {
		row          []pgoutput.Tuple
		expectedPK   []string
		name         string
		expectedKind string
		decoder      relationDecoder
		expectedData map[string]interface{}
		relationID   uint32
		fails        bool
	}{{
		name: "empty row and relationID", fails: true,
	}, {
		name: "non existent relation", relationID: 1, fails: true,
	}, {
		name: "empty row", decoder: relationDecoder{1: exampleRelation}, relationID: 1, fails: true,
	}, {
		name:       "no primary key",
		decoder:    relationDecoder{1: relation{Name: "rel"}},
		relationID: 1,
		fails:      true,
	}, {
		name:         "correct message",
		decoder:      relationDecoder{1: exampleRelation},
		relationID:   1,
		row:          exampleRow,
		expectedKind: exampleRelation.Name,
		expectedData: exampleRowData,
		expectedPK:   examplePK,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// when
			kind, pk, data, err := tt.decoder.DecodeRelationData(tt.relationID, tt.row)

			// then
			if tt.fails {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedKind, kind)
			assert.Equal(t, tt.expectedPK, pk)
			assert.Equal(t, tt.expectedData, data)
		})
	}
}

func TestRelationDecoderAddRelation(t *testing.T) {
	tests := []struct {
		name     string
		decoder  relationDecoder
		relation pgoutput.Relation
		expected relationDecoder
	}{{
		name:     "empty",
		decoder:  relationDecoder{},
		expected: relationDecoder{0: relation{}},
	}, {
		name:     "new relation",
		decoder:  relationDecoder{},
		relation: pgoutput.Relation{ID: 1337},
		expected: relationDecoder{1337: relation{ID: 1337}},
	}, {
		name:     "already stored relation",
		decoder:  relationDecoder{1337: relation{Name: "old"}},
		relation: pgoutput.Relation{ID: 1337, Name: "new"},
		expected: relationDecoder{1337: relation{ID: 1337, Name: "new"}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.decoder.AddRelation(tt.relation)
			assert.Equal(t, tt.expected, tt.decoder)
		})
	}
}
