package main

import "fmt"

func isPositive(num int) bool {
	return num > 0
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
		output := fmt.Sprintf("isPositive(%d) \t ->  %t ", test.input, result)

		if result == test.expected {
			output += "\t (OK)"
		} else {
			output += fmt.Sprintf("\t (X) expected %t \n", test.expected)
		}

		fmt.Println(output)
	}
}
