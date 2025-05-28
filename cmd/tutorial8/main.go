package main
import "fmt"

func main(){
	var p *int32 = new(int32)
	var i int32 = 44
	fmt.Printf("The address p points to is: %p\n", p)
	fmt.Printf("The value p points to is: %v\n", *p)
	fmt.Printf("The value of i is: %v\n", i)
	p = &i
	fmt.Printf("The value p points to is: %v\n", *p)
	*p = 30
	fmt.Printf("The value of i is: %v\n\n", i)

	thing1 := []float64{1,2,3,4}
	fmt.Printf("Memory address of thing 1 is: %p\n", &thing1)
	fmt.Printf("Memory address of thing1[0] is: %p\n", &thing1[0])
	var result1 []float64 = squareSlice(thing1)
	fmt.Printf("The result is: %v\n", result1)
	fmt.Printf("The value of thing 1 is: %v\n\n", thing1)


	arr1 := [4]float64{1,2,3,4}
	fmt.Printf("Memory address of arr1 is: %p\n", &arr1)
	fmt.Printf("Memory address of arr1[0] is: %p\n", &arr1[0])
	var result2 [4]float64 = squareArray(arr1)
	fmt.Printf("The result is: %v\n", result2)
	fmt.Printf("The value of arr1 is: %v\n\n", arr1)

	var result3 = squareArrayItself(&arr1)
	fmt.Printf("The result is: %v\n", result3)
	fmt.Printf("The value of arr1 is: %v\n", arr1)

}
func squareSlice(thing2 []float64) []float64{
	fmt.Printf("Memory address of thing2 is: %p\n", &thing2)
	for i:= range thing2{
		thing2[i] = thing2[i]*thing2[i]
	}
	return thing2
}
func squareArray(arr2 [4]float64) [4]float64{
	fmt.Printf("Memory address of arr2 is: %p\n", &arr2)
	for i:= range arr2{
		arr2[i] = arr2[i]*arr2[i]
	}
	return arr2
}

func squareArrayItself(arr3 *[4]float64) [4]float64{
	fmt.Printf("Memory address pointed by arr3 is: %p\n", arr3)
	for i:= range *arr3{
		(*arr3)[i] = (*arr3)[i] * (*arr3)[i]
	}
	return *arr3
}