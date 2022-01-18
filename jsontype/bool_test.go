package jsontype

import (
  "github.com/stretchr/testify/assert"
  "reflect"
  "testing"
)

func TestBool2bool(t *testing.T) {
  a := assert.New(t)

  b := false
  B := Bool(true)
  a.NoError(B.ToValue(&b, dummyName))

  a.Equal(true, b)
}

func TestBool2interface(t *testing.T) {
  a := assert.New(t)

  var b interface{}
  B := Bool(true)
  a.NoError(B.ToValue(&b, dummyName))

  a.Equal(true, b)
}

func TestBool2Value(t *testing.T) {
  a := assert.New(t)

  b := false
  v := reflect.ValueOf(&b)

  B := Bool(true)
  if !a.NoError(B.ToValue(v, dummyName)) {
    return
  }

  a.Equal(true, b)
}
