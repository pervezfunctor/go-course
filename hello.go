package main

import (
	"fmt"
	"math"
)

func Swap[T any](x, y *T) {
	*x, *y = *y, *x
}

type Less func(x, y interface{}) bool

func Max[T any](x, y T, less Less) T {
	if less(x, y) {
		return x
	}
	return y
}

type Lesser[T any] interface {
	Less(that T) bool
}

func MaxV[T Lesser[T]](x, y T) T {
	if x.Less(y) {
		return y
	}
	return x
}

type Ints interface {
	~int8 | ~int16 | ~int32 | ~int64
}

type Foo int8

func IndexOf[T Ints](s []T, e T) int {
	for i, a := range s {
		if a == e {
			return i
		}
	}
	return -1
}

type Point struct {
	X int
	Y int
}

func (pt Point) Distance() float64 {
	return math.Sqrt(float64(pt.X*pt.X + pt.Y*pt.Y))
}

type Size struct {
	Width  int
	Height int
}

type Rect struct {
	Point
	Size
}

func (r Rect) RectArea() int {
	return r.Width * r.Height
}

type Point3D struct {
	Point
	Z int
}

func Distance(pt Point3D) float64 {
	return math.Sqrt(float64(pt.X*pt.X + pt.Y*pt.Y + pt.Z*pt.Z))
}

func PointLess()
func main() {
}

type Equaler[T any] interface {
	Equal(T) bool
}

type Adder[T any] interface {
	Add(T) T
}

type Suber[T any] interface {
	Sub(T) T
}

type Multer[T any] interface {
	Mul(T) T
}

type Diver[T any] interface {
	Div(T) T
}

type Numeric[T any] interface {
	Equaler[T]
	Adder[T]
	Suber[T]
	Multer[T]
	Diver[T]
	fmt.Stringer
	int | float64
	Foo()
}

func Arith[T Numeric[T]](x, y T) (T, T, T, T) {
	return (x.Add(y)), (x.Sub(y)), (x.Mul(y)), (x.Div(y))
}
