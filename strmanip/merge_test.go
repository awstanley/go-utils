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
	"testing"
)

// TestMerge performs testing on 'Merge'
func TestMerge(t *testing.T) {
	cases := []struct {
		test1 []string
		test2 []string
		want  []string
	}{
		// The deliberate 'fail' test
		{
			[]string{"Hello", "hello"},
			[]string{"World", "world"},
			[]string{"Hello", "hello", "World", "world"},
		},

		// Trivial 'strip duplicates' test (as this is all this is)
		{
			[]string{
				"a", "b", "c", "d", "e", "f",
			},
			[]string{
				"0", "1", "2", "3", "4",
				"5", "6", "7", "8", "9",
			},
			[]string{
				"a", "b", "c", "d", "e", "f", "0", "1",
				"2", "3", "4", "5", "6", "7", "8", "9",
			},
		},
	}
	success := false
	for id, c := range cases {
		got := Merge(c.test1, c.test2)
		if got == nil {
			t.Errorf("Merge(%d) == nil (wanted %d)", id, len(c.want))
		} else {
			success = true
			for i, _ := range c.want {
				if got[i] != c.want[i] {
					success = false
					break
				}
			}
			if !success {
				t.Errorf("Merge(%d) == did not match (length: %d; wanted: %d)", id, len(got), len(c.want))
			}
		}
	}
}

// TestMerge performs testing on 'MergeSort'
func TestMergeSort(t *testing.T) {
	cases := []struct {
		test1 []string
		test2 []string
		want  []string
	}{
		// The deliberate 'fail' test
		{
			[]string{"Hello", "hello"},
			[]string{"World", "world"},
			[]string{"Hello", "World", "hello", "world"},
		},

		// Trivial 'strip duplicates' test (as this is all this is)
		{
			[]string{
				"a", "b", "c", "d", "e",
				"f",
			},
			[]string{
				"0", "1", "2", "3", "4",
				"5", "6", "7", "8", "9",
			},
			[]string{
				"0", "1", "2", "3", "4", "5", "6", "7",
				"8", "9", "a", "b", "c", "d", "e", "f",
			},
		},
	}
	success := false
	for id, c := range cases {
		got := MergeSort(c.test1, c.test2)
		if got == nil {
			t.Errorf("MergeSort(%d) == nil (wanted %d)", id, len(c.want))
		} else {
			success = true
			for i, _ := range c.want {
				if got[i] != c.want[i] {
					success = false
					break
				}
			}
			if !success {
				t.Errorf("MergeSort(%d) == did not match (length: %d; wanted: %d)", id, len(got), len(c.want))
			}
		}
	}
}
