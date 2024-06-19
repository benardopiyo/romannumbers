package main

import (
	"os"

	"github.com/01-edu/z01"
)

var romanNumerals = []struct {
	value int
	symbol string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func parseInput(input string) (int, bool) {
	num := 0
	for _, r := range input {
		if r < '0' || r > '9' {
			return 0, false
		}
		num = num*10 + int(r-'0')
	}
	return num, true
}

func toRoman(n int) (string, string) {
	if n <= 0 || n >= 4000 {
		return "", "ERROR: cannot convert to roman digit"
	}

	result := ""
	calculation := ""
	for _, rn := range romanNumerals {
		for n >= rn.value {
			if len(result) > 0 {
				calculation += "+"
			}
			result += rn.symbol
			if rn.value >= 1000 {
				calculation += rn.symbol
			} else {
				calculation += "(" + rn.symbol[:1] + "-" + rn.symbol + ")"
			}
			n -= rn.value
		}
	}
	return result, calculation
}

func main() {
	if len(os.Args) != 2 {
		printStr("ERROR: cannot convert to roman digit\n")
		return
	}

	num, valid := parseInput(os.Args[1])
	if !valid || num == 0 || num >= 4000 {
		printStr("ERROR: cannot convert to roman digit\n")
		return
	}

	roman, calc := toRoman(num)
	if calc == "ERROR: cannot convert to roman digit" {
		printStr(calc + "\n")
		return
	}

	printStr(calc + "\n")
	printStr(roman + "\n")
}
