package main

import ("fmt"
		"strconv")

// defining structure for any entity ex-student

type Student struct{ //data-type decleration (Student is name of user defined data-type)
	id int 
	name string
	class int
	gender string
	age int
}
// another method for data-type decleration
//type Student struct{
//name, gender  string
//id, class,age int
//}

//greeting method(value reciever)
func (s Student) greet()string{
	return "hello i am " +s.name +" and i am " + strconv.Itoa(s.age)
}

// age incrementer(pointer reciver)
func (s *Student) ageir(){   // it is changing value so we dont have to return any thing
	s.age++
}

// get promated in next class
func (s *Student) promotion(){
	s.class++
}

// creating function which perform both age and class increment
func (s *Student) bothInc(){
	s.promotion()
	s.ageir()
}

func main(){
	//initalize student
	s1 := Student{
		id:1,
		name:"vishal singh",
		class:10,
		gender: "male",
		age : 16,

	}
	fmt.Println(s1)
	// short-hand decleration
	s2:=Student{2,"nilesh yadav",10,"male",18}
	fmt.Println(s2)

	// to get single field
	fmt.Println(s2.name)  //nilesh yadav

	s2.age++  // or s2.age = 19
	fmt.Println(s2.age) //age = 19

	fmt.Println(s1.greet()) 


	//changing the value of age
	s1.ageir()   //age is incrimented by 1 of s1.age and
	s1.promotion()
	// class is also incremented but not printed as not required
	fmt.Println(s1.greet()) 
	// printing class 
	fmt.Println("naw class :" ,s1.class) 
	fmt.
	// calling bothInc
	s2.bothInc()
	fmt.Println(s2.greet())
	fmt.Println("naw class :" ,s2.class) 



}