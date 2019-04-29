package main

import "fmt"

func main(){
	// first method
	i:=1
for i<=10{
	fmt.Println(i)
	i++
}

//second method (short hand)
for j:=0; j<10;j++{
	fmt.Printf("j:= %d\n",j)

}

//fizz buzz problem

for i:=1;i<=100;i++{
	if i%15==0{
		fmt.Println("fizzbuzz")
		continue	
	}else if i%5==0{
		fmt.Println("buzz ")
	}else if i%3==0{
		fmt.Println("fizz")
	}else{
		fmt.Println(i)
	}

}



}