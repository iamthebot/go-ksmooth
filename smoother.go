package ksmooth

import (
	"github.com/iamthebot/gostat/stat"
	"math"
)

type NWSmoother struct {
	Weights [][]float32
	Kernel  func(float32, float32) float32
	Radius  int //speed up computation by ignoring points outside the Radius.
}

//Allocates a new Nadraya-Watson kernel smoother using gaussian kernel
//It's your responsibility to make sure the kernel is valid
//Passing the requisite density (must be in (0.0,1.0) range) significantly speeds up smoothing by ignoring distant points
func NewNWGaussianSmoother(size int, bandwidth float32, density float32) (*NWSmoother, error) {
	//Create NWSmoother
	s := NWSmoother{
		Kernel: KernelGaussian,
	}

	//Generate weight matrix
	s.Weights = make([][]float32, size)
	for i := 0; i < size; i++ {
		s.Weights[i] = make([]float32, size)
		for t := 0; t < size; t++ {
			s.Weights[i][t] = s.Kernel(float32(t-i), bandwidth)
		}
	}

	//Calculate Radius
	//we use density as 2-tailed probability in InverseNormalCDF
	tail := (1.0 - density) / 2.0
	s.Radius = int(math.Max(float64(stat.NormalInv_CDF(float64(1.0-tail), float64(bandwidth))), 1.0))

	return &s, nil
}

func (s NWSmoother) Smooth(inputs []float32, outputs []float32) {
}
