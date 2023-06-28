package accounts

// ref https://greedisgoods.com/fv-%E0%B8%84%E0%B8%B7%E0%B8%AD-future-value/
import "github.com/Natchalit/account/src/functions"

// มูลค่าเงินในอนาคตหรือใช้หาดอกเบี้ยเงินกู้ = เงินต้น * (1 + [ดอกเบี้ย %])^จำนวนปี
func FutureValue(fv, i, n, decimal float64) float64 {
	return functions.DecimalPoint(fv*(functions.Pow((1+functions.Percent(i, 2)), n, decimal)), decimal)
}

// ผลตอบแทนตามรอบที่กำหนด = เงินลงทุน / (1+ [ผลตอบแทนเป็น % ])^จำนวนครั้งที่จ่าย
func FutureValuePeriod(fv, i, n, decimal float64) float64 {
	return functions.DecimalPoint(fv/(functions.Pow((1+functions.Percent(i, 2)), n, decimal)), decimal)
}

// เงินเฟ้อ = เงินต้น * (1 - [อัตตราการเฟ้อ %])^จำนวนปี
func FutureValueInflation(fv, i, n, decimal float64) float64 {
	return functions.DecimalPoint(fv*(functions.Pow((1-functions.Percent(i, 2)), n, decimal)), decimal)
}
