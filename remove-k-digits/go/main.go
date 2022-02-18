package main

import (
	"fmt"
)

// min finds the smallest digit in a string of decimal digits
func min(s string, start, end int) int {
	minIndex := start

	for i := start; i < end; i++ {
		if s[i] < s[minIndex] {
			minIndex = i
		}
	}

	return minIndex
}

// trimLeadingZeros removes the leading zeros from a string
func trimLeadingZeros(s string) string {
	firstNonZero := len(s) - 1

	for i := 0; i < len(s); i++ {
		if string(s[i]) != "0" {
			firstNonZero = i
			break
		}
	}

	return s[firstNonZero:]
}

// removeKdigits removes K digits from a string of decimal digits to get
// the smallest possible number
func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}

	discarded := 0
	smallest := ""
	minIdx := -1

	for {
		start := minIdx + 1
		end := start + (k - discarded) + 1

		minIdx = min(num, start, end)
		digit := string(num[minIdx])
		smallest += digit

		discarded += minIdx - start
		if discarded == k {
			smallest += num[end:]
			break
		}

		if len(smallest) == len(num)-k {
			break
		}
	}

	smallest = trimLeadingZeros(smallest)

	return smallest
}

func main() {
	smallest := removeKdigits("7401625293823", 5)
	fmt.Printf("smallest: %v\n", smallest)

	smallest = removeKdigits("654398166842", 6)
	fmt.Printf("smallest: %v\n", smallest)

	smallest = removeKdigits("1432219", 3)
	fmt.Printf("smallest: %v\n", smallest)

	smallest = removeKdigits("10200", 1)
	fmt.Printf("smallest: %v\n", smallest)

	smallest = removeKdigits("10", 2)
	fmt.Printf("smallest: %v\n", smallest)

	smallest = removeKdigits("10", 1)
	fmt.Printf("smallest: %v\n", smallest)

	smallest = removeKdigits("112", 1)
	fmt.Printf("smallest: %v\n", smallest)

	smallest = removeKdigits("10001", 1)
	fmt.Printf("smallest: %v\n", smallest)

	smallest = removeKdigits("100", 1)
	fmt.Printf("smallest: %v\n", smallest)
}
