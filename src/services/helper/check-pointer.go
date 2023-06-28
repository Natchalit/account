package helper

import "reflect"

func IsNil(s any) bool {
	if s == nil {
		return true
	}
	switch reflect.TypeOf(s).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(s).IsNil()
	}
	return false
}
