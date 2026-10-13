package raindrops

import "strconv"

func Convert(number int) string {
	result := ""

	if number % 3 == 0 {
		result += "Pling"
	}
	if number % 5 == 0 {
		result += "Plang"
	}
	if number % 7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(number)
	}

	return result
}


// func Convert(number int) string {
// 	result := ""
//
// 	var remainder3 = number % 3
// 	var remainder5 = number % 5
// 	var remainder7 = number % 7
//
// 	if remainder3 == 0 {
// 		result = result + "Pling"
// 	}
// 	if remainder5 == 0 {
// 		result = result + "Plang"
// 	}
// 	if remainder7 == 0 {
// 		result = result + "Plong"
// 	}
//
// 	if remainder3 != 0 && remainder5 != 0 && remainder7 != 0 {
// 		result = strconv.Itoa(number)
// 	}
//
// 	return result
// }
//
