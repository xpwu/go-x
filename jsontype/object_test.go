package jsontype

import (
  "bytes"
  "encoding/json"
  "fmt"
  "github.com/stretchr/testify/assert"
  "reflect"
  "testing"
)

var object Object

func init() {
	object = Object{{Key: "stringKey", Value: String("value1"), Tips: "this is a string"},
		{Key: "numberKey", Value: NumberWithInt(24), Tips: "this is a number"},
		{Key: "nullKey", Value: Null{}, Tips: "this is a null"},
		{Key: "sliceKey", Value: Slice{String("str"), NumberWithInt(11), Bool(true), Null{}}, Tips: "this is a slice"},
	}
	object1 := Object{{Key: "stringKey", Value: String("value1"), Tips: "this is a string"},
		{Key: "numberKey", Value: NumberWithInt(24), Tips: "this is a number"},
		{Key: "nullKey", Value: Null{}, Tips: "this is a null"},
		{Key: "sliceKey", Value: Slice{String("str"), NumberWithInt(11), Bool(true), Null{}}, Tips: "this is a slice"},
	}
	object = append(object, objectElem{Key: "structKey", Value: object1, Tips: "this is a struct"})
}

func TestObject_String(t *testing.T) {
	j := object.String()
	//t.Log(j)

	buffer := bytes.Buffer{}
	err := json.Compact(&buffer, []byte(j))

	a := assert.New(t)
	if !a.NoError(err) {
		return
	}

	v := "{\"stringKey-tips\":\"this is a string\",\"stringKey\":\"value1\",\"numberKey-tips\":\"this is a number\",\"numberKey\":24,\"nullKey-tips\":\"this is a null\",\"nullKey\":null,\"sliceKey-tips\":\"this is a slice\",\"sliceKey\":[\"str\",11,true,null],\"structKey-tips\":\"this is a struct\",\"structKey\":{\"stringKey-tips\":\"this is a string\",\"stringKey\":\"value1\",\"numberKey-tips\":\"this is a number\",\"numberKey\":24,\"nullKey-tips\":\"this is a null\",\"nullKey\":null,\"sliceKey-tips\":\"this is a slice\",\"sliceKey\":[\"str\",11,true,null]}}"
	//t.Log(buffer.String())
	a.Equal(v, buffer.String())
}

func TestObject_Unmarshal_2Interface(t *testing.T) {
	object = append(object, objectElem{
		Key:   "boolKey",
		Value: Bool(true),
		Tips:  "",
	})

	a := assert.New(t)
	var i interface{}
	err := object.Unmarshal(&i, dummyName)
	if !a.NoError(err) {
		return
	}

	m, ok := i.(map[string]interface{})
	if !a.Equalf(true, ok, "map[string]Type expected") {
		return
	}
	if !a.Equalf(6, len(m), "len(map[string]Type) expected") {
		return
	}

	m, ok = m["structKey"].(map[string]interface{})
	if !a.Equalf(true, ok, "map[string]Type expected") {
		return
	}
	if !a.Equalf(4, len(m), "len(map[structKey]) expected") {
		return
	}

	//t.Log(i)
}

func TestObject_Unmarshal_2Struct(t *testing.T) {
	a := assert.New(t)

	s := struct {
		StringKey string `test:"stringKey"`
		NumberKey *Number `test:"numberKey"`
		NullKey   *int   `test:"nullKey"`
		SliceKey  []interface{} `test:"sliceKey"`
		StructKey struct {
			StringKey string `test:"stringKey"`
		} `test:"structKey"`
	}{}

  err := object.Unmarshal(&s, func(tag reflect.StructTag) (name string) {
    return tag.Get("test")
  })
  if !a.NoError(err) {
    return
  }

  a.Equal("{value1 24 <nil> [str 11 true <nil>] {value1}}", fmt.Sprint(s))
}

func TestObject_Unmarshal_2Map(t *testing.T) {
	a := assert.New(t)

	m := map[string]interface{}{}
	s0 := ""
	m["stringKey"] = &s0
	n1 := 5
	m["numberKey"] = &n1
	var s3 []interface{}
	m["sliceKey"] = &s3
	st := struct {
		StringKey string `test:"stringKey"`
	}{}
	m["structKey"] = &st

	err := object.Unmarshal(&m, func(tag reflect.StructTag) (name string) {
		return tag.Get("test")
	})
	if !a.NoError(err) {
		return
	}

	a.Equal(24, n1)
	a.Equal("value1", s0)
	a.Equal("value1", st.StringKey)
	a.Equal(4, len(s3), "len(slice) error")
	a.Equal("11", s3[1].(Number).String())
	a.True(s3[2].(bool))
}
