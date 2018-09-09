package main

import "fmt"

// v.(type)
// v.(int)
func main() {
	var interfaces []interface{}
	interfaces = append(interfaces, 1)
	interfaces = append(interfaces, "str")
	interfaces = append(interfaces, true)
	for _, v := range interfaces {
		switch v.(type) {
		case string:
			fmt.Printf("%T\n", v)
		case bool:
			fmt.Printf("%T\n", v)
		case int:
			fmt.Printf("%T\n", v)
		}
		if _, ok := v.(int); ok {
			fmt.Printf("%T\n", v)
		}
		if _, ok := v.(bool); ok {
			fmt.Printf("%T\n", v)
		}
		if _, ok := v.(string); ok {
			fmt.Printf("%T\n", v)
		}
	}

}

