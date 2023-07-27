package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter number only: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Your input is:", text)

	fmt.Println("Length more than six char:", lengthCheck(text))
	fmt.Println("Not increase or decrease three digit in order", incAndDecNum(text))
	fmt.Println("Not repeat three digit in order:", numberRepeatThreeDigit(text))
	fmt.Println("Not repeat two digit in three set", numberRepeatTwoTimeAndThreeSet(text))

	// testCase_1 := "1255567"
	// testCase_2 := "5734514"
	// testCase_3 := "162987"

	// testCase_4 := "111111"
	// testCase_5 := "5718880"
	// testCase_6 := "1927333"
	// testCase_4 := "11232246"
	// testCase_5 := "5718823492834902834982903400000"
	// testCase_6 := "192733"

	// testCase_7 := "112231344"
	// testCase_8 := "3144556647"
	// testCase_9 := "17887766"

	// fmt.Println(lengthCheck(testCase_1))

	// fmt.Println(incAndDecNum(testCase_1))
	// fmt.Println(incAndDecNum(testCase_2))
	// fmt.Println(incAndDecNum(testCase_3))

	// fmt.Println(numberRepeatThreeDigit(testCase_4))
	// fmt.Println(numberRepeatThreeDigit(testCase_5))
	// fmt.Println(numberRepeatThreeDigit(testCase_6))

	// fmt.Println(numberRepeatTwoTimeAndThreeSet(testCase_7))
	// fmt.Println(numberRepeatTwoTimeAndThreeSet(testCase_8))
	// fmt.Println(numberRepeatTwoTimeAndThreeSet(testCase_9))
}

func lengthCheck(s string) bool {
	charLength := len(s)
	if charLength >= 6 {
		return true
	} else {
		return false
	}
}

func incAndDecNum(s string) bool {
	validatePass := true
	threshold := 3
	repeatCountInc := 1
	repeatCountDec := 1
	stringSlice := strings.Split(s, "") // it is string slice
	var intSlice []int
	// fmt.Println(stringSlice)

	//convert element inside slice from string to int for easier manipulation string
	for _, value := range stringSlice {
		number, _ := strconv.Atoi(value)
		intSlice = append(intSlice, number)
	}

	// check increase first
	for i := 1; i < len(intSlice); i++ {
		if intSlice[i] == (intSlice[i-1] + 1) {
			repeatCountInc++

			if threshold == repeatCountInc {
				validatePass = false
				break
			}
		} else {
			repeatCountInc = 1
		}

		// fmt.Println("Currecn is:", intSlice[i])
		// fmt.Println("Previous is:", intSlice[i-1])
		// fmt.Println("Repeat count is:", repeatCountInc)
		// fmt.Println("----------------------")

	}

	// check decrease then
	for i := 1; i < len(intSlice); i++ {
		if intSlice[i] == (intSlice[i-1] - 1) {
			repeatCountDec++
			if threshold == repeatCountDec {
				validatePass = false
				break
			}

		} else {
			repeatCountDec = 1
		}

	}

	// fmt.Println(`Repeat Count Increase :`, repeatCountInc)
	// fmt.Println(`Repeat Count Decrease :`, repeatCountDec)

	return validatePass

}

func numberRepeatThreeDigit(s string) bool {
	validatePass := true
	threshold := 3
	repeatCount := 1

	stringSlice := strings.Split(s, "")

	for i := 1; i < len(stringSlice); i++ {
		if stringSlice[i] == stringSlice[i-1] {
			repeatCount++
			if repeatCount == threshold {
				validatePass = false
				break
			}
		} else {
			repeatCount = 1
		}
	}

	return validatePass
}

func numberRepeatTwoTimeAndThreeSet(s string) bool {
	validatePass := true

	stringSlice := strings.Split(s, "")

	for i := 0; i < len(stringSlice); i++ {
		if i+6 > len(stringSlice) {
			validatePass = true
			break
		}

		if stringSlice[i] == stringSlice[i+1] && stringSlice[i+2] == stringSlice[i+3] && stringSlice[i+4] == stringSlice[i+5] {
			validatePass = false
			break
		}
	}

	return validatePass
}
