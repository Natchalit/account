package profitabilityratios

import "github.com/Natchalit/account/src/functions"

// อัตรากำไรสุทธิ = กำไรสุทธิ ÷ ยอดขาย x 100 (%) [กำไรสุทธิ = กำไรหลักหักค่าใช้จ่ายต่างๆ ดอกเบี้ย และภาษีเงินได้แล้ว]
func NetProfitMargin(netProfit, sales, decimal float64) float64 {
	return functions.DecimalPoint((netProfit/sales)*100.0, decimal)
}
