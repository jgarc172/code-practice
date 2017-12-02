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
	fmt.Printf("function: %#v - %#v - %v \n", myFunc, val, val.Kind())
	fmt.Printf("Func type: %v - %T \n", typ, aFunc)
}

func myFunc(num int) int {
	return 10
}

func aFunc() {

}
