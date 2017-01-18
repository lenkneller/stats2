// Copyright (c) 2017 Leonid Kneller
//
// See the LICENSE file for an open source license information.

package stats2

import (
	"math"
	"sort"
)

// Summary5 computes a 5-number summary s of x[0],...,x[n-1]:
//	s[0]	minimum
//	s[1]	lower hinge
//	s[2]	median
//	s[3]	upper hinge
//	s[4]	maximum
// Precondition: n ≥ 5.
func Summary5(x []float64) (s [5]float64) {
	var n, m, k int
	var L, U float64
	n = len(x)
	if n < 5 {
		s[0] = math.NaN()
		s[1], s[2], s[3], s[4] = s[0], s[0], s[0], s[0]
		return
	}
	//
	if even_is(n) {
		m = n / 2
	} else {
		m = (n + 1) / 2
	}
	if even_is(m) {
		k = m / 2
	} else {
		k = (m + 1) / 2
	}
	//
	y := make([]float64, n)
	copy(y, x)
	sort.Float64s(y)
	//
	s[0] = y[0]
	s[4] = y[n-1]
	//
	if odd_is(n) {
		s[2] = y[m-1]
	} else {
		L, U = y[m-1], y[m]
		s[2] = L + (U-L)/2
	}
	//
	if odd_is(m) {
		s[1] = y[k-1]
		s[3] = y[n-k]
	} else {
		L, U = y[k-1], y[k]
		s[1] = L + (U-L)/2
		L, U = y[n-k-1], y[n-k]
		s[3] = L + (U-L)/2
	}
	return
}
