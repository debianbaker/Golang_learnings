// BASIC DATA TYPES [INTRO and OPERATIONS]: 
package main
import "fmt"
import "unicode/utf8"

func main(){
	var intNum1 int8 = 127
	fmt.Println(intNum1) 

	var intNum2 uint8 = 255
	fmt.Println(intNum2) 

	var floatNum1 float32 = 24241645334.9
	fmt.Println(floatNum1) 

	var intNum3 int = 8
	var intNum4 int = 3
	var res int = intNum3/intNum4
	var result float32 = float32(intNum3)/float32(intNum4)
	fmt.Println(res)
	fmt.Println(result)

	var remainder = 107/2.5
	fmt.Println(remainder)

	var myString string= "Hello"
	fmt.Println(myString)
	fmt.Println(len(myString)) // Gives 5 - Ascii character of 1 byte, so 5 bytes = len 5
	fmt.Println(len("γ")) // Gives 2
	fmt.Println(utf8.RuneCountInString("γ")) // Gives actual Character count


	var a rune= 'A'   
	fmt.Println(a) //65

	var truth bool= false
	fmt.Println(truth)

	var var1, var2, var3 = "hi", 10.5, 20
	fmt.Println(var1, var2, var3)

	var4 := 10
	fmt.Println(var4)

	var beta, heta int
	beta = 10
	heta = 40
	beta = beta + 20
	heta = heta + 100
	fmt.Println(beta, heta)
 

	const st = "hithere" 
	// const str
	// str = "hiter" 
	fmt.Println(st)

}	