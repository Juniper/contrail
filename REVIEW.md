# Code review guidlines

This file contains guidlines for commiters and code reviewers. This file is open
for updates.

## References
Code should be compliant with rules from:
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Code review comments](https://github.com/golang/go/wiki/CodeReviewComments)

## Extra rules

### Function signature formatting
Funtion signatures that are too long (120 chars) should be splitted.
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
    ```
3. Multiline long - if input or output arguments (or both) are longer than 120 chars.
    ```go
    func decodeRowData(
	    relation pgoutput.Relation,
	    row []pgoutput.Tuple,
    ) (pk string, data map[string]interface{}, err error) {
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
    ```

## Common mistakes

### Named Result Parameters
Result parameters can be named for documentational purposes, however "bare returns"
(return that doesn't reference variables that are returned) are discoraged.

```diff
- // Don't
- func parseRouteTarget(rtName string) (ip net.IP, asn int, target int, err error) {
-     // ...
-     return
- }
+ // Do
+ func parseRouteTarget(rtName string) (ip net.IP, asn int, target int, err error) {
+     // ...
+     return ip, asn, target, err
+ }
```
See: [CRC: Named Result Parameters](https://github.com/golang/go/wiki/CodeReviewComments#named-result-parameters)

### Package and type names
New package names should be in singular form with short word or abbreviation
descripting package contents.  Names like `common` or `util` are discoraged.

If the package name shouldn't be repeated in type/variable/interface name,
because client packages would reference package name twice.

```diff
- // Don't
- package types
-
- type TypesService {
- // ...
- }
+ // Do
+ package types
+
+ type Service {
+ // ...
+ }
```
See: [CRC: Package Names](https://github.com/golang/go/wiki/CodeReviewComments#package-names)
