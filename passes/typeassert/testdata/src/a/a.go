package a

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string) // want `type assertion must be checked`
	fmt.Println(s)

	s, v := i.(string), "foo" // want `type assertion must be checked`
	fmt.Println(v)

	s, ok := i.(string) // ok
	fmt.Println(s, ok)

	if s := i.(string); s != "" { // want `type assertion must be checked`
		println(s)
	}

	if s, ok := i.(string); ok { // ok
		println(s)
	}

	switch n := i.(type) { // ok
	case string:
		fmt.Println(n)
	}
}
