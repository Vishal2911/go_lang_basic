package main

import "fmt"

func main() {
	var name = "vishal"
	var vis = "aaj"
	var last = "singh"
	job := "student"
	var age = 24
	first,middle := "vishal" , "singh"
	fmt.Println("name : ", name, last, "age = ", age, job, vis)
	fmt.Printf("%T %T %T %T %T \n",name,last,age,job,vis)
	fmt.Println(first, middle)
 
}
