package tdist

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mathext"
)

// A TDist is a Student's t-distribution with V degrees of freedom.
type TDist struct {
	V float64
}

func (t TDist) GInv(p float64) float64 {
	if p < 0 || p > 1 {
		panic("Error")
	}
	// F(x) = 1 - 0.5 * I_t(x)(nu/2, 1/2)
	// t(x) = nu/(t^2 + nu)
	var y float64
	if p > 0.5 {
		// Know t > 0
		g := mathext.InvRegIncBeta(t.V/2, 0.5, 2*(1-p))
		y = math.Sqrt(t.V * (1 - g) / g)
	} else {
		g := mathext.InvRegIncBeta(t.V/2, 0.5, 2*p)
		y = -math.Sqrt(t.V * (1 - g) / g)
	}
	// Convert out of standard normal
	return y
}
func (t TDist) Inv(p float64) float64 {
	// var e = a.ibetainv(2 * b.min(c, 1 - c), .5 * d, .5);
	// return e = b.sqrt(d * (1 - e) / e),
	// c > .5 ? e : -e
	if p > 0.5 {
		// Know t > 0
		e := IBetaInv(2*(1-p), 0.5*t.V, 0.5)
		return math.Sqrt(t.V * (1 - e) / e)
	} else {
		e := IBetaInv(2*p, 0.5*t.V, 0.5)
		return -math.Sqrt(t.V * (1 - e) / e)
	}
	// e := IBetaInv(2*math.Min(p, 1-p), 0.5*t.V, 0.5)
	//
	// a := math.Sqrt(t.V * (1 - e) / e)
	// if p > 0.5 {
	// 	return a
	// } else {
	// 	return -a
	// }

}

func GIBeta(a, b, x float64) float64 {
	return mathext.RegIncBeta(a, b, x)
}

func GIBetaInv(a, b, y float64) float64 {
	return mathext.InvRegIncBeta(a, b, y)
}

// Returns the inverse of the incomplete beta function
//
//		jStat.ibetainv = function ibetainv(p, a, b) {
//		  var EPS = 1e-8;
//		  var a1 = a - 1;
//		  var b1 = b - 1;
//		  var j = 0;
//		  var lna, lnb, pp, t, u, err, x, al, h, w, afac;
//		  if (p <= 0)
//		    return 0;
//		  if (p >= 1)
//		    return 1;
//		  if (a >= 1 && b >= 1) {
//		    pp = (p < 0.5) ? p : 1 - p;
//		    t = Math.sqrt(-2 * Math.log(pp));
//		    x = (2.30753 + t * 0.27061) / (1 + t* (0.99229 + t * 0.04481)) - t;
//		    if (p < 0.5)
//		      x = -x;
//		    al = (x * x - 3) / 6;
//		    h = 2 / (1 / (2 * a - 1)  + 1 / (2 * b - 1));
//		    w = (x * Math.sqrt(al + h) / h) - (1 / (2 * b - 1) - 1 / (2 * a - 1)) *
//		        (al + 5 / 6 - 2 / (3 * h));
//		    x = a / (a + b * Math.exp(2 * w));
//		  } else {
//		    lna = Math.log(a / (a + b));
//		    lnb = Math.log(b / (a + b));
//		    t = Math.exp(a * lna) / a;
//		    u = Math.exp(b * lnb) / b;
//		    w = t + u;
//		    if (p < t / w)
//		      x = Math.pow(a * w * p, 1 / a);
//		    else
//		      x = 1 - Math.pow(b * w * (1 - p), 1 / b);
//		  }
//		  afac = -jStat.gammaln(a) - jStat.gammaln(b) + jStat.gammaln(a + b);
//		  for(; j < 10; j++) {
//		    if (x === 0 || x === 1)
//		      return x;
//		    err = jStat.ibeta(x, a, b) - p;
//		    t = Math.exp(a1 * Math.log(x) + b1 * Math.log(1 - x) + afac);
//		    u = err / t;
//	     TODO: check this rule in js (if else and equal)
//		    x -= (t = u / (1 - 0.5 * Math.min(1, u * (a1 / x - b1 / (1 - x)))));
//		    if (x <= 0)
//		      x = 0.5 * (x + t);
//		    if (x >= 1)
//		      x = 0.5 * (x + t + 1);
//		    if (Math.abs(t) < EPS * x && j > 0)
//		      break;
//		  }
//		  return x;
//		}
func IBetaInv(p, a, b float64) float64 {
	EPS := 1e-8
	a1 := a - 1
	b1 := b - 1
	var lna, lnb, pp, t, u, err, x, al, h, w, afac float64

	if p <= 0 {
		return 0
	}
	if p >= 1 {
		return 1
	}
	if a >= 1 && b >= 1 {
		pp = 1 - p
		if p < 0.5 {
			pp = p
		}
		t = math.Sqrt(-2 * math.Log(pp))
		x = (2.30753+t*0.27061)/(1+t*(0.99229+t*0.04481)) - t

		if p < 0.5 {
			x = -x
		}
		al = (x*x - 3) / 6
		h = 2 / (1/(2*a-1) + 1/(2*b-1))
		w = (x * math.Sqrt(al+h) / h) - (1/(2*b-1)-1/(2*a-1))*
			(al+5/6-2/(3*h))

		x = a / (a + b*math.Exp(2*w))
	} else {
		lna = math.Log(a / (a + b))
		lnb = math.Log(b / (a + b))
		t = math.Exp(a*lna) / a
		u = math.Exp(b*lnb) / b
		w = t + u
		if p < t/w {
			x = math.Pow(a*w*p, 1/a)
		} else {
			x = 1 - math.Pow(b*w*(1-p), 1/b)
		}
	}
	afac = -Gammaln(a) - Gammaln(b) + Gammaln(a+b)
	for j := 0; j < 10; j++ {
		if math.IsNaN(x) {
			break
		}
		if x == 0 || x == 1 {
			return x
		}
		// err = Ibeta(x, a, b) - p
		err = GIBeta(x, a, b) - p
		fmt.Println(err / t)
		t = math.Exp(a1*math.Log(x) + b1*math.Log(1-x) + afac)
		u = err / t
		t = u

		x -= (u / (1 - 0.5*math.Min(1, u*(a1/x-b1/(1-x)))))
		if x <= 0 {
			x = 0.5 * (x + t)
		}
		if x >= 1 {
			x = 0.5 * (x + t + 1)
		}
		if math.Abs(t) < EPS*x && j > 0 {
			break
		}
	}

	return x
}

// Log-gamma function
// jStat.gammaln = function gammaln(x) {
//   var j = 0;
//   var cof = [
//     76.18009172947146, -86.50532032941677, 24.01409824083091,
//     -1.231739572450155, 0.1208650973866179e-2, -0.5395239384953e-5
//   ];
//   var ser = 1.000000000190015;
//   var xx, y, tmp;
//   tmp = (y = xx = x) + 5.5;
//   tmp -= (xx + 0.5) * Math.log(tmp);
//   for (; j < 6; j++)
//     ser += cof[j] / ++y;
//   return Math.log(2.5066282746310005 * ser / xx) - tmp;
// };

func Gammaln(x float64) float64 {
	//  TODO: pre-incr / post-incr
	cof := []float64{
		76.18009172947146, -86.50532032941677, 24.01409824083091,
		-1.231739572450155, 0.1208650973866179e-2, -0.5395239384953e-5,
	}
	ser := 1.000000000190015
	var xx, y, tmp float64
	y = x
	xx = x
	tmp = y + 5.5
	tmp -= (xx + 0.5) * math.Log(tmp)
	for j := 0; j < 6; j++ {
		y += 1
		ser += cof[j] / y
	}
	return math.Log(2.5066282746310005*ser/xx) - tmp
}

// // Returns the incomplete beta function I_x(a,b)
// jStat.ibeta = function ibeta(x, a, b) {
//   // Factors in front of the continued fraction.
//   var bt = (x === 0 || x === 1) ?  0 :
//     Math.exp(jStat.gammaln(a + b) - jStat.gammaln(a) -
//              jStat.gammaln(b) + a * Math.log(x) + b *
//              Math.log(1 - x));
//   if (x < 0 || x > 1)
//     return false;
//   if (x < (a + 1) / (a + b + 2))
//     // Use continued fraction directly.
//     return bt * jStat.betacf(x, a, b) / a;
//   // else use continued fraction after making the symmetry transformation.
//   return 1 - bt * jStat.betacf(1 - x, b, a) / b;
// };

func Ibeta(x, a, b float64) float64 {
	// Factors in front of the continued fraction.
	var bt float64
	if x == 0 || x == 1 {
		bt = 0
	} else {
		bt = math.Exp(Gammaln(a+b) - Gammaln(a) -
			Gammaln(b) + a*math.Log(x) + b*
			math.Log(1-x))
	}

	if x < 0 || x > 1 {
		panic("parameter out of range")
	}
	if x < (a+1)/(a+b+2) {
		// Use continued fraction directly.
		return bt * Betacf(x, a, b) / a
	}
	// else use continued fraction after making the symmetry transformation.

	return 1 - bt*Betacf(1-x, b, a)/b
}

// // Evaluates the continued fraction for incomplete beta function by modified
// // Lentz's method.
// jStat.betacf = function betacf(x, a, b) {
//   var fpmin = 1e-30;
//   var m = 1;
//   var qab = a + b;
//   var qap = a + 1;
//   var qam = a - 1;
//   var c = 1;
//   var d = 1 - qab * x / qap;
//   var m2, aa, del, h;
//
//   // These q's will be used in factors that occur in the coefficients
//   if (Math.abs(d) < fpmin)
//     d = fpmin;
//   d = 1 / d;
//   h = d;
//
//   for (; m <= 100; m++) {
//     m2 = 2 * m;
//     aa = m * (b - m) * x / ((qam + m2) * (a + m2));
//     // One step (the even one) of the recurrence
//     d = 1 + aa * d;
//     if (Math.abs(d) < fpmin)
//       d = fpmin;
//     c = 1 + aa / c;
//     if (Math.abs(c) < fpmin)
//       c = fpmin;
//     d = 1 / d;
//     h *= d * c;
//     aa = -(a + m) * (qab + m) * x / ((a + m2) * (qap + m2));
//     // Next step of the recurrence (the odd one)
//     d = 1 + aa * d;
//     if (Math.abs(d) < fpmin)
//       d = fpmin;
//     c = 1 + aa / c;
//     if (Math.abs(c) < fpmin)
//       c = fpmin;
//     d = 1 / d;
//     del = d * c;
//     h *= del;
//     if (Math.abs(del - 1.0) < 3e-7)
//       break;
//   }
//
//   return h;
// };

func Betacf(x, a, b float64) float64 {

	fpmin := 1e-30
	qab := a + b
	qap := a + 1
	qam := a - 1
	var c float64 = 1
	var d float64 = 1 - qab*x/qap
	var m2, aa, del, h float64

	// These q's will be used in factors that occur in the coefficients
	if math.Abs(d) < fpmin {
		d = fpmin
	}
	d = 1 / d
	h = d

	for m := 1; m <= 100; m++ {
		mF := float64(m)
		m2 = 2 * mF
		aa = mF * (b - mF) * x / ((qam + m2) * (a + m2))
		// One step (the even one) of the recurrence
		d = 1 + aa*d
		if math.Abs(d) < fpmin {
			d = fpmin
		}
		c = 1 + aa/c
		if math.Abs(c) < fpmin {
			c = fpmin
		}
		d = 1 / d
		h *= d * c

		aa = -(a + mF) * (qab + mF) * x / ((a + m2) * (qap + m2))
		// Next step of the recurrence (the odd one)
		d = 1 + aa*d
		if math.Abs(d) < fpmin {
			d = fpmin
		}
		c = 1 + aa/c
		if math.Abs(c) < fpmin {
			c = fpmin
		}
		d = 1 / d
		del = d * c
		h *= del
		if math.Abs(del-1.0) < 3e-7 {
			break
		}
	}

	return h
}
