package sqly

type Map map[string]interface{}

type ColumnType struct {
	Name             string `json:"name"`
	DatabaseTypeName string `json:"database_type_name"`
}

type ResponseData struct {
	Rows        []Map        `json:"data"`
	Columns     []string     `json:"columns"`
	ColumnTypes []ColumnType `json:"columns_type"`
}
