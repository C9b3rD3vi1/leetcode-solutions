package main
// Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.
import "fmt"


func addDigits(num int) int {
	
	if num == 0 && num < 10 {
		return 0
	}
	return 1 +((num-1) % 9)
}


func main() {
	fmt.Println(addDigits(38))  // Output: 2
	fmt.Println(addDigits(0))   // Output: 0
	fmt.Println(addDigits(123)) // Output: 6
	fmt.Println(addDigits(999)) // Output: 9
	fmt.Println(addDigits(12345)) // Output: 6
}
