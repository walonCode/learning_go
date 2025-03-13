package main

import (
	"fmt"
	"math"
)

//working with struct

//define a struct
type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1,y1,x2,y2 float64
}

//defining a complex struct
type Person struct{
	name string;
}
func (p *Person) speak(){
	fmt.Println("My name is",p.name)
}

type Andriod struct{
	Person Person;
	Model string
}

//working with interface
//It check if the struct has an area() method
// but to know that we must use the memery locator(&)
type Shape interface{
	area() float64
}

func main(){
	fmt.Println("Struct and interface")

	

	//initializing the struct
	var c Circle = Circle{0,0,5}

	//another way of initializing a struct using the new key word
	cirle := new (Circle)
	cirle.x = 5
	cirle.y = 7
	cirle.r = 10



	//using the struct
	fmt.Println(c,c.r)
	fmt.Println(cirle)

	//working with the struct
	fmt.Println(c.area())

	//creating a reactangle
	var rect Rectangle = Rectangle{0,0,10,10}
	fmt.Println("rectangle area: ",rect.area())

	//defining a new andriod
	var phone Andriod = Andriod{
		Person: Person{name: "walon"},
		Model: "Samsung",
	}

	phone.Person.speak()
	fmt.Println(phone.Model)

	fmt.Println("total area: ",totalArea(&rect, &c))

}


//this is way of adding method to struct
func (c *Circle) area()float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area()float64 {
	l:= distance(r.x1,r.y1,r.x2,r.y2)
	w:= distance(r.x1,r.y1,r.x2,r.y2)
	return l*w
}

func distance(x1,y1,x2,y2 float64)float64{
	a:= x2 - x1
	b:= y2 - y1	
	return math.Sqrt(a*a + b*b)
}

func totalArea(shapes ...Shape)float64{
	var area float64
	for _, s:= range shapes {
		area += s.area()
	}
	return area
}
	