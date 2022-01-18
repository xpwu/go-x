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
}

