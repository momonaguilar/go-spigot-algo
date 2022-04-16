package spigot

import (
	"math/big"
	"strings"
)

var (
	Precision = 1
)

type lft struct {
	q, r, s, t big.Int
}

func (t *lft) extr(x *big.Int) *big.Rat {
	var n, d big.Int
	var r big.Rat
	return r.SetFrac(
		n.Add(n.Mul(&t.q, x), &t.r),
		d.Add(d.Mul(&t.s, x), &t.t))
}

var three = big.NewInt(3)
var four = big.NewInt(4)

func (t *lft) next() *big.Int {
	r := t.extr(three)
	var f big.Int
	return f.Div(r.Num(), r.Denom())
}

func (t *lft) safe(n *big.Int) bool {
	r := t.extr(four)
	var f big.Int
	return n.Cmp(f.Div(r.Num(), r.Denom())) == 0
}

func (t *lft) comp(u *lft) *lft {
	var r lft
	var a, b big.Int
	r.q.Add(a.Mul(&t.q, &u.q), b.Mul(&t.r, &u.s))
	r.r.Add(a.Mul(&t.q, &u.r), b.Mul(&t.r, &u.t))
	r.s.Add(a.Mul(&t.s, &u.q), b.Mul(&t.t, &u.s))
	r.t.Add(a.Mul(&t.s, &u.r), b.Mul(&t.t, &u.t))
	return &r
}

func (t *lft) prod(n *big.Int) *lft {
	var r lft
	r.q.SetInt64(10)
	r.r.Mul(r.r.SetInt64(-10), n)
	r.t.SetInt64(1)
	return r.comp(t)
}

// GetPi returns the Pi value without decimal point (e.g. 31415xxx)
// containing Y digits representing precision.
func GetPi() string {
	// init z to unit
	z := new(lft)
	z.q.SetInt64(1)
	z.t.SetInt64(1)

	// lfts generator
	var k int64
	lfts := func() *lft {
		k++
		r := new(lft)
		r.q.SetInt64(k)
		r.r.SetInt64(4*k + 2)
		r.t.SetInt64(2*k + 1)
		return r
	}

	pi := []string{}
	for {
		y := z.next()
		if z.safe(y) {
			pi = append(pi, y.String())
			z = z.prod(y)
		} else {
			z = z.comp(lfts())
		}

		if len(pi) > Precision {
			break
		}
	}

	result := strings.Join(pi, "")
	return result
}
