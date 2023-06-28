package functions

import "math"

// ปรับจุดทศนิยม (ค่าที่ต้องการเปลี่ยน,จำนวนจุดทศนิยมที่ต้องการ)
func DecimalPoint(num float64, mdp float64) float64 {

	if num == 0.0 {
		return 0
	}

	dp := 10.0

	if mdp != 0 {
		dp = Pow(dp, mdp, 0)
		return math.Round(num*dp) / dp
	} else {
		return math.Round(num*1.0) / 1.0
	}

}
