package jsontype

import (
  "encoding/base64"
  "encoding/json"
  "github.com/stretchr/testify/assert"
  "reflect"
  "testing"
)

func TestString_ToValue_2Slice(t *testing.T) {
  a := assert.New(t)
  s0 := []byte{0, 1, 2}

  n0 := String(base64.StdEncoding.EncodeToString(s0))

  var s []byte

  if a.NoError(n0.ToValue(&s, dummyName)) {
    a.Equal(s0, s)
  }
}

func TestString_ToValue_2String(t *testing.T) {
  a := assert.New(t)
  s0 := "this is a string"
  s := String(s0)

  sn := "this is another string"

  if a.NoError(s.ToValue(&sn, dummyName)) {
    a.Equal(s0, sn)
  }

  // reflect.Value
  sn = "this is another string"
  v := reflect.ValueOf(&sn)
  if a.NoError(s.ToValue(&v, dummyName)) {
    a.Equal(s0, sn)
  }

  // self
  sm := String("error")
  if a.NoError(s.ToValue(&sm, dummyName)) {
    a.Equal(s, sm)
  }
}

func TestString_ToValue_2Interface(t *testing.T) {
  a := assert.New(t)
  s0 := "this is a string"
  s := String(s0)

  var sn interface{}

  if a.NoError(s.ToValue(&sn, dummyName)) {
    a.Equal(s0, sn)
  }
}

func TestString_ToValue_2Number(t *testing.T) {
  a := assert.New(t)
  s0 := "123"
  s := String(s0)

  var sn Number

  if a.NoError(s.ToValue(&sn, dummyName)) {
    a.Equal(s0, sn.String())
  }

  var sm json.Number
  if a.NoError(s.ToValue(&sm, dummyName)) {
    a.Equal(s0, sm.String())
  }
}

