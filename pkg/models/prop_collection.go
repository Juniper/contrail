package models

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// PropCollectionUpdate operations.
const (
	PropCollectionUpdateOperationAdd    = "add"
	PropCollectionUpdateOperationModify = "modify"
	PropCollectionUpdateOperationDelete = "delete"
	PropCollectionUpdateOperationSet    = "set"
)

type PropCollectionUpdate struct {
	Field     string          `json:"field"`
	Operation string          `json:"operation"`
	Value     json.RawMessage `json:"value"`
	Position  string          `json:"position"`
}

func (p *PropCollectionUpdate) positionForList() (pos int, err error) {
	op := strings.ToLower(p.Operation)
	switch op {
	case PropCollectionUpdateOperationAdd:
		if len(p.Value) == 0 {
			return 0, errors.Errorf("add operation needs field value")
		}
	case PropCollectionUpdateOperationModify:
		if len(p.Value) == 0 {
			return 0, errors.Errorf("modify operation needs field value")
		}
		pos, err = parseListPosition(p.Position)
		if err != nil {
			return 0, errors.Wrap(err, "modify operation needs position")
		}
	case PropCollectionUpdateOperationDelete:
		pos, err = parseListPosition(p.Position)
		if err != nil {
			return 0, errors.Wrap(err, "delete operation needs position")
		}
	default:
		return 0, errors.Errorf("unsupported operation: %s", p.Operation)
	}
	return pos, nil
}

func parseListPosition(s string) (int, error) {
	position, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.Wrap(err, "list position must be a number")
	}
	return position, err
}

func (p *PropCollectionUpdate) validateForMap() error {
	op := strings.ToLower(p.Operation)
	switch op {
	case PropCollectionUpdateOperationSet:
		if len(p.Value) == 0 {
			return errors.Errorf("set operation needs field value")
		}
	case PropCollectionUpdateOperationDelete:
		if p.Position == "" {
			return errors.New("delete operation needs position")
		}
	default:
		return errors.Errorf("unsupported operation: %s", p.Operation)
	}
	return nil
}
