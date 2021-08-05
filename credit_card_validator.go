// Package creditcardvalidator implements Valid function for validation
package creditcardvalidator

import (
	"strconv"
	"strings"
)

// Valid validates given string of credit number with luhn algorithm
func Valid(s string) bool {
	r := strings.NewReplacer(" ", "", "-", "", "_", "")
	s = r.Replace(s)

	// Check if string contains non numeric characters
	_, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	// Reject if creadit card is equal to 0000 0000 0000 0000 or length is not 16
	if len(s) != 16 || s == "0000000000000000" {
		return false
	}

	var is []int
	for _, v := range s {
		// Check if rune is a number's unicode
		if v >= 48 && v <= 57 {
			// Convert rune to string and then convert it to integer
			i, err := strconv.Atoi(string(v))
			if err != nil {
				return false
			}
			is = append(is, i)
		}
	}

	// Luhn Algorithm
	total := 0
	for i, v := range is {
		if i%2 == len(is)%2 {
			if v*2 > 9 {
				total += (v * 2) - 9
			} else {
				total += v * 2
			}
		} else {
			total += v
		}
	}
	if total%10 == 0 {
		return total%10 == 0
	}

	return false
}
