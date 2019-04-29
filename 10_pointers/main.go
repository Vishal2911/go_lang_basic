package main

import "fmt"

func main(){
	a:=4
	b:=&a // store add of a 
	c:=b // store value of b which is addr of a
	
	fmt.Println(a,&a,b,c,*b,*c) //&a print addr of a ,*b and *c print value at addr of a =4

	//change value with pointer
	*c=5
	fmt.Println(a)  //a=5
}