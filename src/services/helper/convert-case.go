package helper

import (
	"github.com/iancoleman/strcase"
)

func ToSnake(s string) string {
	return strcase.ToSnake(s)
}

func ToPascal(s string) string {
	return strcase.ToCamel(s)
}

func ToCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
