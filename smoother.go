package ksmooth

import (
	"errors"
	"github.com/iamthebot/gostat/stat"
	"math"
)

type NWSmoother struct {
	Bandwidth float64
	Radius    int
	Kernel    func(float64, float64) float64
}

//Allocates a new Nadraya-Watson kernel smoother using gaussian kernel
func NewNWGaussianSmoother(bandwidth float64, density float64) (*NWSmoother, error) {
	//Create NWSmoother
	s := NWSmoother{
		Bandwidth: bandwidth,
		Kernel:    KernelGaussian,
	}

	//Calculate Radius
	//we use density as 2-tailed probability in InverseNormalCDF
	tail := (1.0 - density) / 2.0
	s.Radius = int(math.Max(float64(stat.NormalInv_CDF(float64(1.0-tail), float64(bandwidth/2.0))), 1.0))

	return &s, nil
}

//Computes Nadraya-Watson regression at given point
//x is the x value which we want smoothed
//inputs is the slice of y values.
//the length we should assume for the input array
func (s NWSmoother) SmoothPoint(x int, inputs []float64, length int) (float64, error) {
	if x < 0 {
		return 0.0, errors.New("x is out of range")
	} else if length <= 0 {
		return 0.0, errors.New("length must be a positive integer")
	}

	low := 0
	if x-s.Radius > low {
		low = x - s.Radius
	}
	high := length - 1

	if x+s.Radius/2 < high {
		high = x + s.Radius/2
	}

	numSum := 0.0
	denSum := 0.0
	for i := low; i <= high; i++ {
		weight := s.Kernel(float64(i-x), s.Bandwidth/2.0)
		numSum += weight * inputs[i]
		denSum += weight
	}

	return numSum / denSum, nil
}
