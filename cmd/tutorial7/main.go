package main
import "fmt"

// Custom named types - just naming the Go types
type intNum int32                       
type tracker map[string]int32
type engine1 struct{
	mpg int16
	gallons int16
}
type engine2 struct{
	mpg int16
	gallons int16
	ownerInfo owner        // custom type owner: is a struct
}
type owner struct{
	name string
	age int16
}

func (e engine1) milesLeft() int16{               
	return e.gallons*e.mpg
}
func milesLeft(e engine1) int16{
	return e.gallons*e.mpg
}
// Above two functions look the same - accept only engine1 type, return the same.
// Key diff is - 1st one gives OOP behaviour in go - Method on struct engine1
// Understand Like- engine1 is class, myEngine1 is object and 1st fn. is fn of the class which object calls.


func (e engine2) milesLeft() int16{               
	return e.gallons*e.mpg
}
type engine interface{
	milesLeft() int16
}

func main(){
	var intNum1 intNum = 44
	fmt.Println(intNum1)

	var myTracker tracker = tracker{
		"banana": 2,
	}
	myTracker["apple"] = 1
	fmt.Println(myTracker)

	var myEngine1 = engine1{ 10,20 }
	fmt.Println(myEngine1.mpg, myEngine1.gallons)

	var myEngine2 = engine2{
		15, 20,
		owner{
			name: "Puneet",
			age: 21,
		},
	}
	fmt.Println(myEngine2, myEngine2.ownerInfo.age)

	var localStructEg = struct{               //Type on right side in must
		name string
		age int16
	}{
		"Puneet",
		21,
	}
	fmt.Println(localStructEg.name, localStructEg.age)

	fmt.Println("Miles left for engine 1:", milesLeft(myEngine1))       
	fmt.Println("Miles left for engine 1:", myEngine1.milesLeft())

	canMakeIt(myEngine1, 250)
	canMakeIt(myEngine2, 250)

}	
func canMakeIt(e engine, miles int16){
	if(miles <= e.milesLeft()){
		fmt.Println("You can make it there")
	}else{
		fmt.Println("You cannot make it there")
	}
}