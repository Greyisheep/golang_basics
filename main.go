package main

import "fmt"

func main() {
	// // strings
	// var nameOne string = "claret"
	// var nameTwo = "martins"
	// var nameThree string

	// // fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "peach"
	// nameThree = "bowser"

	// // fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "Jessica"

	// fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// ints
	// var ageOne int = 20
	// var ageTwo int = 30
	// ageThree := 40

	// fmt.Println(ageOne, ageTwo, ageThree)

	// // bits & memory
	// // var numOne int8 = 25
	// // var numTwo int8 = -128
	// // var numThree uint8 = 255

	// var scoreOne float32 = 25.98
	// var scoreTwo float64 = 343435543236.7655
	// scoreThree := 1.5

	age := 35
	name := "shaun"

	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted strings)
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %v and my name is %q \n", age, name)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %0.1f points! \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
}
