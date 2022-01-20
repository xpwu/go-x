package jsontype

import (
  "bytes"
  "github.com/stretchr/testify/assert"
  "reflect"
  "testing"
)

func dummyName(tag reflect.StructTag) (name string) {
  return ""
}

var gotype = struct {
  In int `test:"it,this is int"`
  Flt float64 `test:"fl,this is float"`
  Str string
  Sli []string
  Stu struct{
    Uin uint32
    Flt32 float32 `test:"Arr,be covered"`
  }
  Arr [2]int
}{Sli:[]string{"s", "t", "r"}}

var tips = func(tag reflect.StructTag) (key, tips string) {
  t := tag.Get("test")
  ts := bytes.SplitN([]byte(t), []byte(","), 2)
  key = string(ts[0])
  if len(ts) == 2 {
    tips = string(ts[1])
  }
  return
}

func TestFromGoType(t *testing.T) {
  ty := FromGoType(&gotype, tips)

  a := assert.New(t)
  a.Equal(ObjectK, ty.Kind())
  t1 := ty.(Object)
  a.Equal(6, len(t1))
  a.Equal(NumberK, t1[0].Value.Kind())
  a.Equal(NumberK, t1[1].Value.Kind())
  a.Equal(StringK, t1[2].Value.Kind())
  a.Equal(SliceK, t1[3].Value.Kind())
  a.Equal(ObjectK, t1[4].Value.Kind())
  a.Equal(SliceK, t1[5].Value.Kind())
  a.Equal(Slice([]Type{String("s"), String("t"), String("r")}), t1[3].Value.(Slice))
}

func TestFromJson(t *testing.T) {
  ty := FromGoType(gotype, tips)
  a := assert.New(t)

  jsn := "{\"it-tips\":\"this is int\",\"it\":0,\"fl-tips\":\"this is float\",\"fl\":0,\"Str\":\"\",\"Sli\":[\"s\",\"t\",\"r\"],\"Stu\":{\"Uin\":0,\"Arr-tips\":\"be covered\",\"Arr\":0},\"Arr\":[0,0]}"
  tp,err := FromJson([]byte(jsn))
  if !a.NoError(err) {
    return
  }

  a.True(tp.Include(ty))
}

