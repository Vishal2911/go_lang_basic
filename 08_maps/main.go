package main

import "fmt"

func main() {
	//map is a key value pair
	//define map
	student := make(map[string]string)

	//assigning value to map
	student["vishal"] = "7021359838"
	student["nilesh"] = "7021073478"
	student["suraj"] = "70213598914"
	fmt.Println(student)           //complete map printing
	fmt.Println(student["vishal"]) //individual key pair printing
	fmt.Println(len(student))      //length of map 3

	//delete from map
	delete(student, "suraj")
	fmt.Println(student)      //complete map printing
	fmt.Println(len(student)) //length of map 2


	//direct assign

	emails :=map[string]string{"vishal":"vishal@gmail.com", "nilesh":"nilesh@gmail.com", "suraj":"suraj@gmail.com"}
	fmt.Println(emails)
	emails["pratik"] = "pratik@gmail.com"
	fmt.Println(emails)

	// this can be used for varification
}
