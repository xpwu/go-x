package jsontype

import (
  "github.com/stretchr/testify/assert"
  "reflect"
  "testing"
)

func TestBool_Unmarshal_2bool(t *testing.T) {
  a := assert.New(t)

  b := false
  B := Bool(true)
  a.NoError(B.Unmarshal(&b, dummyName))

  a.Equal(true, b)
}

func TestBool_Unmarshal_2interface(t *testing.T) {
  a := assert.New(t)

  var b interface{}
  B := Bool(true)
  a.NoError(B.Unmarshal(&b, dummyName))

  a.Equal(true, b)
}

func TestBool_Unmarshal_2Value(t *testing.T) {
  a := assert.New(t)

  b := false
  v := reflect.ValueOf(&b)

  B := Bool(true)
  if !a.NoError(B.Unmarshal(v, dummyName)) {
    return
  }

  a.Equal(true, b)
}

func TestBool_Unmarshal_2Ptr(t *testing.T) {
  a := assert.New(t)
  
  var b *bool
  B := Bool(true)
  a.NoError(B.Unmarshal(&b, dummyName))

  a.Equal(true, *b)
}
