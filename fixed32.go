package fixed32

import "math"

type Fixed32 int32

const (
	fractionalBits = 16
	scale          = 1 << fractionalBits
)

var (
	zero = Fixed32(0)     // 0.0
	one  = Fixed32(scale) // 1.0
	min  = Fixed32(math.MinInt32)
	prec = Fixed32(1) // highest precision
)

func FromFloat32(f float32) Fixed32 {
	return Fixed32(f * float32(scale))
}

func FromFloat64(f float64) Fixed32 {
	return Fixed32(f * float64(scale))
}

func (f Fixed32) Float32() float32 {
	return float32(f) / float32(scale)
}

func (f Fixed32) Float64() float64 {
	return float64(f) / float64(scale)
}

func (f Fixed32) Add(other Fixed32) Fixed32 {
	return f + other
}

func (f Fixed32) Sub(other Fixed32) Fixed32 {
	return f - other
}

func (f Fixed32) Mul(other Fixed32) Fixed32 {
	return Fixed32((int64(f) * int64(other)) >> fractionalBits)
}

func (f Fixed32) Div(other Fixed32) Fixed32 {
	return Fixed32((int64(f) << fractionalBits) / int64(other))
}

func (f Fixed32) Abs() Fixed32 {
	if f < 0 {
		return -f
	}
	return f
}

func (f Fixed32) Exp() Fixed32 {
	result := one
	term := one
	x := f
	for i := 1; i < 10; i++ {
		term = term.Mul(x).Div(Fixed32(float32(i) * scale))
		result = result.Add(term)
		if term.Abs() < prec {
			break
		}
	}
	return result
}

func (f Fixed32) Pow(exponent Fixed32) Fixed32 {
	lnBase := f.Ln()
	x := lnBase.Mul(exponent)
	return x.Exp()
}

func (f Fixed32) Ln() Fixed32 {
	if f <= 0 {
		return min
	}
	x := float32(f) / float32(scale)
	y := (x - 1) / (x + 1)
	result := zero
	term := Fixed32(float32(scale) * y)
	for i := 1; i <= 15; i += 2 {
		result = result.Add(term.Div(Fixed32(i * scale)))
		term = term.Mul(Fixed32(float32(scale) * y * y)).Div(one)
		if term.Abs() < prec {
			break
		}
	}
	return result.Mul(Fixed32(2 * scale))
}
