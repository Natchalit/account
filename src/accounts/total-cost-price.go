package accounts

import "github.com/Natchalit/account/src/functions"

// ราคาต้นทุนรวมทั้งหมด = (ปริมาณที่ใช้ * ราคา) / ปริมาณทั้งหมด
func TotalCostPrice(amountUsed, price, totalVolume, decimal float64) float64 {
	priceX := functions.Multiply(amountUsed, price, decimal)
	return functions.Divine(priceX, totalVolume, decimal)
}
