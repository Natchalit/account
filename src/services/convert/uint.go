package convert

import (
	"fmt"
	"reflect"
	"strconv"
)

var UINT uint
var UINT8 uint8
var UINT16 uint16
var UINT32 uint32
var UINT64 uint64

func Uint(n interface{}) uint {
	r := ProcessConvertUintPtr(n)
	if r == nil {
		return UINT
	}
	return uint(*r)
}
func Uint8(n interface{}) uint8 {
	r := ProcessConvertUintPtr(n)
	if r == nil {
		return UINT8
	}
	return uint8(*r)
}
func Uint16(n interface{}) uint16 {
	r := ProcessConvertUintPtr(n)
	if r == nil {
		return UINT16
	}
	return uint16(*r)
}
func Uint32(n interface{}) uint32 {
	r := ProcessConvertUintPtr(n)
	if r == nil {
		return UINT32
	}
	return uint32(*r)
}
func Uint64(n interface{}) uint64 {
	r := ProcessConvertUintPtr(n)
	if r == nil {
		return UINT64
	}
	return *r
}

func ProcessConvertUintPtr(n interface{}) *uint64 {
	if n == nil {
		return nil
	}

	// Is pointer?
	if reflect.TypeOf(n).Kind() == reflect.Pointer {
		vn := reflect.ValueOf(n)
		// Has value?
		if vn.IsNil() {
			return nil
		}
		n = reflect.Indirect(vn).Interface()
	}

	vx, ex := strconv.ParseUint(fmt.Sprintf(`%v`, n), 10, 64)
	if ex != nil {
		return nil
	}

	return &vx
}
