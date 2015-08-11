package ksmooth

import (
	"math"
)

var __OSQRT2PI__ float64 = 0.3989422804014326779399460599343818684758586311649346

func KernelGaussian(x float64, s float64) float64 {
	return __OSQRT2PI__ * (1.0 / s) * float64(math.Exp(-0.5*(math.Pow(float64(x), 2.0)/math.Pow(float64(s), 2.0))))
}
