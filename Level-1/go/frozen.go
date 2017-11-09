package main

import "fmt"

func isFrozen(temp int) bool {
	return temp <= 32
}

func main() {
	// isFrozen : int -> bool
	tests := []test{
		{100, false},
		{32, true},
		{0, true},
	}

	runTests(tests)
}

type test struct {
	temp int
	expected bool
}

func runTests(tests []test) {
	fmt.Println()
	for _, test := range tests {
		result := isFrozen(test.temp)

		testStr := testString(test.temp, result)
		okStr := okString(result, test.expected)
		fmt.Println(testStr + okStr)
	}
}

func testString(input1 int, result bool) string {
	testStr := "isFrozen(%d) \t -> %t "
	return fmt.Sprintf(testStr, input1, result)
}

func okString(result bool, expected bool) string {
	if result == expected {
		return "\t (OK)"
	} else {
		return fmt.Sprintf("\t (X) expected %t", expected)
	}
}
