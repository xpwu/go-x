package jsontype

import (
  "encoding/json"
  "reflect"
)

func (b Bool) Kind() Kind {
  return BoolK
}

func (b Bool) ToValue(v interface{}, name func(tag reflect.StructTag) (name string)) error {
  value := indirect(v, false)
  if !value.IsValid() {
    // skip
    return nil
  }

  switch value.Kind() {
  case reflect.Bool:
    value.SetBool(bool(b))
    return nil
  case reflect.Interface:
    if value.NumMethod() == 0 {
      value.Set(reflect.ValueOf(bool(b)))
      return nil
    }
  }

  return &json.UnmarshalTypeError{Value: "bool", Type: value.Type()}
}

