package main

import (
	"fmt"
	"stepik_golang/concurency"
	ifstatement "stepik_golang/if_statement"
)

func main() {
	thischan := make(chan int, 2)
	concurency.DddDataChan(thischan, 3)
	fmt.Println(<-thischan)

	fmt.Println(ifstatement.ReturnData(5))
}
