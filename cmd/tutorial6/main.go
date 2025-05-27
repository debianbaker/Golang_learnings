package main

import (
	"fmt"
	"strings"
)

func main(){
	var myString = "résumé"
	fmt.Println(myString[1])        // gives 1st byte only
	for i, v := range myString{
		fmt.Println(i, v)           //skips index
	}
	var intArr = []rune(myString)  // OR int32(alias)
	fmt.Println(intArr[1])

	for i, v := range intArr{
		fmt.Println(i, v)
	}

	var a rune = 'a'  // 
	fmt.Println(a)

	strSlice := []string{"s", "a", "l", "é"}
	var catSlice = ""
	for i := range strSlice{
		catSlice = catSlice + strSlice[i]
	}
	fmt.Println(catSlice)

	var strBuilder strings.Builder
	for i:= range strSlice{
		strBuilder.WriteString(strSlice[i])	
	}
	var catStr = strBuilder.String()
	fmt.Println(catStr)

	var bytes1 byte = 'i'   // OR uint8 i.e 256 characters
	fmt.Println(bytes1)

	var byteArr = []byte("é")   // String is utf8 so 2 bytes here
	fmt.Println(byteArr[1])
}