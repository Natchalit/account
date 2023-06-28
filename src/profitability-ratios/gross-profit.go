package profitabilityratios

import "github.com/Natchalit/account/src/functions"

// กำไรขั้นต้น = รายได้จากการขาย - ต้นทุนขายและบริการ
func GrossProfit(salesRevenue, costOfSalesAndServices, decimal float64) float64 {
	return functions.Minus(salesRevenue, costOfSalesAndServices, decimal)
}
