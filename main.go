package main

import (
	"fmt"
	"sync"
	"time"
	"trygo/calc"
	"trygo/greet"
)

const Pi = 3.14

func test1(x, y string) (a, b string, c int) {
	defer fmt.Println("Print second after function returns") // mind blow
	defer fmt.Println("Print first after function returns")  // mind blow
	a = y
	b = x
	return // naked return
}

func main() {
	greet.SayHello()
	fmt.Println(calc.Sum(2, 3))

	_, name, num := test1("Toan", "Nguyen")
	fmt.Println(name, num)
	fmt.Printf("%v is type %T \n", name, name)
	fmt.Printf("%v is type %T \n", Pi, Pi)

	floatNum := float64(2) // casting
	fmt.Println(floatNum)

	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	fmt.Println()

	if i := num + 1; i == 1 {
		fmt.Println(true)
	}

	today := time.Now().Weekday()
	fmt.Println(today + 1)
	switch time.Saturday {
	case today + 1:
		fmt.Println("Tomorrow")
	default:
		fmt.Println("Far away")
	}

	// POINTER
	number := 1000
	pointer1 := &number // type *int, zero value is nil
	fmt.Printf("%v is type %T \n", pointer1, pointer1)
	fmt.Printf("%v is type %T \n", *pointer1, *pointer1)

	// STRUCT
	type User struct {
		name string
		age  uint
	}
	user1 := User{name: "Toan", age: 30}
	user2 := User{"Khanh", 29}
	pointerUser2 := &user2
	(*pointerUser2).name = "Thuy Khanh"
	fmt.Printf("%v is type %T \n", user1, user1)
	fmt.Printf("%v is type %T \n", pointerUser2.name, (*pointerUser2).name)

	// ARRAY
	var array [5]int = [5]int{1, 2, 3, 4}
	var slice []int = array[2:]
	var slice2 []int = array[1:2]
	slice[0] = 100
	// var literalSlice []string = []string{"Toan", "Khanh"} // create an array then a slice
	fmt.Printf("\n %v is type %T \n", array, array)
	fmt.Printf("%v is type %T \n", slice, slice)
	fmt.Printf("len=%d cap=%d %v\n", len(slice2), cap(slice2), slice2)
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
	slice = slice[0:2] // re-slice the slice
	fmt.Printf("len=%d cap=%d %v 🍕 \n", len(slice), cap(slice), slice)
	// dynamic array
	slice3 := make([]int, 10)
	fmt.Printf("len=%d cap=%d %v 🍕 \n", len(slice3), cap(slice3), slice3)
	slice3 = append(slice3, 1, 2, 3)
	fmt.Printf("len=%d cap=%d %v 🍕 \n", len(slice3), cap(slice3), slice3)

	for i, v := range slice3 {
		fmt.Printf("%d, %d\n", i, v)
	}

	// MAP
	var map1 = make(map[string]int)
	var literalMap map[string]string = map[string]string{
		"toan": "gia",
	}
	map1["toan"] = 1
	fmt.Printf("%v is type %T \n", map1, map1)
	fmt.Printf("%v is type %T \n", literalMap, literalMap)

	// FIRST CLASS FUNCTION has closure too
	// var myFunc func(int, int) int = func(x, y int) int { // define 2 types!!!
	//   return x + y
	// }

	// GO ROUTINE
	var say = func(s string) {
		for i := 0; i < 5; i++ {
			time.Sleep(100 * time.Millisecond)
			fmt.Println(s)
		}
	}
	go say("world")
	say("hello")

	// CHANNEL
	c1 := make(chan string)
	var dothing = func(c chan string) {
		time.Sleep(3 * time.Second)
		c <- "Done"
	}
	go dothing(c1)
	result := <-c1
	fmt.Println(result)

	// MUTEX
	var m1 sync.Mutex
	var m2 sync.Mutex
	numx := 0
	fmt.Println(m1 == m2)
	var test2 = func(x *int) {
		m1.Lock()
		m2.Lock()
		*x = *x + 1
		m2.Unlock()
	}
	fmt.Println(numx)
	for i := 0; i < 1000; i++ {
		go test2(&numx)
	}
	time.Sleep(time.Second)
	fmt.Println(numx)
}
