package main

import "fmt"

func or(val1 bool, val2 bool) bool {
	return val1 || val2
}

func main() {
	// test function or
	// bool, bool -> bool

	tests := []test{
		{false, false, false},
		{false, true, true},
		{true, false, true},
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
		result := or(test.val1, test.val2)

		fmt.Println(testString(test, result))
	}
}

func testString(t test, result bool) string {
	format := "or(%t, %t)  \t -> %t "
	testStr := fmt.Sprintf(format, t.val1, t.val2, result)
	okStr := okString(result, t.expected)
	return testStr + okStr
}

func okString(result bool, expected bool) string {
	if result == expected {
		return "\t (OK)"
	}
	return fmt.Sprintf("\t (X) expected %t", expected)
}
