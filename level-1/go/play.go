package main

import (
	"fmt"
	"reflect"
)

func main() {
	val := reflect.ValueOf(myFunc)
	typ := reflect.TypeOf(myFunc)
	inputs := typ.NumIn()

	fmt.Println("value:", val)
	fmt.Println("type:", typ)
	fmt.Println("inputs:", inputs)
	fmt.Printf("function: %v - %v \n", val, val.Kind())
	fmt.Printf("aFunc type: %v \n", reflect.TypeOf(aFunc))
}

func myFunc(num int) int {
	return 10
}

func aFunc() {

}
