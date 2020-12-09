// Code adapted from the Tour of Go (tour.golang.org)
package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

var cLang, python, java bool
var g, h int = 1, 2

var (
	// ToBe : test bool
	ToBe bool = false
	// MaxInt : test uint64
	MaxInt uint64 = 1<<64 - 1
	// s : test complex128
	s complex128 = cmplx.Sqrt(-5 + 12i)
)

// IStillLovePython : first constant
const IStillLovePython = true

const (
	// Big : Shift 1, 100 bits to the left.
	Big = 1 << 100
	// Small : Shift Big, 99 places to the right (2).
	Small = Big >> 99
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// "short" return statements should only be used with short functions for best readability
	return
}

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {
	fmt.Println("Hello, JP!")
	fmt.Println("The random number is", rand.Intn(10))
	fmt.Println("PI is", math.Pi)
	fmt.Println("2 + 3 =", add(2, 3))

	a := "Hi"
	b := "Jackson"
	c, d := swap(a, b)

	fmt.Println(a, b, c, d)

	e, f := split(27)
	fmt.Println("27 becomes", e, "and", f)

	var i int
	fmt.Println(i, cLang, python, java)

	var j, k, l = true, false, "no!"
	fmt.Println(g, h, j, k, l)

	var m, n int = 1, 2
	o := 3
	p, q, r := true, false, "no!"

	fmt.Println(m, n, o, p, q, r)

	fmtStr := "Type: %T Value: %v\n"

	fmt.Printf(fmtStr, ToBe, ToBe)
	fmt.Printf(fmtStr, MaxInt, MaxInt)
	fmt.Printf(fmtStr, s, s)

	var t int
	var u float64
	var v bool
	var w string

	fmt.Printf("%v %v %v %q\n", t, u, v, w)

	var x, y int = 3, 4
	var z float64 = math.Sqrt(float64(x*x + y*y))
	var aa uint = uint(z)
	fmt.Println(x, y, z, aa)

	fmt.Println("Do I still love python?", IStillLovePython)

	fmt.Println(needInt(Small))
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
