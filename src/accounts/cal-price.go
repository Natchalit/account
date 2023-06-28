package accounts

import "github.com/Natchalit/account/src/functions"

// ราคาสินค้าทั้งหมด = ราคาต่อหน่วย * จำนวนหน่วยสินค้า (ไม่รวม vat)
func CalPrice(pricePerItem, item, decimal float64) float64 {
	return functions.Multiply(pricePerItem, item, decimal)
}
