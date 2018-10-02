package cache

import (
	"fmt"

	"github.com/Juniper/contrail/pkg/models/basemodels"
	uuid "github.com/satori/go.uuid"
)

//UUIDIndex implement uuid indexer for go-memdb.
type UUIDIndex struct {
}

//FromObject makes index from object.
func (i *UUIDIndex) FromObject(o interface{}) (bool, []byte, error) {
	object, ok := o.(basemodels.Object)
	if !ok {
		return false, nil, nil
	}
	s := object.GetUUID() + "\x00"
	u, err := uuid.FromString(s)
	if err != nil {
		return false, nil, nil
	}
	return true, u.Bytes(), nil
}

//FromArgs makes uuid from args
func (i *UUIDIndex) FromArgs(args ...interface{}) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("must provide only a single argument")
	}
	switch arg := args[0].(type) {
	case string:
		u, err := uuid.FromString(arg)
		if err != nil {
			return nil, err
		}
		return u.Bytes(), nil
	case []byte:
		if len(arg) != 16 {
			return nil, fmt.Errorf("byte slice must be 16 characters")
		}
		return arg, nil
	default:
		return nil,
			fmt.Errorf("argument must be a string or byte slice: %#v", args[0])
	}
}

//PrefixFromArgs makes a prefix for search.
func (i *UUIDIndex) PrefixFromArgs(args ...interface{}) ([]byte, error) {
	val, err := i.FromArgs(args...)
	if err != nil {
		return nil, err
	}

	// Strip the null terminator, the rest is a prefix
	n := len(val)
	if n > 0 {
		return val[:n-1], nil
	}
	return val, nil
}

//OwnerIndex implement owner based index.
type OwnerIndex struct {
}

// FromObject makes index from object.
func (i *OwnerIndex) FromObject(o interface{}) (bool, []byte, error) {
	object, ok := o.(basemodels.Ownable)
	if !ok {
		return false, nil, nil
	}
	owner := object.GetPerms2Owner() + "\x00"
	return true, []byte(owner), nil
}

// FromArgs makes uuid from args.
func (i *OwnerIndex) FromArgs(args ...interface{}) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("must provide only a single argument")
	}
	arg, ok := args[0].(string)
	if !ok {
		return nil, fmt.Errorf("argument must be a string: %#v", args[0])
	}
	return []byte(arg + "\x00"), nil
}

//PrefixFromArgs makes a prefix for search.
func (i *OwnerIndex) PrefixFromArgs(args ...interface{}) ([]byte, error) {
	val, err := i.FromArgs(args...)
	if err != nil {
		return nil, err
	}

	// Strip the null terminator, the rest is a prefix
	n := len(val)
	if n > 0 {
		return val[:n-1], nil
	}
	return val, nil
}
