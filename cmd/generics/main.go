package main

import "fmt"

// generic type constraint
type Number interface {
	int64 | float64
}

func SumNumbers3[K comparable, V Number](m map[K]V) V {
	var value V

	for _, val := range m {
		value += val
	}

	return value
}

// struct generic
type userModel[N Number] struct {
	name  string
	score N
}

func (um *userModel[N]) SetName(name string) {
	um.name = name
}

func (um *userModel[N]) SetScore(score N) {
	um.score = score
}

func main() {
	var m1 userModel[float64]
	m1.SetName("fani alfirdaus")
	m1.SetScore(12.2)
	fmt.Printf("%#v\n", m1)

	m2 := userModel[int64]{name: "fanialfi", score: 99}
	fmt.Printf("%#v\n", m2)
}

// menggunakan keyword comparable
// jika menggunakan keyword comparable maka bisa support semua tipe data,
// tapi tidak untuk slice, map, dan function

// non Generics
func SumNumbers1(m map[string]int64) int64 {
	var s int64

	for _, value := range m {
		s += value
	}

	return s
}

func SumNumbers2[K comparable, V int64 | float64](m map[K]V) V {
	var s V

	for _, value := range m {
		s += value
	}

	return s
}
