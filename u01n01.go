// Copyright (c) 2017 Leonid Kneller
//
// See the LICENSE file for an open source license information.

package stats2

import (
	"crypto/rand"
	"math"
	"math/big"
)

var two53 *big.Int = big.NewInt(1 << 53)

// U01 returns a pseudo-random number uniformly distributed on (0,1).
// U01 uses a cryptographically strong pseudo-random number generator.
// U01 is safe for concurrent use by multiple goroutines.
//
// U01 returns a pseudo-random number u satisfying the following:
//	ε ≤ u ≤ 1-ε, where ε = 2^-53 = 1.1102230246251565404236316680908203125₁₀-16.
func U01() float64 {
	for {
		b, err := rand.Int(rand.Reader, two53)
		if err != nil {
			panic(err)
		}
		n := b.Int64()
		if n != 0 {
			return float64(n) / (1 << 53)
		}
	}
}

// N01 returns a pseudo-random number normally distributed with mean 0 and variance 1.
// N01 uses a cryptographically strong pseudo-random number generator.
// N01 is safe for concurrent use by multiple goroutines.
//
// See: Knuth, Seminumerical Algorithms, 3rd ed, 130-131 (1998).
func N01() float64 {
	const C1 = 1.7155277699214135929603792825575449562415972155051 // sqrt(8/e)
	const C2 = 5.1361016667509659362936822722497458333451234611259 // 4*sqrt(sqrt(e))
	for {
		u, v := U01(), U01()
		x := C1 * (v - 0.5) / u
		x2 := x * x
		if x2 <= 5-C2*u {
			return x
		}
		if x2 <= -4*math.Log(u) {
			return x
		}
	}
}
