package main

import (
	"fmt"
	"reflect"
)

func Type_Switch_Sample() {
	var t1 string = "this is a string"
	discoverType(t1)

	var t2 *string = &t1
	discoverType(t2)

	var t3 int = 123
	discoverType(t3)

	discoverType(nil)
}

func discoverType(t any) {
	switch v := t.(type) {
	case string:
		t2 := v + "..."
		fmt.Printf("string found: %s\n", t2)
	case *string:
		fmt.Printf("pointer string found: %v\n", *v)
	case int:
		fmt.Printf("we have an integer: %d\n", t)
	default:
		myType := reflect.TypeOf(t)
		if myType == nil {
			fmt.Printf("type is a nil")
		} else {
			fmt.Printf("Type not found: %s\n", reflect.TypeOf(t))
		}
	}
}
