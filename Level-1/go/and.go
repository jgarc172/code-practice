package main

import "fmt"

func and(val1 bool, val2 bool) bool {
	return val1 && val2
}

func main() {
	// test function and
	// bool, bool -> bool

	tests := []test{
		{false, false, false},
		{false, true, false},
		{true, false, false},
		{true, true, true},
	}

	runTests(tests)
}

type test struct {
	val1     bool
	val2     bool
	expected bool
}

func runTests(tests []test) {
	fmt.Println()
	for _, test := range tests {
		result := and(test.val1, test.val2)

		testStr := testString(test.val1, test.val2, result)
		testStr += okString(result, test.expected)
		fmt.Println(testStr)
	}
}

func testString(input1 bool, input2 bool, result bool) string {
	testStr := "and(%t, %t) \t -> %t "
	return fmt.Sprintf(testStr, input1, input2, result)
}

func okString(result bool, expected bool) string {
	var okStr string
	if result == expected {
		okStr = "\t (OK)"
	} else {
		okStr = fmt.Sprintf("\t (X) expected %t", expected)
	}

	return okStr
}
