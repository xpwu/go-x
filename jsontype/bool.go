package jsontype

import (
  "encoding/json"
  "fmt"
  "reflect"
)

func (b Bool) Kind() Kind {
  return BoolK
}

func (b Bool) String() string {
  return fmt.Sprint(bool(b))
}

func (b Bool) ToValue(v interface{}, name func(tag reflect.StructTag) (name string)) error {
  value := indirect(v, false)
  if !value.IsValid() {
    // skip
    return nil
  }

  if value.Type() == reflect.TypeOf(b) {
    value.Set(reflect.ValueOf(b))
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

