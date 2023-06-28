package accounts

import (
	"github.com/Natchalit/account/src/functions"
)

// ราคาหลังหักส่วนลด = ราคาเต็ม - (ราคาเต็ม * ส่วนลด %)
func Discount(price, precent, decimal float64) float64 {
	return functions.DecimalPoint(price-(price*functions.Percent(precent, 2)), decimal)
}
