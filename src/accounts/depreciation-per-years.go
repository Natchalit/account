package accounts

import "github.com/Natchalit/account/src/functions"

// ค่าเสื่อมราคาต่อปี = [ราคาทุนของสินทรัพย์ – ราคาซาก (ถ้ามี)] ÷ อายุการใช้งาน
// สูตรการคิดค่าเสื่อมราคาสูตรนี้ จะเหมาะกับอุปกรณ์ เครื่องจักร ข้าวของเครื่องใช้
func DepreciationPerYears(costPrice, carcassPrice, serviceLife, decimal float64) float64 {
	return functions.DecimalPoint((costPrice-carcassPrice)/serviceLife, decimal)
}
