package apiserver

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"unicode"

	"github.com/go-openapi/spec"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// PluginToOpenAPIPaths creates OpenAPI paths based on plugins' endpoints.
func PluginToOpenAPIPaths(ps ...APIPlugin) (spec.Paths, error) {
	paths := spec.Paths{Paths: map[string]spec.PathItem{}}
	for _, p := range ps {
		registerer := &openAPIEndpointRegisterer{
			paths: paths,
			tag:   strings.TrimLeft(fmt.Sprintf("%T", p), "*"),
		}
		p.RegisterHTTPAPI(registerer)
		if registerer.err != nil {
			return spec.Paths{}, registerer.err
		}
	}
	return paths, nil
}

type openAPIEndpointRegisterer struct {
	paths spec.Paths
	tag   string
	err   error
}

// GET registers a GET handler.
func (r *openAPIEndpointRegisterer) GET(path string, _ HandlerFunc, options ...RouteOption) {
	path = sanitizePath(path)
	item := r.paths.Paths[path]
	op, err := optionsToOperation(loadOptions(options...), path, http.MethodGet, r.tag)
	if err != nil {
		r.err = err
		return
	}
	item.Get = op
	r.paths.Paths[path] = item
}

// POST registers a POST handler.
func (r *openAPIEndpointRegisterer) POST(path string, _ HandlerFunc, options ...RouteOption) {
	path = sanitizePath(path)
	item := r.paths.Paths[path]
	op, err := optionsToOperation(loadOptions(options...), path, http.MethodPost, r.tag)
	if err != nil {
		r.err = err
		return
	}
	item.Post = op
	r.paths.Paths[path] = item
}

// PUT registers a PUT handler.
func (r *openAPIEndpointRegisterer) PUT(path string, _ HandlerFunc, options ...RouteOption) {
	path = sanitizePath(path)
	item := r.paths.Paths[path]
	op, err := optionsToOperation(loadOptions(options...), path, http.MethodPut, r.tag)
	if err != nil {
		r.err = err
		return
	}
	item.Put = op
	r.paths.Paths[path] = item
}

// DELETE registers a DELETE handler.
func (r *openAPIEndpointRegisterer) DELETE(path string, _ HandlerFunc, options ...RouteOption) {
	path = sanitizePath(path)
	item := r.paths.Paths[path]
	op, err := optionsToOperation(loadOptions(options...), path, http.MethodDelete, r.tag)
	if err != nil {
		r.err = err
		return
	}
	item.Delete = op
	r.paths.Paths[path] = item
}

func optionsToOperation(ro RouteOptions, path, method, tag string) (*spec.Operation, error) {
	op := spec.NewOperation(operationID(path, method)).WithTags(tag)
	for _, param := range ro.queryParams {
		op = op.AddParam(spec.QueryParam(param))
	}

	if ro.request != nil {
		schema, err := objectToSchema(ro.request)
		if err != nil {
			return nil, err
		}
		op = op.AddParam(spec.BodyParam(typeName(ro.request), schema))
	}

	for code, response := range ro.responses {
		schema, err := objectToSchema(response)
		if err != nil {
			return nil, err
		}
		op = op.RespondsWith(code, spec.NewResponse().WithSchema(schema))
	}

	return op, nil
}

func operationID(path, method string) string {
	return fmt.Sprintf("%s-%s", method, strings.ReplaceAll(strings.TrimLeft(path, "/"), "/", "-"))
}

func typeName(i interface{}) string {
	return reflect.TypeOf(i).Name()
}

func objectToSchema(o interface{}) (*spec.Schema, error) {
	if o == nil {
		return nil, nil
	}
	return typeToSchema(reflect.TypeOf(o))
}

func typeToSchema(t reflect.Type) (*spec.Schema, error) {
	switch t.Kind() {
	case reflect.Int8:
		return spec.Int8Property(), nil
	case reflect.Int16:
		return spec.Int16Property(), nil
	case reflect.Int32:
		return spec.Int32Property(), nil
	case reflect.Int, reflect.Int64:
		return spec.Int64Property(), nil
	case reflect.Uint8:
		return spec.Int8Property().WithMinimum(0, false), nil
	case reflect.Uint16:
		return spec.Int16Property().WithMinimum(0, false), nil
	case reflect.Uint32:
		return spec.Int32Property().WithMinimum(0, false), nil
	case reflect.Uint, reflect.Uint64:
		return spec.Int64Property().WithMinimum(0, false), nil
	case reflect.Bool:
		return spec.BooleanProperty(), nil
	case reflect.String:
		return spec.StringProperty(), nil
	case reflect.Float32:
		return spec.Float32Property(), nil
	case reflect.Float64:
		return spec.Float64Property(), nil
	case reflect.Struct:
		return structToSchema(t)
	case reflect.Slice:
		return sliceToSchema(t)
	case reflect.Array:
		sliceSchema, err := sliceToSchema(t)
		if err != nil {
			return nil, err
		}
		return sliceSchema.WithMinItems(int64(t.Len())).WithMaxItems(int64(t.Len())), nil
	case reflect.Map:
		return new(spec.Schema).Typed("object", ""), nil
	case reflect.Interface:
		return &spec.Schema{}, nil
	default:
		return nil, errors.Errorf("typeToSchema not implemented for type %q\n", t.Kind())
	}
}

func structToSchema(t reflect.Type) (*spec.Schema, error) {
	schema := new(spec.Schema).Typed("object", "")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		propName := f.Tag.Get("json")
		if propName == "-" || !isFieldPublic(f.Name) {
			continue
		}
		propName = strings.Split(propName, ",")[0]
		if propName == "" {
			propName = f.Name
		}
		fieldSchema, err := typeToSchema(f.Type)
		if err != nil {
			return nil, err
		}
		schema = schema.SetProperty(propName, *fieldSchema)

	}
	return schema, nil
}

func isFieldPublic(name string) bool {
	return unicode.IsUpper([]rune(name)[0])
}

func sliceToSchema(t reflect.Type) (*spec.Schema, error) {
	schema, err := typeToSchema(t.Elem())
	if err != nil {
		return nil, err
	}
	return spec.ArrayProperty(schema), nil
}

// Group satisfies the HTTPRouter interface.
func (r *openAPIEndpointRegisterer) Group(prefix string, _ ...RouteOption) {
	logrus.WithField("prefix", prefix).Debug("Registering groups of paths is not supported - ignoring the group")
}

func loadOptions(options ...RouteOption) RouteOptions {
	r := RouteOptions{}
	r.apply(options)
	return r
}

// Use satisfies the HTTPRouter interface.
func (r *openAPIEndpointRegisterer) Use(_ ...MiddlewareFunc) {}
