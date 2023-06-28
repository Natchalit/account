package sqly

import (
	"github.com/Natchalit/account/src/services/convert"
	"github.com/Natchalit/account/src/services/helper"
)

func (m *Map) GetString(key string) string {
	if m == nil {
		return convert.STRING
	}
	return convert.String(m.getCase(key))
}

func (m *Map) GetFloat(key string) float64 {
	if m == nil {
		return convert.FLOAT64
	}
	return convert.Float64(m.getCase(key))
}

func (m *Map) GetInt(key string) int64 {
	if m == nil {
		return convert.INT64
	}
	return convert.Int64(m.getCase(key))
}

func (m *Map) GetBool(key string) bool {
	if m == nil {
		return false
	}
	return convert.Bool(m.getCase(key))
}

func (m *Map) GetStringPtr(key string) *string {
	if m == nil {
		return &convert.STRING
	}
	return convert.StringPtr(m.getCase(key))
}

func (m *Map) GetFloatPtr(key string) *float64 {
	if m == nil {
		return &convert.FLOAT64
	}
	return convert.Float64Ptr(m.getCase(key))
}

func (m *Map) GetIntPtr(key string) *int64 {
	if m == nil {
		return &convert.INT64
	}
	return convert.IntPtr(m.getCase(key))
}

func (m *Map) GetBoolPtr(key string) *bool {
	if m == nil {
		return &convert.FALSE
	}
	return convert.BoolPtr(m.getCase(key))
}

func (m *Map) getCase(col string) interface{} {
	if helper.IsSnakeCase(col) {
		if v, ok := (*m)[col]; ok {
			return v
		}
	} else {
		return m.getCase(helper.ToSnake(col))
	}
	return nil
}
