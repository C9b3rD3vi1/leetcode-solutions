package main

import (
	"fmt"
)

/*
 * Roman to Integer

 Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
 
 Symbol       Value
 I             1
 V             5
 X             10
 L             50
 C             100
 D             500
 M             1000
 For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
 
 Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
 
 I can be placed before V (5) and X (10) to make 4 and 9. 
 X can be placed before L (50) and C (100) to make 40 and 90. 
 C can be placed before D (500) and M (1000) to make 400 and 900.
 Given a roman numeral, convert it to an integer.
 */

func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	
	total := 0
	
	lengthOfString := len(s)
	
	for i := 0; i < lengthOfString; i++ {
		
		// If the current value is less than the next one, subtract it
		if i < lengthOfString-1 && roman[string(s[i])] < roman[string(s[i+1])] {
			total -= roman[string(s[i])]
		} else {
			total += roman[string(s[i])]
		}
	}
	
	return total
}


func main() {
	fmt.Println(romanToInt("III"))     // 3
	fmt.Println(romanToInt("IV"))      // 4
	fmt.Println(romanToInt("IX"))      // 9
	fmt.Println(romanToInt("LVIII"))   // 58
	fmt.Println(romanToInt("MCMXCIV")) // 1994
	fmt.Println(romanToInt("MMXXIII")) // 2023
	fmt.Println(romanToInt("MMXXIV"))  // 2024
	fmt.Println(romanToInt("MMXXV"))   // 2025
}
