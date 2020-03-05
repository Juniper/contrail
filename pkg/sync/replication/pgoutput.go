package replication

import (
	"fmt"

	"github.com/Juniper/asf/pkg/logutil"
	"github.com/kyleconroy/pgoutput"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type pgoutputDecoder struct {
	decoder       relationDecoder
	txnInProgress []Change

	log *logrus.Entry
}

func newPgoutputDecoder() *pgoutputDecoder {
	return &pgoutputDecoder{
		decoder: relationDecoder{},
		log:     logutil.NewLogger("pgoutput-writer"),
	}
}

func (p *pgoutputDecoder) DecodeChanges(d []byte) ([]Change, error) {
	msg, err := pgoutput.Parse(d)
	if err != nil {
		return nil, errors.Wrap(err, "invalid pgoutput message")
	}

	switch v := msg.(type) {
	case pgoutput.Relation:
		p.log.Debug("received RELATION message")
		p.decoder.AddRelation(v)
	case pgoutput.Begin:
		p.log.Debug("received BEGIN message")
		p.txnInProgress = []Change{}
	case pgoutput.Commit:
		p.log.Debug("received COMMIT message")
		return p.commit(), nil
	case pgoutput.Insert:
		p.log.Debug("received INSERT message")
		return p.handleDataMessage(CreateOperation, v.RelationID, v.Row)
	case pgoutput.Update:
		p.log.Debug("received UPDATE message")
		return p.handleDataMessage(UpdateOperation, v.RelationID, v.Row)
	case pgoutput.Delete:
		p.log.Debug("received DELETE message")
		return p.handleDataMessage(DeleteOperation, v.RelationID, v.Row)
	}
	return nil, nil
}

func (p *pgoutputDecoder) commit() []Change {
	currentTxn := p.txnInProgress
	p.txnInProgress = nil
	return currentTxn
}

func (p *pgoutputDecoder) handleDataMessage(operation ChangeOperation, relationID uint32, row []pgoutput.Tuple) ([]Change, error) {
	kind, pk, data, err := p.decoder.DecodeRelationData(relationID, row)
	if err != nil {
		return nil, err
	}

	c := change{kind: kind, pk: pk, data: data, operation: operation}
	if p.txnInProgress == nil {
		return []Change{c}, nil
	}

	p.txnInProgress = append(p.txnInProgress, c)
	return nil, nil
}

type change struct {
	kind      string
	pk        []string
	data      map[string]interface{}
	operation ChangeOperation
}

func (c change) Kind() string                 { return c.kind }
func (c change) PK() []string                 { return c.pk }
func (c change) Data() map[string]interface{} { return c.data }
func (c change) Operation() ChangeOperation   { return c.operation }

type relationDecoder map[uint32]relation

func (d relationDecoder) AddRelation(r pgoutput.Relation) {
	d[r.ID] = relation(r)
}

func (d relationDecoder) DecodeRelationData(
	relationID uint32, row []pgoutput.Tuple,
) (kind string, pk []string, data map[string]interface{}, err error) {
	rel, ok := d[relationID]
	if !ok {
		return "", nil, nil, fmt.Errorf("no relation for %d", relationID)
	}

	pk, data, err = rel.DecodeRow(row)
	if err != nil {
		return "", nil, nil, fmt.Errorf("error decoding row: %v", err)
	}
	return rel.Name, pk, data, nil
}

type relation pgoutput.Relation

func (r relation) DecodeRow(row []pgoutput.Tuple) (pk []string, data map[string]interface{}, err error) {
	return decodeRowData(pgoutput.Relation(r), row)
}

func decodeRowData(
	relation pgoutput.Relation,
	row []pgoutput.Tuple,
) (pk []string, data map[string]interface{}, err error) {
	keys, data := []interface{}{}, map[string]interface{}{}

	if t, c := len(row), len(relation.Columns); t != c {
		return nil, nil, fmt.Errorf("malformed message or relation columns, got %d values but relation has %d columns", t, c)
	}

	for i, tuple := range row {
		col := relation.Columns[i]
		decoder := col.Decoder()
		if err = decoder.DecodeText(nil, tuple.Value); err != nil {
			return nil, nil, fmt.Errorf("error decoding column '%v': %s", col.Name, err)
		}
		value := decoder.Get()
		data[col.Name] = value
		if col.Key {
			keys = append(keys, value)
		}

	}

	pk, err = primaryKeyToStringSlice(keys)
	if err != nil {
		return nil, nil, fmt.Errorf("error creating PK: %v", err)
	}
	if len(pk) == 0 {
		return nil, nil, fmt.Errorf("no primary key specified for row: %v", row)
	}

	return pk, data, nil
}

func primaryKeyToStringSlice(keyValues []interface{}) ([]string, error) {
	keys := []string{}
	for i, pk := range keyValues {
		if pk == nil || pk == "" {
			return nil, fmt.Errorf("primary key value is nil or empty on key element at index %v", i)
		}
		keys = append(keys, fmt.Sprint(pk))
	}
	return keys, nil
}
