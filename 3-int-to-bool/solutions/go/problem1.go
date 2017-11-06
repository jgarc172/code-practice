package main

import "fmt"

func isPositive(num int) bool {
	return false
}

func main() {
	// test function isPositive
	// int -> bool

	tests := []test{
		{-4, false},
		{0, false},
		{10, true},
	}

	runTests(tests)
}

type test struct {
	input    int
	expected bool
}

func runTests(tests []test) {
	fmt.Println()
	for _, test := range tests {
		result := isPositive(test.input)

		output := testString(test.input, result)
		output += okString(result, test.expected)
		fmt.Println(output)
	}
}

func testString(input int, result bool) string {
	str := "isPositive(%d) \t -> %t "
	return fmt.Sprintf(str, input, result)
}

func okString(result bool, expected bool) string {
	var okStr string
	if result == expected {
		okStr = "\t (OK)"
	} else {
		okStr = fmt.Sprintf("\t (X) expected %t \n", expected)
	}

	return okStr
}
