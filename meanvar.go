// Copyright (c) 2017 Leonid Kneller
//
// See the LICENSE file for an open source license information.

package stats2

import (
	"math"
)

// MeanVar computes the sample mean (m) and the sample variance (s2) of x[0],...,x[n-1], n = len(x):
//	m = (Σx[i])/n,
//	s2 = (Σ(x[i]-m)²)/(n-1),
//	i = 0,...,n-1.
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
