package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"reflect"
	"strconv"
)

var I int = 1

const (
	x1 = iota
	x2 = iota
	x3 = iota
)

type Employee struct {
	id             int
	employeeName   string
	departmentName string
}

func main() {

	var i int
	i = 42
	//i = 30

	j := 50
	fmt.Println("This is the first time i am running the code.")
	fmt.Println(i + j)
	fmt.Println(i+j == 90)

	fmt.Printf(" %v, %T \n", I, i)

	var theUrl string = "https://www.google.com"

	fmt.Println(theUrl, " any")
	fmt.Println(theUrl, " 2any")

	var convertedString string
	convertedString = strconv.Itoa(j)

	fmt.Printf("converted string is %v and type is %T \n", convertedString, convertedString)

	// agenda Boolean Values

	n := 1

	var result bool

	result = n == 1

	fmt.Println(result)

	// numeric Values int8 --> -128 to 128
	// int 16, int32, int64
	// uint16, uint32

	a := 10
	b := 12
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b) //  not producing float values
	fmt.Println(a * b)
	fmt.Println(b % a)

	a = 0
	b = 1
	//  Binary Operators
	fmt.Println("Binary Operations")
	fmt.Println(a | b)
	fmt.Println(a & b)
	fmt.Println(a ^ b)
	fmt.Println(b &^ a)

	// Shift Operations

	a = 8

	fmt.Println(a << 1)
	fmt.Println(a >> 1)

	n1 := 3.14
	n1 = 13.7e72
	n1 = 2.1e9
	fmt.Printf("%v %T \n", n1/2, n1)

	// complex numbers

	// var complex complex64 = 2i
	var complex complex64 = complex(2, 1222)
	fmt.Printf("%v %v %T \n", real(complex), imag(complex), complex)

	// Strings of text -- strings are immuatable

	s := "Hello I am Learning Go Programming Language"
	s1 := " <<<This is new String>>>"
	fmt.Printf(" %v , %T \n", s+s1, s)

	// Byte Slices of strings
	s3 := []byte(s)
	fmt.Printf(" %v , %T \n", s3, s3)

	// characters  Rune is integer is 32

	c := 'a'

	fmt.Printf("%v, %T \n", c, c)

	// Constants
	// 1. Naming Convention
	// 2. Typed Constatns
	// 3. UnTyped Constants
	// 4. Enumerated Constants
	// 5. Enumeration Expressions

	// const keyword is used to define the contant.
	// If first letter is capital then the const will be exported.
	// const is like a final keyword in java.
	// once assigned, cant be reassigned again.

	const myConst int = 42
	var integerA int = 32
	fmt.Printf("%v %T", myConst, myConst)

	r1 := myConst + integerA
	fmt.Println(r1)

	// Enumerated Constants

	fmt.Printf("%v %T\n", x1, x1)
	fmt.Printf("%v %T\n", x2, x2)
	fmt.Printf("%v %T\n", x3, x3)

	const (
		x4 = iota
		x5 = iota
		x6
	)

	fmt.Println(x4)
	fmt.Println(x6) // answer 2, values are automatically assigned

	const (
		_ = iota + 5
		catSpecialist
		dogSpecialist
		pigSpecialist
	)

	const cat = 6
	const dog = 7

	fmt.Println(catSpecialist == cat)
	fmt.Println(dogSpecialist != dog)

	// Arrays
	// 1.Creation
	// 2.Built In Functions
	// 3.Working with Arrays
	// 4. make operation is used to create an array

	grades := [3]int{12, 13, 14}

	fmt.Println(grades[1])

	fmt.Printf(" %v %T\n", grades[1], grades)

	var students [3]string

	students[0] = "student1"
	students[1] = "student2"
	students[2] = "student3"

	fmt.Println("length of the array", students, len(students))
	fmt.Println("capacity of the array", students, cap(students))

	// 3 x 3 matrix

	var twoDimentionalMatrix [3][3]string

	twoDimentionalMatrix[0] = [3]string{"firstRow1", "firstRow2", "firstRow3"}
	twoDimentionalMatrix[1] = [3]string{"secoundRow1", "secoundRow2", "secoundRow3"}
	twoDimentionalMatrix[2] = [3]string{"thirdRow1", "thirdRow2", "thirdRow3"}

	fmt.Printf("%v, %T", twoDimentionalMatrix, twoDimentionalMatrix)

	// copying arrays and new reference is created

	array1 := [...]int{1, 2, 3}
	array2 := array1
	array2[1] = 5
	fmt.Println("\n", array1, array2)

	array3 := &array1

	array3[1] = 100001

	fmt.Println(array1, array3)

	make1 := make([]int, 10)
	make2 := make([]int, 10, 100) // capacoty = 100
	fmt.Println("Using make Funcion creating the array", make1)
	fmt.Printf("Using make %v, and Capacity = %v", len(make2), cap(make2))

	// Slices
	// 1.Creation
	// 2.Built In Functions
	// 3.Working with Slices

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	slice2 := slice1

	fmt.Println("before Changing the value", slice1, slice2)

	slice2[1] = 33333

	fmt.Println("After changing the value", slice1, slice2)

	slice3 := slice1[:]
	slice4 := slice1[:3]
	slice5 := slice1[3:]
	slice6 := slice1[3:7]
	fmt.Println("slice all elements", slice3)
	fmt.Println("slice upto index 3", slice4)
	fmt.Println("slice from index 3 to end", slice5)
	fmt.Println("slice from index 3 to 6", slice6)

	// NOTE :: Slicing operations also works with arrays

	array4 := array1[1:2]
	fmt.Println(array4)

	// when we append the new elements to the array, then the capacity of the array increases
	// if length crosses the capacity, then the capaciuty will be incremented by 2 times
	// Example :: If length is 7 and capacity is 8. New element(9) is added.
	// 9 ---> add operation --> len = 8, cap = 8
	// 10 ---> add operation --> len = 9, cap = 8 * 2 = 16
	// upto length 16, it accepts all add opertioans. When length crosses the 16 the capacity is multiplied by 2
	// 16 --> len = 17, cap = 32

	cap1 := []int{}
	fmt.Println(cap1, len(cap1), cap(cap1))
	cap1 = append(cap1, 2, 3, 4)
	fmt.Println(cap1, len(cap1), cap(cap1))
	cap1 = append(cap1, 2, 3, 4, 5)
	fmt.Println(cap1, len(cap1), cap(cap1))
	cap1 = append(cap1, 2, 3, 4, 5)
	fmt.Println(cap1, len(cap1), cap(cap1))
	cap1 = append(cap1, 2, 3, 4, 5, 7, 8)
	fmt.Println(cap1, len(cap1), cap(cap1))
	cap2 := []int{1, 2, 3}
	cap2 = append(cap2, []int{4, 5, 6}...)
	fmt.Println("append operation output using three dots", cap2, len(cap2))

	// Remove operations in slices
	deleteOp := []int{2, 3, 4, 5, 6}
	fmt.Println(deleteOp)
	deleteOp1 := append(deleteOp[:1], deleteOp[2:]...)
	fmt.Println("delete the secound element", deleteOp1)

	// Maps
	// what are they ?
	// Creating
	// Manipulation

	statePupulations := map[string]int{
		"california":      12345,
		"Texas":           27345,
		"Florida":         12345,
		"Andhara Pradesh": 234555,
	}

	fmt.Println(statePupulations)

	m := map[string][]int{
		"numbers":      []int{1, 2, 3, 4, 5},
		"even_numbers": []int{0, 2, 4, 6, 8},
	}
	m["odd_numbers"] = []int{1, 3, 5, 7, 9}
	m["whole_numbers"] = []int{0, 1, 2, 3, 4, 5}

	fmt.Println(m)

	val, ok := m["whole_numbers"]
	fmt.Println(val, ok)

	// delete the elements in the map

	delete(m, "whole_numbers")

	fmt.Println("After delete the key", m)

	// Structs
	// What are they ?
	// Creating
	// Naming Conventions
	// Embedding
	// Tags

	aEmployee := Employee{
		id:             3,
		employeeName:   "Subhash", // using variables names is suggested
		departmentName: "SOFTWARE",
	}

	fmt.Println(aEmployee, aEmployee.employeeName)

	// short lived time

	aEmployeeDifferent := struct{ name string }{name: "manikanta"}
	aEmployeeDifferent1 := &aEmployeeDifferent
	aEmployeeDifferent1.name = "manikanta_1"
	fmt.Println(aEmployeeDifferent.name)
	fmt.Println(aEmployeeDifferent1.name)

	// Embedding -- No Inheritance creates a relation
	// Composition -- But use composition to create a relation between the structs

	type Animal struct {
		name   string
		origin string
	}

	type Bird struct {
		Animal
		speedKph float32
		canFly   bool
	}

	bird := Bird{
		Animal:   Animal{name: "parrot", origin: "india"},
		speedKph: 48,
		canFly:   true,
	}

	fmt.Println(bird, "\n", bird.canFly)

	// reflect packages to be used to get the type

	type Vehicle struct {
		Name   string `required max:"100"`
		Origin string
	}

	vehicle_name_type := reflect.TypeOf(Vehicle{})
	field, _ := vehicle_name_type.FieldByName("Name")
	fmt.Println(field, "\n", field.Tag)

	// if statements
	// Operators
	// if else and if else if statements

	if true {
		fmt.Println("If statemtn is true")
	} // always curly braces should be there

	if ok {
		fmt.Println("MAP IS TRUE")
	}

	_, telanganaExists := m["Telangana"]
	if !telanganaExists {
		fmt.Println(" Telangana not exists")
	}

	n11 := 50
	n2 := 70

	var result1 string
	if n11 > n2 {
		fmt.Println(" N1 is greateer than n2")

	}
	if n11 < n2 {
		fmt.Println(" n2 is greater than n2")
		result1 = "n2 is grerater than n1"
	}
	if n11 == n2 {
		fmt.Println("n1 is equal to n2")
	}

	if n11 <= n2 {
		//fmt.Println("N1 is less than or equal to n2")
		result1 = "N1 is less than or equal to n2"
	}

	if n11 < n2 || n11 == n2 {
		//fmt.Println("N1 is less than or equal to n2")
		result1 = "N1 is less than or equal to n2"
	}

	if false {
		fmt.Println("Never will be executed")
	} else {
		fmt.Println("Always else block will be executed")
	}

	fmt.Println(result1, returnTrue())

	n111 := math.Pow(2, 10)

	fmt.Println(n111)

	// Switch statements
	// simple Cases
	// cases with multiple tests
	// Falling through
	// Type switches

	switchStatements()
	loopingOperations()
	loopingOperationsForMaps(statePupulations)

	// Defer -- executes after the function completes
	// open and close functions are using defer keyword
	// so that resources are open --
	deferExamples() // stack manner it executes
	// Panic
	//runTheWebServer()
	// if panic and defer statements are present , then defer statements will be
	// executed then the panic statements are executed.
	// if panic is not handled then the program will exit
	panicAndDeferStatementsAtSameTime()
	// Recover also covered in above function
	// Only useful in defer functions,
	//

}

func returnTrue() bool {
	fmt.Println("function returnTrue() got called")
	return true
}

func switchStatements() {

	//Observe: No breaks are implemented
	n1 := 1 + 5
	switch n1 {
	case 1:
		fmt.Println("1 is prinitted")
	case 2, 4, 6, 8, 10:
		fmt.Println("2,4,6,8,10 is printed")
	default:
		fmt.Println("default is printed")
	}
	// Tag Syntaxes

	switch {
	case n1 < 10:
		fmt.Println(" n1 is less than 10")
	case n1 < 20:
		fmt.Println(" n1 is less thane 20")
	}
	// fallthrough --> still executes the next case irrespective of true or false.

	// type switches

	var n1111 interface{} = "subhash"

	switch n1111.(type) {
	case int:
		fmt.Println("n1 is of type integer")
	case float64:
		fmt.Println("n1 is type of float64")
	default:
		fmt.Println(" n1 is none of the above types")
	}
}

func loopingOperations() {
	fmt.Println(" Printing the for loop for 3 times")
	for i := 0; i < 3; i++ {
		fmt.Println("Print i for 3 times", i)
	}

	printEvenNumbersUpto10()
}

func printEvenNumbersUpto10() {
	var x int = 10
	fmt.Print("even numbers upto 10 ")
	for i := 0; i < x; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	// break keywords
	i := 0
	for ; i < 5; i++ {
		if i == 4 {
			fmt.Println("\nBreak Usage", i)
			break
		}
		if i < 4 {
			fmt.Println(" Continue usage", i)
			continue
		}
	}

	// Print 5 * 5 stars(*)

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// range keywords
	s := []int{1, 2, 3}

	for index, val := range s {
		fmt.Println(index, val)
	}
}

func loopingOperationsForMaps(statePop map[string]int) {
	for ke, val := range statePop {
		fmt.Printf(" Key : %v, value : %v \n", ke, val)
	}
	// we want only values instead of ke, replace with _
	var tempVariable1 int
	for _, val := range statePop {
		tempVariable1 = val
	}
	fmt.Println(tempVariable1)
}

func deferExamples() {
	defer fmt.Println("first")
	fmt.Println("secound")
	defer fmt.Println("third")

	deferExampleInRealTime()
}

func deferExampleInRealTime() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", robots)
}

func runTheWebServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func panicAndDeferStatementsAtSameTime() {

	fmt.Println("Panic And Defer Example")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("Something Bad Happen")
}
