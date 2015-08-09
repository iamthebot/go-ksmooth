package ksmooth

import (
	"math"
)

var __OSQRT2PI__ float32 = 0.3989422804014326779399460599343818684758586311649346

func KernelGaussian(x float32, s float32) float32 {
	return __OSQRT2PI__ * (1.0 / s) * float32(math.Exp(-0.5*(math.Pow(float64(x), 2.0)/math.Pow(float64(s), 2.0))))
}
