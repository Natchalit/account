package profitabilityratios

import "github.com/Natchalit/account/src/functions"

// กำไรก่อนภาษี = กำไรก่อนดอกเบี้ยและภาษี - ดอกเบี้ย
func ProfitBeforeTax(profitBeforeInterestAndTax, interest, decimal float64) float64 {
	return functions.Minus(profitBeforeInterestAndTax, interest, decimal)
}
