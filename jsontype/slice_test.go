package jsontype

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestSlice_String(t *testing.T) {
  s0 := []Type{String("str"), NumberWithInt(11), Bool(true), Null{}}
  s := Slice(s0)

  a := assert.New(t)
  a.Equal("[str 11 true nil]", s.String())
}

func TestSlice_ToValue_2Interface(t *testing.T) {
  a := assert.New(t)

  s0 := []Type{String("str"), NumberWithInt(11), Bool(true), Null{}}

  s := Slice(s0)

  var sn interface{}

  if !a.NoError(s.ToValue(&sn, dummyName)) {
    return
  }

  slice,ok := sn.([]interface{})
  if !a.Equalf(true, ok, "[]Type expected") {
    return
  }

  a.Equalf(len(s0), len(slice), "len is not equal")

  a.Equal(s0[0].String(), slice[0])
  a.Equal(s0[1], slice[1])
  a.Equal(bool(s0[2].(Bool)), slice[2])
  a.Equal(nil, slice[3])
}

func TestSlice_ToValue_2Array(t *testing.T) {
  a := assert.New(t)

  s0 := []Type{String("str1"), String("str2"), String("str3"), String("str4")}
  s := Slice(s0)

  var sn [2]string
  if !a.NoError(s.ToValue(&sn, dummyName)) {
    return
  }
  a.Equal(s0[0].String(), sn[0])
  a.Equal(s0[1].String(), sn[1])

  s5 := [5]string{"1", "2", "3", "4", "5"}
  if !a.NoError(s.ToValue(&s5, dummyName)) {
    return
  }
  a.Equal(s0[0].String(), s5[0])
  a.Equal(s0[1].String(), s5[1])
  a.Equal(s0[2].String(), s5[2])
  a.Equal(s0[3].String(), s5[3])
  a.Equal("", s5[4])

}

func TestSlice_ToValue_2Slice(t *testing.T) {
  a := assert.New(t)

  s0 := []Type{String("str1"), String("str2"), String("str3"), String("str4")}
  s := Slice(s0)

  var sn []string
  if !a.NoError(s.ToValue(&sn, dummyName)) {
    return
  }
  a.Equal(len(s0), len(sn), "len()")

  a.Equal(s0[0].String(), sn[0])
  a.Equal(s0[1].String(), sn[1])
  a.Equal(s0[2].String(), sn[2])
  a.Equal(s0[3].String(), sn[3])

  var slice Slice
  if !a.NoError(s.ToValue(&slice, dummyName)) {
    return
  }
  a.Equal(s, slice)
}

