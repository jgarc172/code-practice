package main

import "fmt"

/*
	isPositive indicates if a number is positive based on the definition of a positive number: a number is positive if it is greater than 0.

	isPositive takes an integer, num, and it returns true if num is positive

		isPositive : integer -> boolean
			Ex:
		isPositive(100)	 -> true 	 (OK)
		isPositive(0)	 -> false 	 (OK)
		isPositive(-5)	 -> false 	 (OK)

	translations:
		- num is positive if it is greater than 0
		- true when num > 0
		  false otherwise (when num <= 0)
	solution:
		return num > 0
*/

func isPositive(num int) bool {
	return num > 0
}

func main() {
	// isPositive : int -> bool
	tests := []test{
		{100, true},
		{0, false},
		{-5, false},
	}

	runTests(tests)
}

type test struct {
	num      int
	expected bool
}

func runTests(tests []test) {
	fmt.Println(header())
	for _, test := range tests {
		result := isPositive(test.num)
		fmt.Println(testString(test, result))
	}
}

func testString(t test, result bool) string {
	testStr := fmt.Sprintf(format(), t.num, result)
	okStr := okString(result, t.expected)
	return testStr + okStr
}

func okString(result bool, expected bool) string {
	if result == expected {
		return "\t (OK)"
	}
	return fmt.Sprintf("\t (X) expected %t", expected)
}

func header() string {
	header :=
		`isPositive(num) -> bool
-----------------------`
	return header
}

func format() string {
	name := "isPositive"
	in := "%d"
	out := "%t"
	return name + "(" + in + ")\t -> " + out
}
