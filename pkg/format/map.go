package format

import (
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

type MapApplier interface {
	ApplyMap(map[string]interface{}) error
}

func ApplyMap(m map[string]interface{}, o interface{}) error {
	if len(m) == 0 {
		return nil
	}
	ov := reflect.ValueOf(o)
	if ov.Kind() != reflect.Ptr {
		return errors.Errorf("cannot apply to non pointer %T", o)
	}
	if ov.IsNil() {
		return errors.Errorf("cannot apply to nil pointer %T", o)
	}
	iv := reflect.Indirect(ov)
	if indirect(iv.Type()).Kind() != reflect.Struct {
		return errors.Errorf("cannot apply map to %T", iv.Interface())
	}
	if iv.Kind() == reflect.Ptr && iv.IsNil() {
		iv.Set(reflect.New(iv.Type().Elem()))
	}
	sv := reflect.Indirect(iv)
	if sv.Kind() != reflect.Struct {
		return errors.Errorf("cannot apply map to %T", sv.Interface())
	}

	for i := 0; i < sv.NumField(); i++ {
		k := fieldKey(sv.Type().Field(i))
		fieldValue, ok := m[k]
		if !ok {
			continue
		}
		if fieldValue == nil {
			continue
		}
		f := sv.Field(i)
		var err error
		if pointerToNonStruct(f) {
			return errors.Errorf("only pointer to struct fields can be applied")
		}

		if !f.CanSet() || !f.CanInterface() {
			return errors.Errorf("cannot set value of %s", k)
		}
		switch f.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			f.SetInt(InterfaceToInt64(fieldValue))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			f.SetUint(InterfaceToUint64(fieldValue))
		case reflect.Bool:
			f.SetBool(InterfaceToBool(fieldValue))
		case reflect.String:
			f.SetString(InterfaceToString(fieldValue))
		case reflect.Float32, reflect.Float64:
			f.SetFloat(InterfaceToFloat(fieldValue))
		case reflect.Array, reflect.Slice:
			vv := reflect.ValueOf(fieldValue)
			if vv.IsNil() {
				continue
			}
			if !(vv.Type().Kind() == reflect.Array || vv.Type().Kind() == reflect.Slice) {
				err = errors.Errorf("cannot apply %T onto %T", fieldValue, f.Interface())
				break
			}
			for n := 0; n < vv.Len(); n++ {
				if vv.Index(n).Type() != f.Type().Elem() {
					if !(vv.Index(n).Kind() == reflect.Interface && vv.Index(n).Elem().Type() == f.Type().Elem()){
						err = errors.Errorf("incompatible array item types '%s' != '%s'", vv.Index(n).Type(), f.Type().Elem())
						break
					}
				}
				f.Set(reflect.Append(f,reflect.ValueOf(vv.Index(n).Interface())))
			}
		case reflect.Ptr, reflect.Struct:
			sm, ok := fieldValue.(map[string]interface{})
			if !ok {
				err = errors.Errorf("cannot apply %T onto %T", fieldValue, f.Interface())
				break
			}
			err = ApplyMap(sm, f.Addr().Interface())
		case reflect.Interface:
			if f.IsNil() {
				continue
			}
			sm, ok := fieldValue.(map[string]interface{})
			if !ok {
				err = errors.Errorf("cannot apply %T onto %T", fieldValue, f.Interface())
				break
			}
			if f.Elem().Kind() != reflect.Ptr {
				err = errors.Errorf("cannot mutate non pointer %T", f.Interface())
				break
			}
			err = ApplyMap(sm, f.Interface())
		case reflect.Map:
			sm, ok := fieldValue.(map[string]interface{})
			if !ok {
				err = errors.Errorf("cannot apply %T onto %T", fieldValue, f.Interface())
				break
			}
			if f.IsNil() {
				f.Set(reflect.MakeMap(f.Type()))
			}
			a, ok := f.Addr().Interface().(MapApplier)
			if ok {
				return a.ApplyMap(sm)
			}
			err = errors.New("map field needs to implement MapApplier interface")
		default:
			return errors.Errorf("applying field of type: '%s' not implemented", f.Kind())
		}
		if err != nil {
			return errors.Wrapf(err, "failed to apply field %s", k)
		}
	}
	return nil
}

func pointerToNonStruct(v reflect.Value) bool {
	return v.Kind() == reflect.Ptr && v.Type().Elem().Kind() != reflect.Struct
}

func fieldKey(s reflect.StructField) string {
	tag, ok := s.Tag.Lookup("json")
	if !ok {
		return s.Name
	}
	ss := strings.Split(tag, ",")
	return ss[0]
}

func indirect(t reflect.Type) reflect.Type {
	for t != nil && t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}
