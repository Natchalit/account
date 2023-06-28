package accounts

import "github.com/Natchalit/account/src/functions"

// รวมรายรับ
func Income(nums []float64, decimal float64) float64 {
	return functions.SumTotal(nums, decimal)
}
