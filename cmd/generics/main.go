package main

import "fmt"

func main() {
	ints := map[string]int64{"first": 34, "second": 12}
	floats := map[string]float64{"first": 35.98, "second": 26.99}

	fmt.Printf("Generic Sums with Constraint : %v and %v\n",
		SumNumbers2(ints),
		SumNumbers2(floats),
	)

	intJuga := map[int]int64{1: 34, 2: 12}
	fmt.Printf("Generic Sums with Constraint comparable : %d\n", SumNumbers2(intJuga))
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
