package main

import "fmt"

/*
	goSwim indicates when it is time to go swimming.  People mostly go swiming during the Summer or when it is not cold outside.

	goSwim takes two booleans, isSummer and isCold, and it returns true when it is Summer or it is not cold

	translations:
		it is Summer -> isSummer is true
		it is not cold -> isCold is false -> !isCold is true
		or -> ||
		true when isSummer is true or !isCold is true
		true when isSummer || !isCold

	solution:
		return isSummer || !isCold
*/
func goSwim(isSummer bool, isCold bool) bool {
	return isSummer || !isCold
}

func main() {
	// goSwim : bool, bool -> bool
	tests := []test{
		{false, false, true},
		{false, true, false},
		{true, false, true},
		{true, true, true},
	}

	runTests(tests)
}

type test struct {
	isSummer bool
	isCold   bool
	expected bool
}

func runTests(tests []test) {
	fmt.Println(header())
	for _, test := range tests {
		result := goSwim(test.isSummer, test.isCold)
		fmt.Println(testString(test, result))
	}
}

func testString(t test, result bool) string {
	testStr := fmt.Sprintf(format(), t.isSummer, t.isCold, result)
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
		`goSwim(isSummer,isCold) -> bool
----------------------------------`
	return header
}

func format() string {
	name := "goSwim"
	in := "%t,%t"
	out := "%t"
	return name + "(" + in + ")\t -> " + out
}
