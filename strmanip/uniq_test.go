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

// TestUnique performs testing on the default Unique function.
func TestUnique(t *testing.T) {
	cases := []struct {
		test []string
		want []string
	}{
		// The deliberate 'fail' test
		{
			[]string{"Hello", "hello", "World", "world"},
			[]string{"Hello", "hello", "World", "world"},
		},

		// Trivial 'strip duplicates' test (as this is all this is)
		{
			[]string{
				"Hello", "hello", "World", "world",
				"Hello", "hello", "World", "world",
				"Hello", "hello", "World", "world",
			},
			[]string{
				"Hello", "hello", "World", "world",
			},
		},
	}
	hitCount := 0
	for id, c := range cases {
		got := Unique(c.test)
		if got == nil {
			t.Errorf("Unique(%d) == nil, want %d", id, len(c.want))
		} else {
			hitCount = 0
			for _, b := range got {
				for _, d := range c.want {
					if b == d {
						hitCount++
						break
					}
				}
			}
			if hitCount != len(c.want) {
				t.Errorf("Unique(%d) == %d, want %d", id, hitCount, len(c.want))
			}
		}
	}
}
