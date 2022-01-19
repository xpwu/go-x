package jsontype

import (
  "encoding/json"
  "fmt"
  "reflect"
)

func (s Slice) Kind() Kind {
  return SliceK
}

func (s Slice) String() string {
  return fmt.Sprint([]Type(s))
}

func (s Slice) ToValue(i interface{}, name func(tag reflect.StructTag) (name string)) error {
  value := indirect(i, false)
  if !value.IsValid() {
    // skip
    return nil
  }

  if value.Type() == reflect.TypeOf(s) {
    value.Set(reflect.ValueOf(s))
    return nil
  }

  // Check type of target.
  switch value.Kind() {
  default:
    return &json.UnmarshalTypeError{Value: "array", Type: value.Type()}

  case reflect.Interface:
    if value.NumMethod() != 0 {
      return &json.UnmarshalTypeError{Value: "array", Type: value.Type()}
    }
    v := make([]interface{}, len(s), len(s))
    for i := 0; i < len(s) && i < len(v); i++ {
      if err := s[i].ToValue(&v[i], name); err != nil {
        return err
      }
    }
    value.Set(reflect.ValueOf(v))
  case reflect.Slice:
    value.Set(reflect.MakeSlice(value.Type(), len(s), len(s)))
    fallthrough
  case reflect.Array:
    for i := len(s); i < value.Len(); i++ {
      // len(value) > len(s), zero value
      value.Index(i).Set(reflect.Zero(value.Type().Elem()))
    }

    for i := 0; i < value.Len() && i < len(s); i++ {
      if err := s[i].ToValue(value.Index(i), name); err != nil {
        return err
      }
    }
  }

  return nil
}

