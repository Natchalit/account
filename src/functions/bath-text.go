package functions

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Natchalit/account/src/services/convert"
)

var numToText = map[string]string{
	`0`: "ศูนย์",
	`1`: "หนึ่ง",
	`2`: "สอง",
	`3`: "สาม",
	`4`: "สี่",
	`5`: "ห้า",
	`6`: "หก",
	`7`: "เจ็ด",
	`8`: "แปด",
	`9`: "เก้า",
}

var digitText = map[uint]string{
	0: "",
	1: "สิบ",
	2: "ร้อย",
	3: "พัน",
	4: "หมื่น",
	5: "แสน",
	6: "ล้าน",
}

func BathText(num float64) string {

	if num == 0 {
		return fmt.Sprintf(`%sบาท `, numToText["0"])
	}

	strNum := fmt.Sprintf(`%f`, num)
	numSplit := strings.Split(strNum, `.`)
	naturalNum := numSplit[0]

	// สตางค์
	// nN, _ := strconv.ParseFloat(naturalNum, 64)
	nN := convert.Float64(naturalNum)
	stangV := num - nN                                                // 0.02334
	stangx := DecimalPoint(stangV, 2)                                 // 0.02
	strStangxSplit := strings.Split(fmt.Sprintf(`%.2f`, stangx), `.`) // '02'
	// flStang, _ := strconv.ParseFloat(strStangxSplit[1], 64)
	flStang := convert.Float64(strStangxSplit[1])

	stangRs := RoundStang(flStang)
	stang := ``
	if stangRs == 100 {
		// numx, _ := strconv.ParseUint(naturalNum, 10, 64)
		numx := convert.Uint64(naturalNum)
		numx += convert.Uint64(stangRs)
		naturalNum = fmt.Sprintf(`%d`, numx)
	} else {
		st := strconv.Itoa(stangRs)
		rangeStang := len(st)
		stang = TextBath(rangeStang, st)
	}

	digitNatural := len(naturalNum)
	naturalV := TextBath(digitNatural, naturalNum)

	strRs := ``
	if stang != `` {
		strRs = naturalV + `บาท` + stang + `สตางค์`
	} else {
		strRs = naturalV + `บาทถ้วน`
	}

	fmt.Println(strRs)

	return strRs
}

func _OnMod(modNo, i uint) uint {
	modx := ((modNo - 1) - i) % 7

	if modx > 6 {
		return _OnMod(modx, i)
	}
	return modx
}

func TextBath(digitNatural int, naturalNum string) string {
	strRs := ``
	for i := 0; i < digitNatural; i++ {
		strNaturalNo := ``
		strCall := ``

		strNaturalNo = numToText[string(naturalNum[i])]

		mod := _OnMod(uint(digitNatural), uint(i))
		if mod > 0 {
			strCall = digitText[uint(mod)]
		} else {
			strCall = digitText[uint(digitNatural-1)-uint(i)]
		}

		if strNaturalNo == numToText[`0`] {
			if digitText[mod] == digitText[6] {
				strNaturalNo = ``
				strCall = digitText[mod]
			} else {
				strCall = ``
				strNaturalNo = ``
			}
		}
		if strNaturalNo == numToText[`1`] && strCall == digitText[1] {
			strNaturalNo = ``
		}
		if strNaturalNo == numToText[`1`] && i == digitNatural-1 {
			strNaturalNo = `เอ็ด`
		}
		if strNaturalNo == numToText[`2`] && strCall == digitText[1] {
			strNaturalNo = `ยี่`
		}

		strRs = strRs + strNaturalNo + strCall
	}

	return strRs
}
