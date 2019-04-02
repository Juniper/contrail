# Code review guidelines

This file contains guidelines for commiters and code reviewers.
It is open for updates.

## References

Code should be compliant with rules from:
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Code review comments](https://github.com/golang/go/wiki/CodeReviewComments)

## Extra rules

### Organizing imports

Imports should be split into 4 groups:
1.  standard library packages, e.g. `net/http`
2.  repo packages, e.g.

```go
"github.com/pkg/errors"
"github.com/Juniper/contrail"
```

3. renamed imports: 
    e.g. `apicommon "github.com/Juniper/contrail/pkg/apisrv/common"`
4. "_" imports (these should only exist in main or tests)
    e.g. `_ "github.com/Juniper/contrail/pkg/apisrv/common"`



Additionally, each of the groups should be sorted alphabetically -
`goimports` will do it for you.

```diff
-// Don't
-import (
-	"fmt"
-
-  "net/http"
-  apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
-  "github.com/Juniper/contrail/pkg/models/basemodels"
-  "github.com/labstack/echo"
-  _ "github.com/Juniper/contrail/pkg/keystone"
-)
+// Do
+import (
+  "fmt"
+  "net/http"
+
+  "github.com/labstack/echo"
+  "github.com/Juniper/contrail/pkg/models/basemodels"
+
+  apicommon "github.com/Juniper/contrail/pkg/apisrv/common"
+
+  _ "github.com/Juniper/contrail/pkg/keystone"
+)
```

### Function signature formatting

Function signatures that are too long (120 chars) should be split.
Below is a list of possible formattings:
1.  Single line - if everything fits in 120 chars

```go
func numericProtocolForEthertype(protocol, ethertype string) (numericProtocol string, err error) {
	// ...
}
```

2.  Multiline short - if function name with arguments is too long, but arguments
    alone aren't longer than 120 chars.

```go
func decodeRowData(
	relation pgoutput.Relation, row []pgoutput.Tuple,
) (pk string, err error) {
	// ...
}
```

3. Multiline long - if input or output arguments (or both) are longer than 120 chars.

```go
func decodeRowData(
	relation pgoutput.Relation,
	row []pgoutput.Tuple,
) (pk string, data map[string]interface{}, err error) {
	// ...
}
```

or even

```go
func decodeRowData(
	relation pgoutput.Relation,
	row []pgoutput.Tuple,
) (
	pk string,
	data map[string]interface{},
	err error,
) {
	// ...
}
```

## Common mistakes

### Named Result Parameters

Result parameters can be named for documentational purposes.
However, "bare returns" (returns that don't reference the variables
that are returned) are discouraged.

```diff
-// Don't
-func parseRouteTarget(rtName string) (ip net.IP, asn int, target int, err error) {
-	// ...
-	return
-}
+// Do
+func parseRouteTarget(rtName string) (ip net.IP, asn int, target int, err error) {
+	// ...
+	return ip, asn, target, err
+}
```

See: [CRC: Named Result Parameters](https://github.com/golang/go/wiki/CodeReviewComments#named-result-parameters)

### Package and type names

New package names should be in singular form with a short word or abbreviation
describing package contents. Names like `common` or `util` are discouraged.

The package name shouldn't be repeated in type/variable/interface name,
because client packages would reference the package name twice.

```diff
-// Don't
-package types
-
-type TypesService {
-	// ...
-}
+// Do
+package types
+
+type Service {
+	// ...
+}
```

See: [CRC: Package Names](https://github.com/golang/go/wiki/CodeReviewComments#package-names)

## Templates

### Template file indentation

New template files should be indented in the same way
as the language they generate.
In particular, Go template files should be indented with tabs.

```diff
-// Don't
-type WriteService interface {
-{% for schema in schemas %}{% if schema.Type != "abstract" and schema.ID %}
-    Create{{ schema.JSONSchema.GoName }}(context.Context, *Create{{ schema.JSONSchema.GoName }}Request) (*Create{{ schema.JSONSchema.GoName }}Response, error)
+// Do
+type WriteService interface {
+{% for schema in schemas %}{% if schema.Type != "abstract" and schema.ID %}
+	Create{{ schema.JSONSchema.GoName }}(context.Context, *Create{{ schema.JSONSchema.GoName }}Request) (*Create{{ schema.JSONSchema.GoName }}Response, error)
```
