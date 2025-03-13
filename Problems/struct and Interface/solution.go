package main

import (
	"fmt"
	"math"
)

type Shape interface{
	area() float64
	perimeter() float64
}

type Rectangle struct{
	x1,y1,x2,y2 float64
}

type Circle struct{
	x,y,r float64
}

func distance(x1,y1,x2,y2 float64)float64{
	a:= x2-x1
	b:= y2-y1
	return math.Sqrt(a*a + b*b)
}

func(c *Rectangle) area()float64{
	l:= distance(c.x1,c.y1,c.x2,c.y2)
	w:= distance(c.x1,c.y1,c.x2,c.y2)
	return l*w
}

func (c *Circle) area()float64{
	return math.Pi * c.r * c.r
}

func (c *Circle)perimeter()float64{
	return 2 * math.Pi * c.r
}

func (c *Rectangle)perimeter()float64{
	l:= distance(c.x1,c.y1,c.x2,c.y2)
	w:= distance(c.x1,c.y1,c.x2,c.y2)
	return 2*l + 2*w
}

func totalAre(shape ...Shape)float64{
	var area float64
	for _, s := range shape{
		area += s.area()
	}
	return area
}

func totalPerimeter(shape ...Shape)float64{
	var perimeter float64
	for _, s := range shape{
		perimeter += s.perimeter()
	}
	return perimeter
}

func main(){
	var c Circle = Circle{0,0,5}
	var r Rectangle = Rectangle{0,0,10,10}
	fmt.Println("total area:",totalAre(&c,&r))
	fmt.Println("total perimeter",totalPerimeter(&c,&r))
}