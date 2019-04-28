package main

import "fmt"

func main(){
	//Arrays
	var fruitarr [2]string

	// assigning value
	fruitarr[0] = "mango"
	fruitarr[1] = "berry"


	fmt.Println(fruitarr)
	fmt.Println(fruitarr[0])
	// create and assign value in array in one step

	tree := [2]string{"neem", "mango"}
	fmt.Println(tree)

	// slices(we can dynamically change the size )
	student := []string{"vishal", "nilesh", "suraj"}
	fmt.Println(student)
	// count number of element in slice/arrey
	fmt.Println("length of arrey(fruitarr) = ", len(fruitarr),", length of slice = ",len(student))

	// range in slice
	fmt.Println(student[0:2])//vishal nilesh
}