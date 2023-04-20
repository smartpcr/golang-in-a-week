package main

import "fmt"

type Comparable interface {
	GreaterThan(to Comparable) bool
}

type Int int

var _ Comparable = Int(0)

func (i Int) GreaterThan(to Comparable) bool {
	return i > to.(Int)
}

type String string

var _ Comparable = String("")

func (i String) GreaterThan(to Comparable) bool {
	return i > to.(String)
}

type Float64 float64

var _ Comparable = Float64(0)

func (i Float64) GreaterThan(to Comparable) bool {
	return i > to.(Float64)
}

func Max[T Comparable](values []T) T {
	if len(values) == 0 {
		panic("empty slice")
	}
	max := values[0]
	for _, value := range values {
		if value.GreaterThan(max) {
			max = value
		}
	}
	return max
}

func main() {
	ints := []Int{1, 2, 3}
	fmt.Println(Max(ints))

	floats := []Float64{1.1, 2.2, 3.3}
	fmt.Println(Max(floats))

	strings := []String{"foo", "bar", "baz"}
	fmt.Println(Max(strings))
}
