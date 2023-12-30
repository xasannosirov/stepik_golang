package main

import "fmt"

func sumInt(numbers ...int) (int, int) {
	var sum, count int
	for _, elem := range numbers {
		sum += elem
		count += 1
	}
	return count, sum
}

func main(){
	fmt.Println(sumInt(1,2,3,4))
}