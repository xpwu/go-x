package flatfield

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func tagName(tag reflect.StructTag) (name string) {
	return tag.Get("flat")
}

func name(field *reflect.StructField) string {
	if n := tagName(field.Tag); n != "" {
		return n
	}
	return field.Name
}

func TestFlattenTop(t *testing.T) {
	a := assert.New(t)
	type top struct {
    unExport0 float32
		Export1   string `flat:"Export2"`
		unExport1 int
		Export2   struct{}
	}
	fields, err := Flatten(&top{}, Name(tagName))
	a.NoError(err)
	a.Equalf(1, len(fields), "len(fields)")
  f := fields[0]
  a.Equalf(true, f.HasValue, "hasValue")
  a.Equalf("Export2", name(f.SField), "name")
  a.Equalf(1, f.SField.Index[0], "index")
  a.Equalf("Export1", f.SField.Name, "field name")

	//for i, f := range fields {
	//	t.Log(i, "->", name(f.SField), ",type:", f.SField.Type.String(), ", hasValue:", f.HasValue)
	//}
}
