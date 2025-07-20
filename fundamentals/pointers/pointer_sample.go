package pointers

import "fmt"

func Pointer_Sample() {
	a := "string"
	test_pointer(&a)
	fmt.Printf("a is %s\n", a)
}

func test_pointer(a *string) {
	*a = "another string"
}

// slices are reference types — they contain a pointer to an underlying array, along with length and capacity. So, when you pass a slice to a function, you're already passing a copy of the slice header, but both the original and the copy share the same underlying array.
func Pointer_Slice_Sample() {
	a := []string{"string_one"}
	test_pointer_slice(a)
	test_pointer_slice_append(&a)
	fmt.Printf("a: %v\n", a)
}

func test_pointer_slice(a []string) {
	a[0] = "Another string"
}

func test_pointer_slice_append(a *[]string) {
	*a = append(*a, "appended string")
}

func Pointer_Map_Sample() {
	a := make(map[string]string)
	a["test"] = "value"
	test_pointer_map(a)
	fmt.Printf("a map: %v\n", a)
}

func test_pointer_map(a map[string]string) {
	a["test"] = "updated value"
	a["test_2"] = "value 2"
	a["test_3"] = "value 3"
}
