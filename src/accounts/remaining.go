package accounts

import "github.com/Natchalit/account/src/functions"

// เงินคงเหลือ = (ตั้งต้น + รายรับ) - รายจ่าย
func Remaining(carriedOver, decimal float64, income, expenses []float64) float64 {
	return functions.DecimalPoint((carriedOver+Income(income, decimal))-Expenses(expenses, decimal), decimal)
}
