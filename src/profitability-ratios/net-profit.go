package profitabilityratios

import "github.com/Natchalit/account/src/functions"

// กำไรสุทธิ = กำไรก่อนภาษี - ภาษี
func NetProfit(profitBeforeTax, tax, decimal float64) float64 {
	return functions.Minus(profitBeforeTax, tax, decimal)
}
