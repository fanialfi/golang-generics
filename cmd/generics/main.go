package main

import "fmt"

// disebut type constraint
type number interface {
	int | float32 | float64
}

func main() {
	total1 := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println("total :", total1)

	total2 := Sum([]float32{1.1, 2.2, 3.3, 4.4, 5.5})
	fmt.Println("total :", total2)

	total3 := Sum([]float64{11.11, 22.22, 33.33, 44.44, 55.55})
	fmt.Println("total :", total3)

	sumInt := Sum[int]
	total4 := sumInt([]int{1, 2, 3, 4, 4, 4, 4, 4, 4, 4})
	fmt.Println("total :", total4)
}

// notasi penulisan function dengan generic kurang lebih seperti ini
// func FunctionName[dataType <ComparableType>](params)

func Sum[V number](numbers []V) V {
	var total V

	for _, val := range numbers {
		total += val
	}

	return total
}
