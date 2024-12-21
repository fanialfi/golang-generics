package main

import "fmt"

func main() {
	total1 := Sum([]int{1, 2, 3, 4, 5})
	fmt.Println("total :", total1)
}

// notasi penulisan function dengan generic kurang lebih seperti ini
// func FunctionName[dataType <ComparableType>](params)

func Sum[V int](numbers []V) V {
	var total V

	for _, val := range numbers {
		total += val
	}

	return total
}
