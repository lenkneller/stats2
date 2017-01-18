// Copyright (c) 2017 Leonid Kneller
//
// See the LICENSE file for an open source license information.

package stats2

// Knuth, Seminumerical Algorithms, 3rd ed, 235-236 (1998).
func knadd(u, v float64) (x, y float64) {
	var up, vpp float64
	x = u + v
	up = x - v
	vpp = x - up
	y = (u - up) + (v - vpp)
	return
}

// Knuth, Seminumerical Algorithms, 3rd ed, 616 (1998).
func knmul(u, v float64) (x, y float64) {
	const c = 134217729 // 2^27+1
	var up, u1, u2, vp, v1, v2 float64
	up, vp = u*c, v*c
	u1, v1 = (u-up)+up, (v-vp)+vp
	u2, v2 = u-u1, v-v1
	x = u * v
	y = ((((u1 * v1) - x) + (u1 * v2)) + (u2 * v1)) + (u2 * v2)
	return
}

// AccuSum computes the sum f(0)+...+f(n-1) in twice the working precision.
//
// See: Ogita, Rump, Oishi, Accurate sum and dot product, SIAM J Sci Comput, 26(6), 1955–1988.
//
// DOI: http://dx.doi.org/10.1137/030601818
func AccuSum(n int, f func(int) float64) float64 {
	var p, s, q float64
	for i := 0; i < n; i++ {
		p, q = knadd(p, f(i))
		s += q
	}
	return p + s
}

// AccuDot computes the dot product f(0)g(0)+...+f(n-1)g(n-1) in twice the working precision.
//
// See: Ogita, Rump, Oishi, Accurate sum and dot product, SIAM J Sci Comput, 26(6), 1955–1988.
//
// DOI: http://dx.doi.org/10.1137/030601818
func AccuDot(n int, f, g func(int) float64) float64 {
	var p, s, q, h, r float64
	for i := 0; i < n; i++ {
		h, r = knmul(f(i), g(i))
		p, q = knadd(p, h)
		s += q + r
	}
	return p + s
}
