// Copyright (c) 2017 Leonid Kneller
//
// See the LICENSE file for an open source license information.

package stats2

import (
	"math"
)

// MeanVar computes the sample mean (m) and the sample variance (s2) of x[0],...,x[n-1]:
//	m = (x[0]+...+x[n-1])/n,
//	s2 = ((x[0]-m)²+...+(x[n-1]-m)²)/(n-1).
// Precondition: n ≥ 2.
func MeanVar(x []float64) (m, s2 float64) {
	n := len(x)
	if n < 2 {
		m, s2 = math.NaN(), math.NaN()
		return
	}
	m = AccuSum(n, func(i int) float64 { return x[i] }) / float64(n)
	s2 = AccuDot(n, func(i int) float64 { return x[i] - m }, func(i int) float64 { return x[i] - m }) / float64(n-1)
	return
}
