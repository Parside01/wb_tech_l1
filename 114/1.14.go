package main

import (
	"fmt"
)

func TypeIdentification(data interface{}) string {
	switch data.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	case chan bool:
		return "chan bool"
	}
	return "unknown"
}

func main() {
	fmt.Println(TypeIdentification(true))
}
