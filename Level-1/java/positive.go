package main

import "fmt"

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
	num int
	expected bool
}

func runTests(tests []test) {
	fmt.Println()
	for _, test := range tests {
		result := isPositive(test.num)

		testStr := testString(test.num, result)
		okStr := okString(result, test.expected)
		fmt.Println(testStr + okStr)
	}
}

func testString(input1 int, result bool) string {
	testStr := "isPositive(%d)\t -> %t "
	return fmt.Sprintf(testStr, input1, result)
}

func okString(result bool, expected bool) string {
	if result == expected {
		return "\t (OK)"
	} else {
		return fmt.Sprintf("\t (X) expected %t", expected)
	}
}
