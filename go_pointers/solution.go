package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return v.X + v.Y
}

func (v *Vertex) Scale(f float64) {
	fmt.Println("hello")
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	var v1, v2 *Vertex
	v1 = &Vertex{1,2}
	v1.Scale(10)
	fmt.Println(v1.Abs())
	v2.Scale(10)
	fmt.Println(v2.Abs())
}