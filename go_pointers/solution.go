package main

import (
	"fmt"
)

type Geom struct{
	name string
}

func (g *Geom) Name() {
	fmt.Println(g.name)
}

type Vertex struct {
	Geom
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return v.X + v.Y
}

func (v *Vertex) Scale(f float64) {
	fmt.Println("hello")
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	// v0 := Vertex{1,2}
	// v0.Scale(1) // hello
	// fmt.Println(v0.Abs()) // 3

	// var v1, v2 *Vertex
	// v1 = &Vertex{1,2}
	// v1.Scale(1) // hello
	// fmt.Println(v1.Abs()) // 3

	// v2.Scale(1) // hello -> panic? 
	// fmt.Println(v2.Abs())


	v4 := Vertex{Geom{"point"}, 1, 1}
	v4.Name()
}
