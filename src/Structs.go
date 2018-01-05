package main

import "fmt"
import "math"

func learnStruct() {
	fmt.Println("Hello struct")
}

type Circle struct {
	x, y, r float64
}

func main() {
	d := Circle{0, 0, 5}
	fmt.Println(d.x, d.y, d.r)
	fmt.Println("Value of d==> ", d)
	fmt.Println("Circle Area===> ", circleArea(&d))
	fmt.Println("Value of d==> ", d)
	learnStruct()

	p := Person{"Muhamad Nizar Iqbal"}
	p.talk()

	a := new(Android)
	a.Person.talk()
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Person struct {
	Name string
}

type Android struct {
	Person Person
	Model  string
}

func circleArea(c *Circle) float64 {
	*c = Circle{0, 0, 10}
	return math.Pi * c.r * c.r
}
