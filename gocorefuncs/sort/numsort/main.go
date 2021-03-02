// ascending sort
package main

import (
	"fmt"
	"sort"
)

func main() {

	// sorts in ascending order
	// func Ints(a []int)
	a := []int{4, 9, 1, 3, 12, 6, 8, 2}
	sort.Ints(a)
	fmt.Println("a:", a)

	// func Float64s(a []float64)
	b := []float64{3.3, 9, 10.111, 2.8, 7.6}
	sort.Float64s(b)
	fmt.Println("b:", b)

	// func Strings(a []string)
	c := []string{"!", "Pear", "apple", "grapes", "pineapple", "&", "3", "orange", "~"}
	sort.Strings(c)
	fmt.Println("c:", c)

	/*
		// func IntsAreSorted(a []int) bool
		d := []int{1, 3, 4, 5, 10}
		e := []int{10, 5, 3, 4, 1}
		fmt.Println("sort.IntsAreSorted(d):", sort.IntsAreSorted(d))
		fmt.Println("sort.IntsAreSorted(e):", sort.IntsAreSorted(e))

		// func Float64sAreSorted(a []float64) bool
		f := []float64{1.1, 3, 3.37, 7.3, 9}
		g := []float64{1.1, 7.3, 9, 3, 3.37}
		fmt.Println("sort.Float64sAreSorted(f):", sort.Float64sAreSorted(f))
		fmt.Println("sort.Float64sAreSorted(g):", sort.Float64sAreSorted(g))

		// func SliceIsSorted(slice interface{}, less func(i, j int) bool) bool
		h := []string{"baseball", "basketball", "football", "hockey", "soccer"}
		i := []string{"hockey", "basketball", "soccer", "football", "baseball"}
		fmt.Println("sort.StringsAreSorted(h):", sort.StringsAreSorted(h))
		fmt.Println("sort.StringsAreSorted(i):", sort.StringsAreSorted(i))
	*/

	/*
		// func SearchInts(a []int, x int) int
		// slice must be sorted in ascending order
		j := []int{1, 2, 3, 4, 4, 7, 7, 9}
		fmt.Println("sort.SearchInts(j, 4):", sort.SearchInts(j, 4))

		// func SearchFloat64s(a []float64, x float64) int
		// slice must be sorted in ascending order
		k := []float64{1.1, 4.1, 5.99, 8.222, 63.7}
		fmt.Println("sort.SearchFloat64s(k, 5.99):", sort.SearchFloat64s(k, 5.99))

		// func SearchStrings(a []string, x string) int
		// slice must be sorted in ascending order
		l := []string{"basketball", "baseball", "football", "hockey", "soccer"}
		fmt.Println(`sort.SearchStrings(l, "soccer"):`, sort.SearchStrings(l, "soccer"))
	*/
}
