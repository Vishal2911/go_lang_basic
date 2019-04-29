package main

import ("fmt"
		"math")

type Shape interface{
	area() float64
}

type Circle struct{
	x,y,radius float64
}
type Rectangle struct{
	width, height float64
}
func(c Circle)area()float64{
	return math.Pi*c.radius*c.radius
}
func(r Rectangle)area()float64{
	return r.width* r.height
}

func getArea(s Shape) float64{
	return s.area()
}

func main(){
	circle :=Circle{x:0, y:0, radius:5}//method of assigning
	rectangle:= Rectangle{10,5}//another mesthod of assigning
	fmt.Printf("circle area :=%f \n",getArea(circle))

	fmt.Printf("rectangle area :=%f \n",getArea(rectangle))

}