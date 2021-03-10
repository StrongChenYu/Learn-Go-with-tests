package main

import (
	"fmt"
	"strings"
)

type RomanNumeral struct {
	Value int
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) getValueBySymbol(symbols ...byte) int {
	symbol := string(symbols)
	for _, item := range allRomanNumerals {
		if symbol == item.Symbol {
			return item.Value
		}
	}
	return 0
}

var allRomanNumerals = RomanNumerals {
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

func ConvertToRoman(arabic int) string {

	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}

	return result.String()
}


func ConvertToArabic(roman string) int {
	res := 0
	for i := 0; i < len(roman); i++ {
		symbol := roman[i]
		if couldBeSubtractive(i, symbol, roman) {
			if value := allRomanNumerals.getValueBySymbol(symbol, roman[i + 1]); value != 0 {
				res += value
				//这个i++会移动两次
				i++
			} else {
				//这种情况是IL这种，这个是不允许减的，所以只需要加一个简单的值就可以了
				res += allRomanNumerals.getValueBySymbol(symbol)
			}
		} else {
			res += allRomanNumerals.getValueBySymbol(symbol)
		}
	}
	return res
}

func couldBeSubtractive(i int, symbol uint8, roman string) bool {
	isSubtractiveSymbol := symbol == 'I' || symbol == 'X' || symbol == 'C'
	return i + 1 < len(roman) && isSubtractiveSymbol
}

func main() {
	if value := 1; value != 0 {
		fmt.Println("chenyu")
	}
}
