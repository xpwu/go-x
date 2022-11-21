package jsontype

import (
	"fmt"
	"reflect"
)

type Kind uint

const (
  NullK Kind = iota
	NumberK
	StringK
	BoolK
	SliceK
	ObjectK
)

type Type interface {
	fmt.Stringer
	Kind() Kind
	// 把值 Unmarshal为go类型的value
	// 类似 json.Unmarshal(data []byte, value interface{})
	// name: tag到name的映射关系。
	// 如果value有struct类型时，field key的取值为：1、name func 的结果；2、如果为""，则取filed Name的值
	// 如果value 为 interface{}，转换关系如下：
	// Object -> Map[Key]interface{}, Tips丢弃
	// Null -> nil
	// Bool -> bool
	// Number 不转换
	// Slice -> []interface{}
	// String -> string
	Unmarshal(value interface{}, name func(tag reflect.StructTag) (name string)) error

	// a.Include(c) c中的字段，在a中都存在，并且类型相同, a >= c
	Include(other Type) bool
	IncludeErr(other Type, path string) error
}

type Null struct {}
type Number struct {str *string; i *int64; f *float64}
type String string
type Bool bool
type Slice []Type
// 为了有序性，所以没有使用map
type Object []objectElem

type objectElem struct {Key string; Value Type; Tips string}

