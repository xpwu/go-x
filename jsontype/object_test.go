package jsontype

import (
  "bytes"
  "encoding/json"
  "github.com/stretchr/testify/assert"
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