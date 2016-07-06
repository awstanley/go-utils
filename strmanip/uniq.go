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

// Unique uses the default unique function (UniqueTestDefault) to remove
// duplicates from the slice.
func Unique(a []string) []string {
	return UniqueF(a, UniqueTestDefault)
}

// UniqueF uses function (fn) to test for uniqueness in removing any
// duplicates.
func UniqueF(a []string, fn FuncUniqueTest) []string {
	size, i, j, k := len(a), 0, 0, 0

	b := make([]string, size, size)

	for i = 0; i < size; i++ {
		// Only scan against newer items
		for j = i + 1; j < size; j++ {
			if fn(a[i], a[j]) {
				break
			}
		}
		if j == size {
			b[k] = a[i]
			k++
		}
	}
	return b[0:k]
}
