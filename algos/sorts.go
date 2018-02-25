package algos

import (
	"math/rand"
)

// ConcurrentMergeSort sorts an array of ints
// turns out its slower than mergesort
func ConcurrentMergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	ch := make(chan []int)
	go concurrentMergeSortHelper(a, ch)
	return <-ch
}

func concurrentMergeSortHelper(a []int, ch chan []int) {
	if len(a) < 2 {
		ch <- a
		return
	}

	b, c := split(a)
	chc := make(chan []int)
	chb := make(chan []int)
	go concurrentMergeSortHelper(b, chb)
	go concurrentMergeSortHelper(c, chc)
	ch <- merge(<-chb, <-chc)
	return
}

// MergeSort sorts an array of ints
func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	b, c := split(a)

	b = MergeSort(b)
	c = MergeSort(c)

	return merge(b, c)
}

func merge(a []int, b []int) []int {

	aL, bL := len(a), len(b)
	c := make([]int, aL+bL)
	ai, bi := 0, 0

	for ai < aL && bi < bL {
		if a[ai] < b[bi] {
			c[ai+bi] = a[ai]
			ai++
		} else {
			c[ai+bi] = b[bi]
			bi++
		}
	}

	if ai < aL {
		c = append(c[:ai+bi], a[ai:]...)
	}

	if bi < bL {
		c = append(c[:ai+bi], b[bi:]...)
	}

	return c
}

func split(a []int) ([]int, []int) {
	aL := len(a)
	if aL < 2 {
		return a, make([]int, 0)
	}

	return a[:aL/2], a[aL/2:]
}

//GenerateRandomIntSlice generates a slice of random ints of length n
//with ints 0..n
func GenerateRandomIntSlice(n int) []int {
	a := make([]int, n)

	if n <= 0 {
		return a
	}

	for i := range a {
		a[i] = rand.Intn(n)
	}

	return a
}
