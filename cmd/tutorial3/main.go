package main
import (
	"fmt"
	"errors"
)

func main(){
	const v1 string = "Hello! World"
	printMe(v1)

	var a, b int = 10, 3
	var result, remainder, err = intDivision(a, b)
	if err!=nil{
		fmt.Println(err.Error())
	} else{
	fmt.Printf("The result of the integer divison is %v with remainder is %v \n", result, remainder) 	
	}

	switch{
	case err!=nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("%v is completely divisible by %v\n", b, a)
	default:
		fmt.Printf("%v when divided by %v gives %v quotient and %v remainder\n", a, b, result, remainder)
	}

	switch remainder{
	case 0: 
		fmt.Println("Divisible by 3")
	case 1,2:
		fmt.Println("Not divisible by 3")
	}
}

func printMe(v1 string){
	fmt.Println(v1)
}

func intDivision(numerator int, denominator int) (int, int, error){
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0,0, err
	}
	return numerator/denominator, numerator%denominator, err
}