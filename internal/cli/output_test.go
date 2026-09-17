package cli

import (
	"reflect"
	"testing"

	"github.com/lumastack/luma-backlog/internal/app"
)

// Every field a listing can show reaches the JSON, and is wired through rather
// than merely declared.
//
// `--json` is published contract (spec.md §9.3), and a field missing from it is
// worse than a field missing from the table: a machine surface that silently
// omits something makes every claim about that something untestable. **Measured:
// `rank` was absent for as long as it existed** --- and ADR-0005's whole
// justification for the ordinal prefix is that a consumer outside this tool can
// sort the field alone, which was therefore false in the tool's own machine
// output and checkable by nobody.
//
// Two failures are possible and this catches both. A field added to View and
// not to itemJSON never appears. A field added to both and forgotten in
// toItemJSON appears, always empty --- which reads as "this record has none"
// and is the more misleading of the two.
func TestEveryViewFieldReachesTheJSON(t *testing.T) {
	// Every field non-zero, so anything that fails to carry through shows up
	// as a zero on the other side.
	var view app.View
	fillNonZero(reflect.ValueOf(&view).Elem())

	item := reflect.ValueOf(toItemJSON(view))
	shape := item.Type()

	viewType := reflect.TypeOf(view)
	for i := 0; i < viewType.NumField(); i++ {
		name := viewType.Field(i).Name
		got, ok := shape.FieldByName(name)
		if !ok {
			t.Errorf("View.%s has no itemJSON field --- it cannot appear in --json at all", name)
			continue
		}
		if item.FieldByIndex(got.Index).IsZero() {
			t.Errorf("View.%s has an itemJSON field and toItemJSON does not fill it --- it would always emit empty", name)
		}
	}
}

// fillNonZero gives every field in a struct a value it does not have by
// default, so a zero on the far side of a conversion means the conversion
// dropped it rather than that the input was empty.
func fillNonZero(v reflect.Value) {
	switch v.Kind() {
	case reflect.String:
		v.SetString("x")
	case reflect.Int:
		v.SetInt(1)
	case reflect.Bool:
		v.SetBool(true)
	case reflect.Pointer:
		v.Set(reflect.New(v.Type().Elem()))
		fillNonZero(v.Elem())
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if v.Field(i).CanSet() {
				fillNonZero(v.Field(i))
			}
		}
	}
}
