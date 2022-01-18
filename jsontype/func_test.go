package jsontype

import (
  "reflect"
)

type S1 struct {
}

func (s1 *S1) F() {

}

func dummyName(tag reflect.StructTag) (name string) {
  return ""
}

