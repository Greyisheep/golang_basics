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

	// age := 35
	// name := "shaun"

	// fmt.Println("my age is", age, "and my name is", name)

	// // Printf (formatted strings)
	// fmt.Printf("my age is %v and my name is %v \n", age, name)
	// fmt.Printf("my age is %v and my name is %q \n", age, name)
	// fmt.Printf("age is of type %T \n", age)
	// fmt.Printf("you scored %0.1f points! \n", 225.55)

	// // Sprintf (save formatted strings)
	// var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	// fmt.Println("the saved string is:", str)

	// Arrays
	var ages = [3]int{20, 25, 30}

	names := [4]string{"claret", "delphine", "martins", "jessica"}
	names[1] = "gloria"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices (use arrays under the hood)
	var scores = []int{25, 75, 50}
	scores[2] = 33
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// Slice ranges

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeThree)

}
