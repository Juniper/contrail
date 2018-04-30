package replication

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/pgtype"
	"github.com/kyleconroy/pgoutput"
)

type relationSet map[uint32]pgoutput.Relation

func (rs relationSet) Add(r pgoutput.Relation) {
	rs[r.ID] = r
}

func (rs relationSet) Get(id uint32) (pgoutput.Relation, error) {
	rel, ok := rs[id]
	if !ok {
		return rel, fmt.Errorf("no relation for %d", id)
	}
	return rel, nil
}

func getDecoder(col pgoutput.Column) pgoutput.DecoderValue {
	switch col.Type {
	case pgtype.JSONBOID:
		return &bytesDecoderValue{}
	case pgtype.JSONOID:
		return &bytesDecoderValue{}
	default:
		return col.Decoder()
	}
}

type bytesDecoderValue struct {
	bytes []byte
}

func (b *bytesDecoderValue) DecodeText(ci *pgtype.ConnInfo, src []byte) error {
	if src == nil {
		*b = bytesDecoderValue{}
		return nil
	}

	*b = bytesDecoderValue{bytes: src}
	return nil
}

func (b *bytesDecoderValue) Set(src interface{}) error {
	switch value := src.(type) {
	case []byte:
		*b = bytesDecoderValue{bytes: value}
	case string:
		*b = bytesDecoderValue{bytes: []byte(value)}
	default:
		buf, err := json.Marshal(value)
		if err != nil {
			return err
		}
		*b = bytesDecoderValue{bytes: buf}
	}
	return nil
}

func (b *bytesDecoderValue) Get() interface{} {
	return b.bytes
}

func (b *bytesDecoderValue) AssignTo(dst interface{}) error {
	switch value := dst.(type) {
	case *[]byte:
		*value = b.bytes
	case *string:
		*value = string(b.bytes)
	default:
		return errors.New("bad data type")
	}
	return nil
}
