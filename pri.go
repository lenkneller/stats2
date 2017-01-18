// Copyright (c) 2017 Leonid Kneller
//
// See the LICENSE file for an open source license information.

package stats2

func odd_is(n int) bool {
	return n&1 != 0
}

func even_is(n int) bool {
	return n&1 == 0
}
