package a

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	s := fmt.Sprint(rand.Int())  // want `don't use fmt.Sprint to convert number to string. Use strconv.Itoa.`
	s = strconv.Itoa(rand.Int()) // OK
	fmt.Println(s)
	s = fmt.Sprint(true)        // OK
	s = fmt.Sprint(uint(10))    // want `don't use fmt.Sprint to convert number to string. Use strconv.Itoa.`
	s = fmt.Sprint(float64(10)) // want `don't use fmt.Sprint to convert number to string. Use strconv.Itoa.`
	v := 100
	s = fmt.Sprintln(v)            // want `don't use fmt.Sprint to convert number to string. Use strconv.Itoa.`
	s = fmt.Sprintln(123, 456)     // want `don't use fmt.Sprint to convert number to string. Use strconv.Itoa.`
	s = fmt.Sprintln(123, "hello") // ok
}
