package jsontype

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestToValueInterface(t *testing.T) {
  a := assert.New(t)

  n := Null{}

  var i interface{}
  if a.NoError(n.ToValue(&i, dummyName)) {
    a.Nil(i)
  }
}

func TestToValueMap(t *testing.T) {
  a := assert.New(t)

  n := Null{}

  m := make(map[int]int)
  if a.NoError(n.ToValue(&m, dummyName)) {
    a.Nil(m)
  }
}

func TestToValueSlice(t *testing.T) {
  a := assert.New(t)

  n := Null{}

  m := make([]int, 0)
  if a.NoError(n.ToValue(&m, dummyName)) {
    a.Nil(m)
  }
}

func TestToValuePtr(t *testing.T) {
  a := assert.New(t)

  n := Null{}

  var m *interface{}
  if a.NoError(n.ToValue(&m, dummyName)) {
    a.Nil(m)
  }
}

