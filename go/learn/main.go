package main

import (
	"fmt"
)


func main() {
    closure()
    //cool()
 //    list :=  LinkedList{}
 //    list.append(40)
 //    list.append(40)
 //    list.append(40)
 //    list.append(40)
 //    list.append(40)
 //    list.display()
	//
	//
	// var fleet = []string{"siji", "sore", "seyi"}
	// var Coordinates = map[string]int{}
	//
	// var fleetSlice []string = fleet[:]
	// fleetSlice[0] = "tola"
	//
	// var goodFleetSlice = []string{}
 //    //var goodFleetSlice = make([]string, 20)
	// for i := 0; i < 15; i++ {
	// 	goodFleetSlice = append(goodFleetSlice, fmt.Sprintf("%d", i))
	// }
	// fmt.Printf("Good fleet size: %v\n", goodFleetSlice)
	//
	// var Coords map[string]int
	// Coords = make(map[string]int)
 //    
 //    var matrix = make([][]int , 10)
	//
 //    for i:= range matrix{
 //        matrix[i] = make([]int, 10)
 //    }
 //    
 //    matrix[0][1] = 40
	//
	// Coords["me"] = 40
	// Coords["you"] = 30
	//
	// fmt.Printf("Coords: %v\n", Coords)
	//
	// Coordinates["x"] = 30
	// Coordinates["y"] = 30
	//
	// fleet = append(fleet, "ay")
	// fmt.Printf("Coordinates: %v\n", Coordinates)
	//
	// fmt.Printf("fleet: %v\n", fleet)
	// fmt.Printf("fleetSlice: %v\n", fleetSlice)
	//
	// length := len(fleet)
	// fmt.Println(length)
	//
	// for i := 0; i < length; i++ {
	// 	fmt.Printf("fleet[i]: %v\n", fleet[i])
	// }
	//
	// charPrintf("sijibomi")
}

func charPrintf(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}
}

// drivable interface, makes the vehicles drive
type Drive interface {
	drive()
}

// Car struct with inline wheel struct
type Car struct {
	make  string
	model string
	wheel struct {
		number int
		size   int
	}
}

type Truct struct {
	Car
	space int
}

func (c Truct) drive() {
	fmt.Println("The Truck is driving")
}
func (c Car) drive() {
	fmt.Println("The car is driving")
}

func nowDrive(vehicle Drive) {
	vehicle.drive()
}

type userError struct {
	name string
}

type insult interface {
	Yab() string
}

func (e userError) Yab() string {
	return fmt.Sprintf("This guy %v is a good guy", e.name)
}

func (e userError) r() string {
	return fmt.Sprintf("This guy %v is a nigga", e.name)
}

func sendSms(msg string, userName string) string {
	return userError{name: userName}.Yab()
}

func getMessage() []string {
	return []string{
		"Hello",
		"Hello",
		"Hello",
	}
}
