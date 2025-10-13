package howtoslices

import "fmt"

type FilterUser struct {
	Name   string
	Active bool
}

func FilterActive(users []FilterUser) []FilterUser {
	active := users[:0] // reuse underlying array
	for _, u := range users {
		if u.Active {
			active = append(active, u)
		}
	}
	return active
}

func Filter_Transform_Sample() {
	fmt.Println("\n===> Filter Transform Sample")
	users := []FilterUser{
		{"Alice", true},
		{"Bob", false},
		{"Charlie", true},
	}
	active := FilterActive(users)
	fmt.Println(active)
}
