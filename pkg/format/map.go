package format

import (
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

type mapApplier interface {
	ApplyMap(map[string]interface{}) error
}

// ApplyMap applies map onto interface which holds a struct.
func ApplyMap(m map[string]interface{}, object interface{}) error {
	if len(m) == 0 {
		return nil
	}
	sv, err := getSettableStructValue(object)
	if err != nil {
		return err
	}
	if ma, ok := sv.Addr().Interface().(mapApplier); ok {
		return ma.ApplyMap(m)
	}
	for i := 0; i < sv.NumField(); i++ {
		k := fieldKey(sv.Type().Field(i))
		fieldValue, ok := m[k]
		if !ok || fieldValue == nil {
			continue
		}
		if isPointerToNonStruct(sv.Field(i)) {
			return errors.Errorf("only pointer to struct fields can be applied, field %s", k)
		}
		err := applyValue(sv.Field(i), fieldValue)
		if err != nil {
			return errors.Wrapf(err, "failed to apply field %s", k)
		}
	}
	return nil
}

func getSettableStructValue(o interface{}) (reflect.Value, error) {
	ov := reflect.ValueOf(o)
	if ov.Kind() != reflect.Ptr {
		return reflect.Value{}, errors.Errorf("cannot apply to non pointer %T", o)
	}
	if ov.IsNil() {
		return reflect.Value{}, errors.Errorf("cannot apply to nil pointer %T", o)
	}
	iv := reflect.Indirect(ov)
	if iv.Kind() == reflect.Ptr && iv.IsNil() {
		iv.Set(reflect.New(iv.Type().Elem()))
	}
	sv := reflect.Indirect(iv)
	if sv.Kind() != reflect.Struct {
		return reflect.Value{}, errors.Errorf("cannot apply map to %T", sv.Interface())
	}
	return sv, nil
}

// nolint: gocyclo
func applyValue(v reflect.Value, i interface{}) error {
	if !isSettable(v) {
		return errors.Errorf("cannot set value of %s", v.Type().Name())
	}
	applied, err := useMapApplierIfAvailable(v, i)
	if err != nil {
		return err
	}
	if applied {
		return nil
	}

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(InterfaceToInt64(i))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(InterfaceToUint64(i))
	case reflect.Bool:
		v.SetBool(InterfaceToBool(i))
	case reflect.String:
		v.SetString(InterfaceToString(i))
	case reflect.Float32, reflect.Float64:
		v.SetFloat(InterfaceToFloat(i))
	case reflect.Array, reflect.Slice:
		return applySlice(v, i)
	case reflect.Ptr, reflect.Struct:
		return applyStruct(i, v.Addr().Interface())
	case reflect.Interface:
		return applyInterface(v, i)
	case reflect.Map:
		return errors.New("map field needs to implement mapApplier interface")
	default:
		return errors.Errorf("applying field of type: '%s' not implemented", v.Kind())
	}
	return nil
}

func isSettable(v reflect.Value) bool {
	return v.CanSet() && v.CanInterface() && v.CanAddr()
}

func useMapApplierIfAvailable(v reflect.Value, i interface{}) (bool, error) {
	a, ok := v.Addr().Interface().(mapApplier)
	if !ok {
		return false, nil
	}
	if v.Addr().IsNil() {
		v.Addr().Set(reflect.New(v.Type()))
	}
	sm, ok := i.(map[string]interface{})
	if !ok {
		return false, errors.Errorf("cannot apply %T onto %T", i, v.Interface())
	}
	return true, a.ApplyMap(sm)
}

func applySlice(v reflect.Value, i interface{}) error {
	vv := reflect.ValueOf(i)
	if vv.IsNil() {
		return nil
	}
	if !(vv.Type().Kind() == reflect.Array || vv.Type().Kind() == reflect.Slice) {
		return errors.Errorf("cannot apply %T onto %T", i, v.Interface())
	}
	for n := 0; n < vv.Len(); n++ {
		tmp := reflect.New(v.Type().Elem())
		err := applyValue(tmp.Elem(), vv.Index(n).Interface())
		if err != nil {
			return err
		}
		v.Set(reflect.Append(v, tmp.Elem()))
	}
	return nil
}

func applyStruct(i interface{}, s interface{}) error {
	sm, ok := i.(map[string]interface{})
	if !ok {
		return errors.Errorf("cannot apply %T onto %T", i, s)
	}
	return ApplyMap(sm, s)
}

func applyInterface(v reflect.Value, i interface{}) error {
	if v.IsNil() {
		return nil
	}
	if v.Elem().Kind() != reflect.Ptr {
		return errors.Errorf("cannot mutate non pointer %T", v.Interface())
	}
	return applyStruct(i, v.Interface())
}

func isPointerToNonStruct(v reflect.Value) bool {
	return v.Kind() == reflect.Ptr && v.Type().Elem().Kind() != reflect.Struct
}

func fieldKey(s reflect.StructField) string {
	tag, ok := s.Tag.Lookup("json")
	if !ok {
		return s.Name
	}
	return strings.Split(tag, ",")[0]
}

func indirect(t reflect.Type) reflect.Type {
	for t != nil && t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}
