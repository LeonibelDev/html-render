/*
Luhn algorithm
Ref: https://en.wikipedia.org/wiki/Luhn_algorithm
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// string to int
func stringToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return num
}

// reverse string
func reverse(s string) string {
	strReverse := strings.Split(s, "")
	strReversed := ""

	for i := len(strReverse) - 1; i >= 0; i-- {
		strReversed += strReverse[i]
	}

	return strReversed
}

// validate credit card
func validateCC(ccnumbers string) bool {
	strReversed := reverse(ccnumbers)

	final := 0

	for i := 0; i < len(strReversed); i++ {
		num := stringToInt(string(strReversed[i]))

		if i%2 == 1 {
			num *= 2

			if num > 9 {
				num -= 9
			}
		}

		final += num

	}

	return final%10 == 0
}
