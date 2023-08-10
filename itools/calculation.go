package itools

func CalculationArrMax(arrint []int) (max int) {
	max = arrint[0]
	for _, v := range arrint {
		if v > max {
			max = v
		}
	}
	return
}

func CheckInListInt(val int, slist []int) bool {
	for _, v := range slist {
		if v == val {
			return true
		}
	}
	return false
}

func CheckInListString(val string, slist []string) bool {
	for _, v := range slist {
		if v == val {
			return true
		}
	}
	return false
}

func DeleteListString(val string, slist []string) []string {
	var resultlist []string
	for _, v := range slist {
		if v != val {
			resultlist = append(resultlist, v)
		}
	}
	return resultlist
}

func DeleteListint(val int, slist []int) []int {
	var resultlist []int
	for _, v := range slist {
		if v != val {
			resultlist = append(resultlist, v)
		}
	}
	return resultlist
}
