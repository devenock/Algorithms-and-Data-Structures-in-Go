package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Frequency Counter Pattern!!")

	fmt.Println(same([]int{1, 2, 3}, []int{7, 1, 9}))
	fmt.Println(sameFreq([]int{1, 2, 3}, []int{4, 1, 9}))
}

// IDEA: Use a map(or set) to collect values and their frequencies. This often avoids nested loops 0(n^2) operations on slices
// and strings and turns then into 0(n)
// Example: same
// Write a function same that accepts two slices. It returns true if every value in the first
// slice has its square in the second slice, and the frequencies of values must match.
// NAIVE SOLUTION
func same(arr1, arr2 []int) bool {
	// first check that they have the same frequency
	if len(arr1) != len(arr2) {
		return false
	}
	// make a copy of arr2 so we do not have to mutate the caller's original slice
	arr2Copy := slices.Clone(arr2)
	// loop through the items in arr1
	for _, i := range arr1 {
		arr1Sqr := i * i
		// find the index of the squared value in arr2Copy
		exist := slices.Index(arr2Copy, arr1Sqr)
		if exist == -1 {
			return false
		}
		// remove the matched element from arr2Copy so it can not be reused
		arr2Copy = slices.Delete(arr2Copy, exist, exist+1)
	}
	return true
}

// FREQUENCY COUNTER SOLUTION
func sameFreq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	// create the frequency maps
	freqA := make(map[int]int)
	freqB := make(map[int]int)

	// loop through each slice and update its frequency
	for _, val := range a {
		freqA[val]++
	}

	for _, val := range b {
		freqB[val]++
	}

	for key, val := range freqA {
		if freqB[key*key] != val {
			return false
		}
	}
	return true
}
