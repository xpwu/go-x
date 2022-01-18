package jsontype

import "reflect"

func (n Null) Kind() Kind {
  return NullK
}

func (n Null) ToValue(i interface{}, name func(tag reflect.StructTag) (name string)) error {
  value := indirect(i, true)
  if !value.IsValid() {
    // skip
    return nil
  }

  switch value.Kind() {
  case reflect.Interface, reflect.Ptr, reflect.Map, reflect.Slice:
    value.Set(reflect.Zero(value.Type()))
    // otherwise, ignore null for primitives/string
  }
  return nil
}

func (n Null) MarshalJSON() ([]byte, error) {
  return []byte("null"), nil
}

