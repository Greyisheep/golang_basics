package main

import "fmt"

// func sayGreeting(n string) {
// 	fmt.Printf("Good morning %v \n", n)
// }
// func sayBye(n string) {
// 	fmt.Printf("Bye bye %v \n", n)
// }
// func cycleNames(n []string, f func(string)) {
// 	for _, v := range n {
// 		f(v)
// 	}
// }

// func circleArea(r float64) float64 {
// 	return math.Pi * r * r
// }

// func getInitials(n string) (string, string) {
// 	s := strings.ToUpper(n)
// 	names := strings.Split(s, " ")

// 	var initials []string
// 	for _, v := range names {
// 		initials = append(initials, v[:1])
// 	}

// 	if len(initials) > 1 {
// 		return initials[0], initials[1]
// 	}

// 	return initials[0], "_"
// }

func updateName(x *string) {
	*x = "wedge"
}

// func updateMenu(y map[string]float64) {
// 	y["coffee"] = 2.99
// }

func main() {

	// group A types -> strings, ints, bools, floats, arrays, structs
	name := "ozii"

	// updateName(name)

	// fmt.Println("memory address of name is: ", &name)

	m := &name
	fmt.Println("memory addresss: ", m)
	fmt.Println("value at memory address", *m)
	fmt.Println(name)

	updateName(m)
	fmt.Println(name)

	// group B types -> slices, maps, functions
	// menu := map[string]float64{
	// 	"coffee":    5.95,
	// 	"ice cream": 3.99,
	// }

	// updateMenu(menu)
	// fmt.Println(menu)

	// menu := map[string]float64{
	// 	"soup":   4.99,
	// 	"pie":    7.99,
	// 	"salad":  6.99,
	// 	"toffee": 3.55,
	// }

	// fmt.Println(menu)
	// fmt.Println(menu["pie"])

	// // looping maps
	// for k, v := range menu {
	// 	fmt.Println(k, "-", v)
	// }

	// // ints as key type
	// phonebook := map[int]string{
	// 	8110774138: "grey",
	// 	8038379343: "dota",
	// 	8034894915: "kc",
	// }

	// fmt.Println(phonebook)
	// fmt.Println(phonebook[8110774138])

	// phonebook[8038379343] = "chinyere"
	// fmt.Println(phonebook)
	// sayHello("claret")

	// for _, v := range points {
	// 	fmt.Println(v)
	// }

	// fn1, sn1 := getInitials("annie leonhart")
	// fmt.Println(fn1, sn1)

	// fn2, sn2 := getInitials("eren jaeger")
	// fmt.Println(fn2, sn2)

	// fn3, sn3 := getInitials("titans")
	// fmt.Println(fn3, sn3)

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
	// var ages = [3]int{20, 25, 30}

	// names := [4]string{"claret", "delphine", "martins", "jessica"}
	// names[1] = "gloria"

	// fmt.Println(ages, len(ages))
	// fmt.Println(names, len(names))

	// // Slices (use arrays under the hood)
	// var scores = []int{25, 75, 50}
	// scores[2] = 33
	// scores = append(scores, 85)

	// fmt.Println(scores, len(scores))

	// // Slice ranges

	// rangeOne := names[1:3]
	// rangeTwo := names[2:]
	// rangeThree := names[:3]

	// fmt.Println(rangeOne)
	// fmt.Println(rangeTwo)
	// fmt.Println(rangeThree)

	// greeting := "hello there friends!"

	// fmt.Println(strings.Index(greeting, "ll"))
	// fmt.Println(strings.Split(greeting, " "))

	// fmt.Println(strings.Contains(greeting, "hello"))
	// fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// fmt.Println(strings.ToUpper(greeting))

	// ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	// sort.Ints(ages)
	// fmt.Println(ages)

	// index := sort.SearchInts(ages, 30)
	// fmt.Println(index)

	// names := []string{"claret", "delphine", "martins", "jessica"}
	// sort.Strings(names)
	// fmt.Println(names)

	// fmt.Println(sort.SearchStrings(names, "claret"))

	// x := 0
	// for x < 5 {
	// 	fmt.Println("value of x is:", x)
	// 	x++
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is:", i)
	// }

	// names := []string{"claret", "delphine", "martins", "jessica"}

	// // for i := 0; i < len(names); i++ {
	// // 	fmt.Println(names[i])
	// // }

	// // for index, value := range names {
	// // 	fmt.Printf("the value at index %v is %v. \n", index, value)
	// // }

	// for _, value := range names {
	// 	fmt.Printf("the value is %v. \n", value)
	// 	value = "new string"
	// }

	// fmt.Println(names)

	// age := 45

	// if age < 30 {
	// 	fmt.Println("age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else {
	// 	fmt.Println("age is not less than 45")
	// }

	// names := []string{"claret", "delphine", "martins", "jessica"}

	// for index, value := range names {
	// 	if index == 1 {
	// 		fmt.Println("continuing at pos", index)
	// 		continue
	// 	}
	// 	if index > 2 {
	// 		fmt.Println("breaking at pos", index)
	// 		break
	// 	}

	// 	fmt.Printf("the value at pos %v is %v \n", index, value)

	// sayGreeting("claret")
	// sayGreeting("martins")
	// sayBye("claret")

	// cycleNames([]string{"gloria", "chinwendu", "kaduru"}, sayGreeting)
	// cycleNames([]string{"gloria", "chinwendu", "kaduru"}, sayBye)

	// a1 := circleArea(10.5)
	// a2 := circleArea(15)

	// fmt.Println(a1, a2)
	// fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f", a1, a2)

}
