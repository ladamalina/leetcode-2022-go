package main

import "sort"

// FenwickTree represents the Fenwick Tree (Binary Indexed Tree) for any type T
// The fenwick tree is a data structure that calculates ranged sums and updates
// each element in O(log N).
//
// Use NewFenwickTree() for basic usage, ranged sum.  When you need more
// complicated range oeprations, use NewFenwickTreeX().
type FenwickTree[T any] struct {
	n    int
	data []T
	add  func(a, b T) T
	sub  func(a, b T) T
}

// NewFenwickTree returns a pointer to FenwickTree with basic operations for [0, n).
func NewFenwickTree[T int | int64](n int) *FenwickTree[T] {
	return NewFenwickTreeX[T](n, func(a, b T) T { return a + b }, func(a, b T) T { return a - b })
}

// NewFenwickTree returns a pointer to FenwickTree for [0, n).
// The given 'add' is a function that adds two values in the tree.
// The given 'sub' is a function that subtructs one value from another one.
func NewFenwickTreeX[T any](n int, add func(a, b T) T, sub func(a, b T) T) *FenwickTree[T] {
	ft := &FenwickTree[T]{
		n:    n,
		data: make([]T, n),
		add:  add,
		sub:  sub,
	}

	return ft
}

// Sum returns a value calculated by summing up the elements in the range [l, r)
func (ft *FenwickTree[T]) Sum(l, r int) T {
	// assert(0 <= l && l <= r && r <= ft.n)
	return ft.sub(ft.sum(r), ft.sum(l))
}

// Add adds x to position p
func (ft *FenwickTree[T]) Add(p int, x T) {
	// assert(0 <= p && p < ft.n)
	p++
	for p <= ft.n {
		ft.data[p-1] = ft.add(ft.data[p-1], x)
		p += p & -p
	}
}

func (ft *FenwickTree[T]) sum(r int) T {
	var s T
	for r > 0 {
		s = ft.add(s, ft.data[r-1])
		r -= r & -r
	}
	return s
}

func minInversionCount(nums []int, k int) int64 {
	n := len(nums)

	xComp := make(map[int]int)
	{
		unique := make([]int, n)
		copy(unique, nums)
		sort.Ints(unique)
		for _, x := range unique {
			_, exists := xComp[x]
			if !exists {
				xComp[x] = len(xComp)
			}
		}
		for i := range n {
			nums[i] = xComp[nums[i]]
		}
	}

	nUnique := len(xComp)
	ft := NewFenwickTree[int](nUnique + 5)

	l, r := 0, -1
	curInvs, minInvs := int64(0), int64(1<<60)
	for r+1 < n {
		for r-l+1 < k {
			r++
			addInvs := ft.Sum(nums[r]+1, nUnique)
			curInvs += int64(addInvs)
			ft.Add(nums[r], 1)
		}
		minInvs = min(minInvs, curInvs)
		ft.Add(nums[l], -1)
		subInvs := ft.Sum(0, nums[l])
		curInvs -= int64(subInvs)
		l++
	}
	return minInvs
}
