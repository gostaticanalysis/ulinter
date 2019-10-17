package b

import (
	fmt2 "fmt"
	"math/rand"
	"strconv"
)

func main() {
	s := fmt2.Sprint(rand.Int()) // want `don't use fmt.Sprint to convert number to string. Use strconv.Itoa.`
	s = strconv.Itoa(rand.Int()) // OK
	fmt2.Println(s)
	s = fmt2.Sprint(true)        // OK
	s = fmt2.Sprint(uint(10))    // want `don't use fmt.Sprint to convert number to string. Use strconv.Itoa.`
	s = fmt2.Sprint(float64(10)) // want `don't use fmt.Sprint to convert number to string. Use strconv.Itoa.`
	v := 100
	s = fmt2.Sprintln(v)            // want `don't use fmt.Sprint to convert number to string. Use strconv.Itoa.`
	s = fmt2.Sprintln(123, 456)     // want `don't use fmt.Sprint to convert number to string. Use strconv.Itoa.`
	s = fmt2.Sprintln(123, "hello") // ok
}
