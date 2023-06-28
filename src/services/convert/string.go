package convert

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gofrs/uuid"
)

var STRING string

func String(s interface{}) string {
	pc := ProcessConvertStringPtr(s)
	if pc == nil {
		return STRING
	}
	return *pc
}

func StringPtr(s interface{}) *string {
	pc := ProcessConvertStringPtr(s)
	if pc == nil {
		return &STRING
	}

	return pc
}

func ProcessConvertStringPtr(s interface{}) *string {

	if s == nil {
		return nil
	}

	// Is pointer?
	if reflect.TypeOf(s).Kind() == reflect.Pointer {
		vs := reflect.ValueOf(s)
		// Has value?
		if vs.IsNil() {
			return nil
		}
		s = reflect.Indirect(vs).Interface()
	}

	// string
	if v, ok := s.(string); ok {
		return &v
	}

	if vs, ok := s.(uuid.UUID); ok {
		if vs == uuid.Nil {
			return &STRING
		}
		vx := vs.String()
		return &vx
	}

	var vs string
	switch reflect.TypeOf(s).Kind() {
	case reflect.Float32:
		vs = strconv.FormatFloat(Float64(s.(float32)), 'f', -1, 64)
	case reflect.Float64:
		vs = strconv.FormatFloat(Float64(s), 'f', -1, 64)
	default:
		vs = fmt.Sprintf(`%f`, s)
	}

	return &vs
}
