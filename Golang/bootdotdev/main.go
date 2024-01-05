package main

import (
	"fmt"
	"time"
)

func main() {
	const number uint = 30

	// basics
	// var answer string
	//
	// if number > 20 {
	// 	answer = fmt.Sprintf("my name is %s", "Ilesanmi oluwasijibomi")
	// }
	//
	// fmt.Println(answer)
	// fmt.Println(addTwoNumber(3, 5))
	// fmt.Println(addTwoNumber(addTwoNumber(3,3), 5))

	//basic memory allocation stuff
	num := 30
	numbe := &num

	*numbe = 31
	fmt.Println(num)
	fmt.Println(&num)
	fmt.Println(numbe)

	increment(numbe)
	fmt.Println(*numbe)

	_, y, _ := return3things(3, 4)
    print(y)

    x, _, _ := returnNaked(2,3)
    fmt.Println(x)

    var Sijibomi Person
    Sijibomi.Experirience.startDate = time.Now()

}


//NOTE: this struct is used to create a person model for our db
type Person struct {
    firstName string
    lastName string
    Experirience WorkExp
}


//NOTE: Nigerian struct
type Nigerian struct {
    Person
    tribe string
}

// a persons user experience
type WorkExp struct {
    role string
    numberOfYears int
    startDate time.Time
}

// this function extends the Person struct to return a Persons object Fullname
func (p Person) Fullname() string {
    return p.firstName + p.lastName
}

//practicing naked returns in go, this function returns height, width and int
func returnNaked (x int , y int)(height int, width int, area int){
    return height , width , area
}

// a function returning 3 things
func return3things(x, y int) (int, int, string) {
	return x, y, fmt.Sprintf("The 2 values passed are %d and %d", x, y)
}

//pass by reference in golang
func increment(x *int) {
	*x++
}

//pass a function into another function in golang
//interesting result: it computes with the values passed into the function from the function call line
func addThreeNumbers(y func(int, int) int, x int) int {
	answer := y(0, 0) + x
	return answer
}

//basic function with 2 parameters
func addTwoNumber(x, y int) int {
	return x + y
}

func print(x any) {
	fmt.Println(x)
}
