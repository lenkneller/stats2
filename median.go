// Copyright (c) 2017 Leonid Kneller
//
// See the LICENSE file for an open source license information.

package stats2

import (
	"math"
	"sort"
)

// Medians computes the low median (lomed) and the high median (himed) of x[0],...,x[n-1]:
// The statistical median of x[0],...,x[n-1] is defined as following:
//	n is odd:  median = lomed = himed,
//	n is even: median = (lomed+himed)/2.
// Precondition: n ≥ 1.
func Medians(x []float64) (lomed, himed float64) {
	n := len(x)
	if n < 1 {
		lomed, himed = math.NaN(), math.NaN()
		return
	}
	//
	var m int
	if even_is(n) {
		m = n / 2
	} else {
		m = (n + 1) / 2
	}
	//
	y := make([]float64, n)
	copy(y, x)
	sort.Float64s(y)
	//
	if odd_is(n) {
		lomed, himed = y[m-1], y[m-1]
	} else {
		lomed, himed = y[m-1], y[m]
	}
	return
}

// MedianMAD computes the statistical median (med) and the median absolute deviation (MAD) of x[0],...,x[n-1].
// The median absolute deviation is defined as following:
//	med = median{x[i]},
//	MAD = median{|med-x[i]|}, i=0,...,n-1.
// Precondition: n ≥ 1.
func MedianMAD(x []float64) (med, MAD float64) {
	n := len(x)
	if n < 1 {
		med, MAD = math.NaN(), math.NaN()
		return
	}
	//
	var m int
	if even_is(n) {
		m = n / 2
	} else {
		m = (n + 1) / 2
	}
	//
	y := make([]float64, n)
	copy(y, x)
	//
	sort.Float64s(y)
	if odd_is(n) {
		med = y[m-1]
	} else {
		lo, hi := y[m-1], y[m]
		med = lo + (hi-lo)/2
	}
	//
	for i, yi := range y {
		y[i] = math.Abs(med - yi)
	}
	//
	sort.Float64s(y)
	if odd_is(n) {
		MAD = y[m-1]
	} else {
		lo, hi := y[m-1], y[m]
		MAD = lo + (hi-lo)/2
	}
	//
	return
}

// FactorMAD is the scale factor to make MAD a consistent estimator for the standard deviation estimation of a normal distribution.
const FactorMAD = 1.4826022185056018605470765293604234313267032025903
