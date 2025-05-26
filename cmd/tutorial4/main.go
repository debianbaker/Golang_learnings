package main
import "fmt"

func main() {
	// Array - size + type is must on RHS o/w error. On LHS, either both(explicit) or none(inferred by Go)
	var intArr1 [5]int32 = [5]int32{1, 2, 3, 4, 5} 
	fmt.Println(len(intArr1))

	var floatArr1 = [...]float32{1.5, 2.56e+1, 20.1} 
	var newFloatArr []float32 = floatArr1[0:1]   // It is a splice not Array => no size on LHS o/w error
	fmt.Println(newFloatArr)

	boolArr := [3]bool{true, false, true}
	fmt.Println(&boolArr[0])

	var intArr2 [3]int32 = [...]int32{1,2,3}
	fmt.Println(intArr2)

	//Slice - size inferred not provided
	slice1 := []int32{1,2,3,4,5,6}
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	
	var modifiedSlice1 = append(slice1, 10) 
	fmt.Println(modifiedSlice1)
	fmt.Println(cap(modifiedSlice1))

	var modifiedSlice2 = append(slice1, []int32{1,2,3}...);
	fmt.Println(modifiedSlice2)

	slice2 := make([]int32, 3, 7)
	fmt.Println(slice2)

	slice3 := intArr1[1:]
	fmt.Println(slice3)

	// Maps:
	var myMap1 map[string]int32 = make(map[string]int32)
	fmt.Println(myMap1["hello"])

	myMap2 := make(map[int32]string)
	fmt.Println(myMap2)

	var myMap3 = map[string]int32{
		"apple": 1,
		"banana": 2,
	}
	var val, isPresent = myMap3["kela"]
	fmt.Println(val, isPresent)

	delete(myMap3, "apple")
	fmt.Println(myMap3)
}
