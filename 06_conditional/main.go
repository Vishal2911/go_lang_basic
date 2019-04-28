package main

import "fmt"

func main(){
	x:=10
	z:=7
	y:=6
	if (x>y && x>z){
	fmt.Printf("%d is the gretest among %d %d %d \n",x,x,y,z)
	} // 10 is the gretest among 10,7,6
}