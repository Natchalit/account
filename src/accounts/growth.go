package accounts

import "github.com/Natchalit/account/src/functions"

// อัตราการเติบโต = [(ยอดขายล่าสุด-ยอดขายปีที่แล้ว)/ยอดขายปีที่แล้ว] x 100
func Growth(now, lastYear, decimal float64) float64 {
	return functions.DecimalPoint((now-lastYear)/100.0, decimal)
}
