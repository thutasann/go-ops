package advancedpointers

import "fmt"

func Syntax() {
	fmt.Println("\n==> Syntax")
}

func Basic_One() {
	fmt.Println("\n==> Basic concept 1")

	userId := 42

	// so by just printing it, we get what is the box: the value
	println(userId)

	// if we print the address, we get where the box is located
	println(&userId)

	anotherUserID := &userId

	fmt.Println(anotherUserID)

	userId = 50

	println("after updated, another user id: ", *anotherUserID, userId)

	fmt.Println("\n==> Basic Concept 2")

	// now that we are taking about it, the * operator is also used to declare a type pointer
	var age *int

	// this will print nil, because we haven't assigned any value to it 0x0
	println("age -> ", age)

	age = &userId

	fmt.Println("age after assigned -> ", age, *age)

	// we can also create pointers using the new keyword
	age = new(int)
	*age = 100
	fmt.Println("after after new keyword -> ", *age)

	fmt.Println("\n==> Basic Concept 3")

	update(age, 42)
	fmt.Println("age after update func -> ", age, *age)
}

func update(val *int, to int) {
	*val = to
}

type PointerUser struct {
	ID   int
	Name string
}

func (u *PointerUser) updateName(name string) {
	u.Name = name
}

func Mutate_UserName() {
	fmt.Println("\n==> Mutate User name")

	u := PointerUser{ID: 42, Name: "Gopher"}
	u.updateName("James")
	fmt.Println(u)
}

func CreatePointer_Sample() {
	fmt.Println("\n==> Create Pointer Sample")
	p := create_pointer()
	fmt.Println(p)
}

func create_pointer() *int {
	x := 42
	return &x
}
