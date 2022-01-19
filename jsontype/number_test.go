package jsontype

import (
  "encoding/json"
  "github.com/stretchr/testify/assert"
  "math"
  "testing"
)

func TestNumber_Int(t *testing.T) {
  a := assert.New(t)
  n := NumberWithStr("5")
  v,err := n.Int64()
  if !a.NoError(err) {
    return
  }
  a.Equal(5, int(v))

  n = NumberWithInt(23)
  v,err = n.Int64()
  if !a.NoError(err) {
    return
  }
  a.Equal(23, int(v))

  n = NumberWithFloat(28.9)
  v,err = n.Int64()
  if !a.NoError(err) {
    return
  }
  a.Equal(28, int(v))

}

func TestNumber_Uint64(t *testing.T) {
  a := assert.New(t)
  n := NumberWithStr("5")
  v,err := n.Uint64()
  if !a.NoError(err) {
    return
  }
  a.Equal(5, int(v))

  n = NumberWithInt(23)
  v,err = n.Uint64()
  if !a.NoError(err) {
    return
  }
  a.Equal(23, int(v))

  n = NumberWithFloat(28.9)
  v,err = n.Uint64()
  if !a.NoError(err) {
    return
  }
  a.Equal(28, int(v))

}

func TestNumber_String(t *testing.T) {
  a := assert.New(t)
  n := NumberWithStr("5")
  v := n.String()
  a.Equal("5", v)

  n = NumberWithInt(23)
  v = n.String()
  a.Equal("23", v)

  n = NumberWithFloat(28.9)
  v = n.String()
  a.Equal("28.9", v)

}

func TestNumber_Float(t *testing.T) {
  a := assert.New(t)
  n := NumberWithStr("5.3333")
  v,err := n.Float()
  if !a.NoError(err) {
    return
  }
  a.LessOrEqual(math.Abs(5.3333-v), math.SmallestNonzeroFloat64)

  n = NumberWithInt(23)
  v,err = n.Float()
  if !a.NoError(err) {
    return
  }
  a.LessOrEqual(math.Abs(23-v), math.SmallestNonzeroFloat64)

  n = NumberWithFloat(28.9)
  v,err = n.Float()
  if !a.NoError(err) {
    return
  }
  a.LessOrEqual(math.Abs(28.9-v), math.SmallestNonzeroFloat64)

}

func TestNumber_ToValue_2Interface(t *testing.T) {
  a := assert.New(t)
  n := NumberWithStr("5")

  var i interface{}
  if !a.NoError(n.ToValue(&i, dummyName)) {
    return
  }

  m,err := i.(Number).Int64()
  if !a.NoError(err) {
    return
  }
  a.Equal(5, int(m))
}

func TestNumber_ToValue_2Number(t *testing.T) {
  a := assert.New(t)
  n := NumberWithStr("5")

  n2 := NumberWithInt(7)
  if !a.NoError(n.ToValue(&n2, dummyName)) {
    return
  }

  a.Equal(n, n2)
}

func TestNumber_ToValue_2JsonNumber(t *testing.T) {
  a := assert.New(t)
  n := NumberWithStr("5")

  n2 := json.Number("10")
  if !a.NoError(n.ToValue(&n2, dummyName)) {
    return
  }

  a.Equal(n.String(), n2.String())
}

func TestNumber_ToValue_2IntXXX(t *testing.T) {
  a := assert.New(t)
  n0 := NumberWithStr("5")

  var n int = 0
  var n8 int8 = 0
  var n16 int16 = 0
  var n32 int32 = 0
  var n64 int64 = 0

  if a.NoError(n0.ToValue(&n, dummyName)) {
    a.Equal(5, n)
  }
  if a.NoError(n0.ToValue(&n8, dummyName)) {
    a.Equal(int8(5), n8)
  }
  if a.NoError(n0.ToValue(&n16, dummyName)) {
    a.Equal(int16(5), n16)
  }
  if a.NoError(n0.ToValue(&n32, dummyName)) {
    a.Equal(int32(5), n32)
  }
  if a.NoError(n0.ToValue(&n64, dummyName)) {
    a.Equal(int64(5), n64)
  }

}

func TestNumber_ToValue_2UintXXX(t *testing.T) {
  a := assert.New(t)
  n0 := NumberWithStr("5")

  var n uint = 0
  var n8 uint8 = 0
  var n16 uint16 = 0
  var n32 uint32 = 0
  var n64 uint64 = 0

  if a.NoError(n0.ToValue(&n, dummyName)) {
    a.Equal(uint(5), n)
  }
  if a.NoError(n0.ToValue(&n8, dummyName)) {
    a.Equal(uint8(5), n8)
  }
  if a.NoError(n0.ToValue(&n16, dummyName)) {
    a.Equal(uint16(5), n16)
  }
  if a.NoError(n0.ToValue(&n32, dummyName)) {
    a.Equal(uint32(5), n32)
  }
  if a.NoError(n0.ToValue(&n64, dummyName)) {
    a.Equal(uint64(5), n64)
  }

}

func TestNumber_ToValue_2FloatXX(t *testing.T) {
  a := assert.New(t)
  n0 := NumberWithStr("5")

  var n float32 = 0
  var n8 float64 = 0

  if a.NoError(n0.ToValue(&n, dummyName)) {
    a.LessOrEqual(math.Abs(float64(5-n)), math.SmallestNonzeroFloat64)
  }
  if a.NoError(n0.ToValue(&n8, dummyName)) {
    a.LessOrEqual(math.Abs(5-n8), math.SmallestNonzeroFloat64)
  }

}

