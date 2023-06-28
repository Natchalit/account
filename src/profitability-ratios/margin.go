package profitabilityratios

import "github.com/Natchalit/account/src/functions"

// อัตรากำไรขั้นต้น = (กำไรขั้นต้น ÷ ยอดขาย) x 100 (%)
func Margin(grossProfit, sales, decimal float64) float64 {
	return functions.DecimalPoint((grossProfit/sales)*100, decimal)
}
