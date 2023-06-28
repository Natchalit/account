package accounts

import "github.com/Natchalit/account/src/functions"

// ค่าใช้จ่าย (x% ของยอดขาย) = (ค่าใช้จ่าย / ยอดขาย)*100
func CompareSalesWithExpenses(expenses, sale, decimal float64) float64 {
	return functions.DecimalPoint((expenses/sale)*100, decimal)
}
