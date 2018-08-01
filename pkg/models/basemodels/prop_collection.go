package basemodels

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
	PropCollectionUpdateOperationSet    = "set"
	PropCollectionUpdateOperationDelete = "delete"
)

// PropCollectionUpdate holds update data for collection property (with CollectionType "map" or "list").
type PropCollectionUpdate struct {
	Field     string          `json:"field"`
	Operation string          `json:"operation"`
	Value     json.RawMessage `json:"value"`
	Position  string          `json:"position"`
}

// PositionForList parses position and validates operation for ListProperty collection update.
func (u *PropCollectionUpdate) PositionForList() (position int, err error) {
	op := strings.ToLower(u.Operation)
	switch op {
	case PropCollectionUpdateOperationAdd:
		if len(u.Value) == 0 {
			return 0, errors.Errorf("add operation needs value")
		}
	case PropCollectionUpdateOperationModify:
		if len(u.Value) == 0 {
			return 0, errors.Errorf("modify operation needs value")
		}
		position, err = parseListPosition(u.Position)
		if err != nil {
			return 0, errors.Wrap(err, "modify operation needs position")
		}
	case PropCollectionUpdateOperationDelete:
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

// ValidateForMap validates MapProperty collection update.
func (u *PropCollectionUpdate) ValidateForMap() error {
	op := strings.ToLower(u.Operation)
	switch op {
	case PropCollectionUpdateOperationSet:
		if len(u.Value) == 0 {
			return errors.Errorf("set operation needs value")
		}
	case PropCollectionUpdateOperationDelete:
		if u.Position == "" {
			return errors.New("delete operation needs position")
		}
	default:
		return errors.Errorf("unsupported operation: %s", u.Operation)
	}
	return nil
}
