package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	var a Abser
	v := Vertex{3, 4}
	//	f := Myfloat(-math.Sqrt2)
	a = &v
	//	a = f
	fmt.Println(a.Abs())
	//	fmt.Println(-math.Sqrt2)
}
