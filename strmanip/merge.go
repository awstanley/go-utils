/**
 * strmanip - String Manipulation
 * https://github.com/awstanley/go-utils/strmanip
 *
 * Copyright (C) 2016 A.W. 'Swixel' Stanley <code@swixel.net>
 *
 * This software is provided 'as-is', without any express or
 * implied warranty. In no event will the authors be held
 * liable for any damages arising from the use of this software.
 *
 * Permission is granted to anyone to use this software for any purpose,
 * including commercial applications, and to alter it and redistribute
 * it freely, subject to the following restrictions:
 *
 *   1. The origin of this software must not be misrepresented;
 *      you must not claim that you wrote the original software.
 *      If you use this software in a product, an acknowledgment
 *      in the product documentation would be appreciated but is
 *      not required.
 *   2. Altered source versions must be plainly marked as such,
 *      and must not be misrepresented as being the original
 *      software.
 *   3. This notice may not be removed or altered from
 *      any source distribution.
**/

package strmanip

// Merge takes two slices of strings (left and right) returns a merged sliced.
// This merged slice may contain duplicates and is not sorted (with right
// appended to left in a new slice).
func Merge(left, right []string) []string {
	size, i, j := len(left)+len(right), 0, 0

	// Allocate using length and capacity
	out := make([]string, size, size)

	for i = 0; i < len(left); i++ {
		out[j] = left[i]
		j++
	}

	for i = 0; i < len(right); i++ {
		out[j] = right[i]
		j++
	}
	return out
}

// MergeUnique takes two slices of strings (left and right) returns a
// merged slice with no duplicates.
//
// Uniqueness here is defined as the string matching (case sensitively).
// To define your own uniqueness test use MergeUniqueF.
// This function calls MergeUniqueF with UniqueTestDefault.
func MergeUnique(left, right []string) []string {
	return MergeUniqueF(left, right, UniqueTestDefault)
}

// MergeUniqueF uses fn (the provided FuncUniqueTest) function to merge
// the two slices (left and right) ensuring uniqueness (as defined by
// the function).
func MergeUniqueF(left, right []string, fn FuncUniqueTest) []string {
	out := Merge(left, right)
	return UniqueF(out, fn)
}

// MergeSort takes two slices of strings (left and right) returning a
// merged and sorted slice.  This slice is not checked for duplicates.
//
// To use custom sorting use MergeSortF.
// This function calls MergeSortF with SortTestDefault.
func MergeSort(left, right []string) []string {
	return MergeSortF(left, right, SortTestDefault)
}

// MergeSort takes two slices of strings (left and right) returning a
// merged and sorted slice.  Sorting is performed using the function
// provided.  This slice is not checked for duplicates.
func MergeSortF(left, right []string, fn FuncSortTest) []string {
	out := Merge(left, right)
	return QuickSort(out, fn)
}

// todo: MergeSortUnique / MergeSortUniqueF

// MergeSortUnique takes two slices (left and right) returning a unique
// array of values which are then sorted.
// This function calls MergeSortUniqueF using the default functions:
// UniqueTestDefault and SortTestDefault.
func MergeSortUnique(left, right []string) []string {
	return MergeSortUniqueF(left, right, UniqueTestDefault, SortTestDefault)
}

// MergeSortUniqueF uses provided sort and uniqueness functions to ensure
// the merged array is sorted and that the values within are unique.
func MergeSortUniqueF(left, right []string, U FuncUniqueTest, S FuncSortTest) []string {
	merged := MergeUniqueF(left, right, U)
	return QuickSort(merged, S)
}
