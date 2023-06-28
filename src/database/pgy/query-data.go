package pgy

import (
	"database/sql"
	"strings"

	"github.com/Natchalit/account/src/services/logx"
	"github.com/Natchalit/account/src/services/sqly"
)

func (db *CusSql) QueryData(query string, arg ...interface{}) (*sqly.ResponseData, error) {

	defer db.DB.Close()

	rows, ex := db.DB.Query(query)
	if ex != nil {
		return nil, ex
	}

	return scanData(rows)
}

func scanData(rows *sql.Rows) (*sqly.ResponseData, error) {

	defer rows.Close()

	// get cols name
	cols, _ := rows.Columns()
	rowx := []sqly.Map{}
	colsTypex := []sqly.ColumnType{}
	chkCols := map[string]bool{}
	// map data
	for rows.Next() {

		values := make([]interface{}, len(cols))
		row := make(map[string]interface{})

		for i := range cols {
			values[i] = new(interface{})
		}

		colsTypes, _ := rows.ColumnTypes()
		colsTypem := map[string]string{}

		for _, v := range colsTypes {
			if ok := chkCols[v.Name()]; !ok {

				colTypex := sqly.ColumnType{
					Name:             v.Name(),
					DatabaseTypeName: strings.ToUpper(v.DatabaseTypeName()),
				}
				colsTypex = append(colsTypex, colTypex)
				colsTypem[colTypex.Name] = colTypex.DatabaseTypeName
				chkCols[v.Name()] = true
			}
		}
		// scan
		if ex := rows.Scan(values...); ex != nil {
			logx.Errorf(ex.Error())
		}

		for i, column := range cols {
			val := *(values[i].(*interface{}))
			row[column] = val
		}

		rowx = append(rowx, row)
	}

	res := sqly.ResponseData{
		Rows:        rowx,
		Columns:     cols,
		ColumnTypes: colsTypex,
	}

	return &res, nil
}
