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
	"strings"
)

// FuncUniqueTest defines the uniqueness test function:
// If a and b are equal then the return shall be true;
// in all other causes the return shall be false.
type FuncUniqueTest func(a string, b string) bool

// UniqueTestDefault performs a trivial a == b test.
func UniqueTestDefault(a string, b string) bool {
	return a == b
}

// UniqueTestCaseInsensitive performs a trivial a == b test
// after lowering both strings.
func UniqueTestCaseInsensitive(a string, b string) bool {
	A := strings.ToLower(a)
	B := strings.ToLower(b)
	return A == B
}

// FuncSortTest defines the sort test function:
// If a should go before b the value shall be true;
// in all other cases the value shall be false.
//
// This is a limited sort function designed to work with
// the included QuickSort.
type FuncSortTest func(a string, b string) bool

// SortTestDefault performs a trivial a < b test.
func SortTestDefault(a string, b string) bool {
	if a < b {
		return true
	}

	return false
}

// SortTestCaseInsensitive performs a trivial a < b test
// after lowering both strings.
func SortTestCaseInsensitive(a string, b string) bool {
	A := strings.ToLower(a)
	B := strings.ToLower(b)

	if A < B {
		return true
	}
	return false
}
