package main

import (
	"fmt"
	"time"
)

func main(){
	myMap := map[string]int32{
		"apple": 1,
		"banana": 2,
	}
	// Map iteration
	for name, num := range myMap{
		fmt.Printf("Name: %v, Value: %v \n", name, num)
	}
	// Array iteration
	intArr := [5]int32{1,2,3,4,5}
	for i, v := range intArr{
		fmt.Printf("Value at index %v is %v \n" ,i, v)
	}
	// Traditional Loop like java
	for i:=1; i<=5; i++{
		fmt.Println(i)
	}

	// While loop
	i := 5
	for i != 0{
		fmt.Println(i)
		i--
	}

	for {
		if(i != 5){
			i++
		}else {
			fmt.Println("Exiting")
			break
		}
	}

	// Small Question-
	// Append 1 to a slice n=1000000 times - without preallocation and with preallocation('make') and check diff in times.

	var n int = 1000000
	slice1 := []int{}
	slice2 := make([]int, 0, n)
	fmt.Printf("Total time without preallocation is: %v\n", timeCalc(slice1, n))
	fmt.Printf("Total time with preallocation is: %v\n", timeCalc(slice2, n))
}
func timeCalc(slice []int, n int) time.Duration{
		var t0 = time.Now()
		for len(slice) < n{
			slice = append(slice, 1)
		}
		return time.Since(t0);
	}