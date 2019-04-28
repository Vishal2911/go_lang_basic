package main

import "fmt"

func greeting(name string) string {
	return "hello " + name
}
func getsum(num1 int, num2 int)int{
	return num1+num2
}
func main() {
	fmt.Println(greeting("vishal"))
	fmt.Println(getsum(4,5))
}
