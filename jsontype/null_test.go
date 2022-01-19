package jsontype

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestNull_Unmarshal_2Interface(t *testing.T) {
  a := assert.New(t)

  n := Null{}

  var i interface{}
  if a.NoError(n.Unmarshal(&i, dummyName)) {
    a.Nil(i)
  }
}

func TestNull_Unmarshal_2Map(t *testing.T) {
  a := assert.New(t)

  n := Null{}

  m := make(map[int]int)
  if a.NoError(n.Unmarshal(&m, dummyName)) {
    a.Nil(m)
  }
}

func TestNull_Unmarshal_2Slice(t *testing.T) {
  a := assert.New(t)

  n := Null{}

  m := make([]int, 0)
  if a.NoError(n.Unmarshal(&m, dummyName)) {
    a.Nil(m)
  }
}

func TestNull_Unmarshal_2Ptr(t *testing.T) {
  a := assert.New(t)

  n := Null{}

  var m *interface{}
  if a.NoError(n.Unmarshal(&m, dummyName)) {
    a.Nil(m)
  }
}

