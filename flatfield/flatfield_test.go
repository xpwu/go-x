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

func TestFlatten(t *testing.T) {
  a := assert.New(t)

  type anony1 struct {
    Anony1Export1   string `flat:"An1Export1"`
    anony1Unexport1 bool
    anony1Unexport2 bool
    Anony1Export2   string `flat:"An1Export2"`
  }

  type anony20 struct {
    Anony20Export1   string `flat:"An20Export1"`
    anony20Unexport1 bool
    anony20Unexport2 bool
    Anony20Export2   string `flat:"An20Export2"`
  }

  type anony2 struct {
    Anony2Export1   string `flat:"An2Export1"`
    anony2Unexport1 bool
    anony1
    anony2Unexport2 bool
    anony20
    Anony2Export2   string `flat:"An2Export2"`
    Anony2Export3   string `flat:"An3Export2" note:"cover An3Export2"`
  }

  type anony3 struct {
    Anony3Export1   string `flat:"An3Export1"`
    anony3Unexport1 bool
    anony3Unexport2 bool
    Anony3Export2   string `flat:"An3Export2"`
  }

  type top struct {
    unExport0 float32
    anony1
    Export1   string `flat:"Export1"`
    anony2
    unExport1 int
    *anony3
    Export2   struct{}
  }

  fields, err := Flatten(&top{}, Name(tagName))
  a.NoError(err)
  //a.Equalf(1, len(fields), "len(fields)")
  //f := fields[0]
  //a.Equalf(true, f.HasValue, "hasValue")
  //a.Equalf("Export2", name(f.SField), "name")
  //a.Equalf(1, f.SField.Index[0], "index")
  //a.Equalf("Export1", f.SField.Name, "field name")

  cases := []struct{
    index []int
    hasValue bool
  }{
    {index: []int{2}, hasValue: true},
    {index: []int{6}, hasValue: true},
    {index: []int{1, 0}, hasValue: true},
    {index: []int{1, 3}, hasValue: true},
    {index: []int{3, 0}, hasValue: true},
    {index: []int{3, 5}, hasValue: true},
    {index: []int{3, 6}, hasValue: true},
    {index: []int{5, 0}, hasValue: false},
    {index: []int{3, 4, 0}, hasValue: true},
    {index: []int{3, 4, 3}, hasValue: true},
  }

  a.Equal(len(cases), len(fields), "len(fields)")

  for i, f := range fields {
    a.Equal(cases[i].index, f.SField.Index, "[]index")
    a.Equal(cases[i].hasValue, f.HasValue, "hasValue")
  	//t.Log(f.SField.Index, "->FieldName:", f.SField.Name, ", tagName:", name(f.SField),
  	//  ", tag:", f.SField.Tag, ", type:", f.SField.Type.String(), ", hasValue:", f.HasValue)
  }
}

// todo 测试最大嵌套；测试自己嵌套自己
