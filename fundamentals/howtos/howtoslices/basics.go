package howtoslices

import (
	"fmt"
)

func Slice_Anatomy() {
	fmt.Println("\n==> Slice Anatomy")
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(len(s), cap(s))
}

func Basic_Slicing() {
	fmt.Println("\n==> Basic Slicing")
	nums := []int{10, 20, 3} //https://github.com/thutasann0, 40, 50}
	a := nums[1:4]           // from index 1 up to 4 (not including 4)
	fmt.Println(a)
}

func Slices_Share_The_Same_Array() {
	fmt.Println("\n==> Slices Share the Same Array")
	nums := []int{1, 2, 3, 4, 5}
	a := nums[1:4] // [2,3,4]
	a[0] = 99
	fmt.Println("a ==> ", a)
	fmt.Println("original ===> ", nums) // [1, 99, 3, 4, 5]
}

// If you slice x[i:j],
//
// len = j - i
//
// cap = cap(x) - i
//
// Length: You are asking for elements from index i up to (but not including) index j.
//
// 👉 That gives exactly j - i elements.
//
// Capacity: Capacity is measured from the new starting index i all the way to the end of the original backing array, because that’s how far appends can go without reallocating.
//
// 👉 That’s why it’s cap(x) - i.
//
// nums := []int{1,2,3,4,5}  // len=5, cap=5
//
// a := nums[1:3]            // slice [2,3]
//
// len(a) = j - i = 3 - 1 = 2
//
// cap(a) = cap(nums) - i = 5 - 1 = 4 → can reach [2, 3, 4, 5]
func Capacity_Controls_Growth() {
	fmt.Println("\n==> Capacity Controls Growth")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("original capcity ==> ", cap(nums))

	a := nums[1:3] // len=2, cap=4 ==> cap(nums) - 1 ==> 5 - 1 ==> cap becomes 4
	fmt.Println("after slice capacity ==> ", len(a), cap(a))

	a = append(a, 99)
	fmt.Println("after append, original capcity ==> ", nums, cap(nums))
}

func Reslicing() {
	fmt.Println("\n==> Reslicing")
	nums := []int{1, 2, 3, 4, 5}
	a := nums[1:3] // [2, 4], len=2, cap=4
	b := a[:4]     // extend to full cap
	fmt.Println(b)
}

func Copying_Slices() {
	fmt.Println("\n==> Copying Slices")

	nums := []int{1, 2, 3, 4, 5}
	a := nums[1:3] // [2,3]
	b := make([]int, len(a))
	copy(b, a)
	a[0] = 99

	fmt.Println("a -> ", a)       // [99 3]
	fmt.Println("nums -> ", nums) // [1 99 3 4 5]
	fmt.Println("b -> ", b)       // [2,3]
}

func Splitting_CSV_Rows() {
	fmt.Println("\n==> Splitting CSV Rows")
	lines := "john,doe,30,developer"
	parts := []rune(lines)
	fmt.Println(string(parts[0:4]))
}
