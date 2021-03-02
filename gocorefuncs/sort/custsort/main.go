package main

import (
	"fmt"
	"sort"
)

/*
type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less reports whether the element with
    // index i should sort before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}
*/

type mySort []int

// Len is the number of elements in the collection.
func (s mySort) Len() int {
	return len(s)
}

// Less reports whether the element with index i should sort before the element with index j.
func (s mySort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Swap swaps the elements with indexes i and j.
func (s mySort) Less(i, j int) bool {
	return s[i] < s[j]
	// return s[i] > s[j]
}

func main() {

	var nums mySort = []int{4, 2, 6, 8, 1, 7, 3, 9, 5}
	// func Sort(data Interface)
	sort.Sort(nums)
	fmt.Println(nums)

	/*
		nums2 := []int{3, 8, 1, 6, 4, 5, 2, 7, 9}
		sort.Sort(mySort(nums2))
		fmt.Printf("nums2  Type:%T  Value:%v\n", nums2, nums2)
		fmt.Printf("mySort(nums2)  Type:%T  Value:%v\n", mySort(nums2), mySort(nums2))
		fmt.Printf("nums2  Type:%T  Value:%v\n", nums2, nums2)
	*/
}
