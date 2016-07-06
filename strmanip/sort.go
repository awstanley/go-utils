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

import (
	"math/rand"
)

// QuickSort performs a quicksort on slice 'a' using sort function 'fn'.
// The returned value is sorted based on the sorting function passed.
//
// This is a pretty typical quicksort, if you're unfamiliar see one of the
// vast sources out there (e.g. https://en.wikipedia.org/wiki/Quicksort).
func QuickSort(a []string, fn FuncSortTest) []string {
	// If it's less than 2 it can't be sorted
	if len(a) < 2 {
		return a
	}

	// Get working bounds
	left, right := 0, len(a)-1

	// Get a random pivot (no need to use crypto/rand,
	// we don't want security or true random, we want speed)
	pvt := rand.Intn(len(a) - 1)

	// Move the pivot to the right
	a[pvt], a[right] = a[right], a[pvt]

	// Move smaller elements to the left, using the function.
	for i := range a {
		if fn(a[i], a[right]) {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Recursion hell.
	QuickSort(a[:left], fn)
	QuickSort(a[left+1:], fn)

	// Done
	return a
}
