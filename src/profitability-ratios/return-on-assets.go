package profitabilityratios

import "github.com/Natchalit/account/src/functions"

// อัตราผลตอบแทนจากสินทรัพย์ = กำไรสุทธิ ÷ สินทรัพย์รวม (เฉลี่ย) x 100 (%)
func ReturnOnAssets(netProfit, TotalAssets, decimal float64) float64 {
	return functions.DecimalPoint((netProfit/TotalAssets)*100.0, decimal)
}
