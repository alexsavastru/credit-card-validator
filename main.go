package main

import (
	"fmt"
	"strings"
)

type Bank struct {
	Name   string
	Prefix string
}

func DetectBank(cardNumber string, banks []Bank) *Bank {

	if cardNumber == "" {
		return nil
	}

	for i := range banks {
		if strings.HasPrefix(cardNumber, banks[i].Prefix) {
			return &banks[i]
		}
	}

	return nil
}

func LuhnCheck(cardNumber string) bool {
	if cardNumber == "" {
		return false
	}

	sum := 0
	double := false
	digits := 0

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := cardNumber[i] - '0'
		if digit > 9 {
			return false
		}

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += int(digit)
		double = !double
		digits++
	}

	return digits > 0 && sum%10 == 0
}

func main() {
	banks := []Bank{
		{Name: "Lunar Bank", Prefix: "4000"},
		{Name: "Mars Credit Union", Prefix: "5000"},
		{Name: "Alfa Credit", Prefix: "8000"},
	}
	fmt.Println(DetectBank("4000123456789017", banks)) // → &Bank{Name: "Lunar Bank", ...}
	fmt.Println(DetectBank("1234567890123456", banks)) // → nil
	fmt.Println(DetectBank("", banks))

	fmt.Println(LuhnCheck("4000123456789017")) // → true  (валиден)
	fmt.Println(LuhnCheck("79927398713"))      // → true
	fmt.Println(LuhnCheck("1234567890123456")) // → false (невалиден по Луне)
	fmt.Println(LuhnCheck(""))                 // → false
	fmt.Println(LuhnCheck("4000abc"))          // → false (нецифры)
}
