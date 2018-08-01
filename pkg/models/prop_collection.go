package models

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// PropCollectionUpdate operations.
const (
	propCollectionUpdateOperationAdd    = "add"
	propCollectionUpdateOperationModify = "modify"
	propCollectionUpdateOperationSet    = "set"
	propCollectionUpdateOperationDelete = "delete"
)

// PropCollectionUpdate holds update data for collection property (with CollectionType "map" or "list").
type PropCollectionUpdate struct {
	Field     string          `json:"field"`
	Operation string          `json:"operation"`
	Value     json.RawMessage `json:"value"`
	Position  string          `json:"position"`
}

func (u *PropCollectionUpdate) positionForList() (position int, err error) {
	op := strings.ToLower(u.Operation)
	switch op {
	case propCollectionUpdateOperationAdd:
		if len(u.Value) == 0 {
			return 0, errors.Errorf("add operation needs value")
		}
	case propCollectionUpdateOperationModify:
		if len(u.Value) == 0 {
			return 0, errors.Errorf("modify operation needs value")
		}
		position, err = parseListPosition(u.Position)
		if err != nil {
			return 0, errors.Wrap(err, "modify operation needs position")
		}
	case propCollectionUpdateOperationDelete:
		position, err = parseListPosition(u.Position)
		if err != nil {
			return 0, errors.Wrap(err, "delete operation needs position")
		}
	default:
		return 0, errors.Errorf("unsupported operation: %s", u.Operation)
	}
	return position, nil
}

func parseListPosition(s string) (int, error) {
	position, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.Wrap(err, "list position must be a base 10 integer")
	}
	return position, err
}

func (u *PropCollectionUpdate) validateForMap() error {
	op := strings.ToLower(u.Operation)
	switch op {
	case propCollectionUpdateOperationSet:
		if len(u.Value) == 0 {
			return errors.Errorf("set operation needs value")
		}
	case propCollectionUpdateOperationDelete:
		if u.Position == "" {
			return errors.New("delete operation needs position")
		}
	default:
		return errors.Errorf("unsupported operation: %s", u.Operation)
	}
	return nil
}
