package main

import "fmt"

func main(){
	fmt.Printf("function: %T \n", myFunc)
	str := 
`myFunc(num) int
----------------`
	fmt.Println(str)
}

func myFunc(num int) int {
	return 10
}