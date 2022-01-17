package flatfield

import (
  "errors"
  "reflect"
  "sort"
)

type option struct {
  nameF func(tag reflect.StructTag) (name string)
}

type Option func(op *option)


// 通过tag获取名字
func Name(nf func(tag reflect.StructTag)(name string)) Option {
  return func(op *option) {
    op.nameF = nf
  }
}

// 1、处理嵌套的匿名struct field(embedded field) 到top-level
// 2、去重所有同名的field: 可能因为取了相同的tag而造成
// 3、去掉所有的非导出field
// 判断重名的名字来源：1、取tag中的名字；2、取FieldName
// 如果匿名struct通过tag能够取到名字，就不认为是匿名struct
// 返回的排序按照struct field的代码书写顺序, breadth-first 而排
// hasValue: 此域是否有合法的值

type FlatField struct {
  SField *reflect.StructField
  HasValue bool
}

func Flatten(st interface{}, opts ...Option) (fields []FlatField, err error) {
  op := &option{
    nameF: func(tag reflect.StructTag) string {
      return ""
    },
  }
  for _,of := range opts {
    of(op)
  }

  typ := reflect.TypeOf(st)
  if typ.Kind() == reflect.Ptr {
    typ = typ.Elem()
  }

  if typ.Kind() != reflect.Struct {
    return nil, errors.New("func Flatten(st, ...) --- st must be a struct or struct pointer")
  }

  nameSet := map[string]bool{}
  flds := make([]reflect.StructField, 0)

  for i := 0; i < typ.NumField(); i++ {
    flat(typ.Field(i), &flds, op)
  }

  byBreadthFirst(flds).Sort()

  fields = make([]FlatField, 0, len(flds))
  for i,f := range flds {
    name := op.nameF(f.Tag)
    if name == "" {
      name = f.Name
    }
    // 重名了，tag设置而引起
    if nameSet[name] {
      continue
    }
    nameSet[name] = true

    fields = append(fields, FlatField{
      SField:   &flds[i],
      HasValue: hasValue(st, flds[i].Index),
    })
  }

  return
}

func hasValue(i interface{}, index []int) (ok bool) {
  defer func() {
   // FieldByIndex 取不到值时，会panic
   if r := recover(); r != nil {
     ok = false
   }
  }()

  value := reflect.ValueOf(i)
  if !value.IsValid() {
    return false
  }
  if value.Kind() == reflect.Ptr && value.IsNil() {
    return false
  }
  value = value.Elem()

  return value.FieldByIndex(index).IsValid()
}

func flat(input reflect.StructField, flds *[]reflect.StructField, opt *option) {
  isUnexported := input.PkgPath != ""
  if input.Anonymous {
    t := input.Type
    if t.Kind() == reflect.Ptr {
      t = t.Elem()
    }
    if isUnexported && t.Kind() != reflect.Struct {
      // Ignore embedded fields of unexported non-struct types.
      return
    }
    // Do not ignore embedded fields of unexported struct types
    // since they may have exported fields.
  } else if isUnexported {
    // Ignore unexported non-embedded fields.
    return
  }

  name := opt.nameF(input.Tag)
  if !input.Anonymous || (input.Anonymous && name != "") {
    *flds = append(*flds, input)
    return
  }

  // follow: input.Anonymous && name == ""
  t := input.Type
  if t.Kind() == reflect.Ptr {
    t = t.Elem()
  }
  if t.Kind() != reflect.Struct {
    // Ignore embedded and name=="" fields  of non-struct types.
    return
  }
  // t is struct
  for i := 0; i < t.NumField(); i++ {
    flat(t.Field(i), flds, opt)
  }
}

// byBreadthFirst: sorts field by byBreadthFirst.
type byBreadthFirst []reflect.StructField

func (x byBreadthFirst) Len() int { return len(x) }

func (x byBreadthFirst) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func (x byBreadthFirst) Less(i, j int) bool {
  indexL := x[i].Index
  indexR := x[j].Index
  // 深度小的在前
  if len(indexL) != len(indexR) {
    return len(indexL) < len(indexR)
  }
  // 序号小的在前
  for ii,lv := range indexL {
    rv := indexR[ii]
    if lv != rv {
      return lv < rv
    }
  }

  // 不会存在两个index完全一样的field
  panic("same []index is impossible")
}

func (x byBreadthFirst) Sort() {
  sort.Sort(x)
}


