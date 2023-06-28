package accounts

import "github.com/Natchalit/account/src/functions"

// จุดคุ้มทุน = ต้นทุนคงที่ / กำไรส่วนเกินต่อหน่วย
func BreakEvenPoint(fixedCost, surplusProfit, decimal float64) float64 {
	return functions.Divine(fixedCost, surplusProfit, decimal)
}
