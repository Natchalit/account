package accounts

import "github.com/Natchalit/account/src/functions"

// รวมรายจ่าย
func Expenses(nums []float64, decimal float64) float64 {
	return functions.SumTotal(nums, decimal)
}
