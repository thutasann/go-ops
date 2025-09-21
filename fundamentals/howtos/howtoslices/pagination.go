package howtoslices

import "fmt"

func paginate(data []string, page, size int) []string {
	start := (page - 1) * size
	end := start + size
	if start > len(data) {
		return []string{}
	}
	if end > len(data) {
		end = len(data)
	}
	return data[start:end]
}

func Slice_Pagination_Sample() {
	fmt.Println("\n==> Slice Pagination Sample")
	users := []string{"Ann", "Bob", "Cara", "Dan", "Eve", "Finn"}
	fmt.Println(paginate(users, 1, 2)) // [Ann Bob]
	fmt.Println(paginate(users, 2, 2)) // [Cara Dan]
	fmt.Println(paginate(users, 3, 2)) // [Eve Finn]
}
