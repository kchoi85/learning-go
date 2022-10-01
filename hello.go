package main

import (
	"fmt"
)

func main() {
	var studentName string = "David Choi"
	studentGradeInferred := 4.1 // short declaration operator
	var studentGradeStatic float32 = 4.1
	const PI float32 = 3.14

	// Arrays (size fixed)
	var lenInferredArr = [...]int{1, 2, 3}
	var intArr = [5]int{1, 2, 3}
	intArr[3] = 4

	// Slices (ArrayList)
	mySlice := []int{}
	mySlice = append(mySlice, 2)

	// Print statements
	fmt.Println(studentName)
	fmt.Println(studentGradeInferred)
	fmt.Println(studentGradeStatic)
	fmt.Println("Hello World from GO")

	// prints
	fmt.Printf("The const PI has value: %v and type: %T\n", PI, PI)

	// Arrays
	fmt.Println(lenInferredArr)
	fmt.Println(intArr)

	// Slices
	fmt.Println(mySlice)
	fmt.Println(cap(intArr)) // 3
	// in slices the cap is resized depending on the no. of elements in the slice

	// Practice
	fruits := [5]string{"a", "b", "c", "d", "e"}
	for i, val := range fruits {
		fmt.Printf("%v\t%v\n", i, val)
	}

	fmt.Println(add(2, 10))

	// Structs
	type Person struct {
		name  string
		age   int
		grade float32
	}

	var person1 Person
	person1.name = "david choi"
	person1.age = 24
	person1.grade = 3.9
	fmt.Println(person1)

	// Map -> maps are references! So changing content will affect the other
	// also unordered DS
	myMap := map[string]int{"status": 1, "state": 0}
	fmt.Println(myMap)
	// creating an empty map
	emptyMap := make(map[string]string)
	emptyMap["cool"] = "nice"
	// delete(emptyMap["cool"])
	fmt.Println(emptyMap)

	for k, v := range myMap {
		fmt.Printf("%v\t%v\n", k, v)
	}

}

// Functions
func add(x int, y int) int {
	return x + y
}
