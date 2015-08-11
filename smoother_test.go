package ksmooth

import (
	"fmt"
	"testing"
)

func TestNewNWGaussianSmoother(t *testing.T) {
	s, _ := NewNWGaussianSmoother(1000, 3600.0, 0.80)
	//test mean
	weight := s.Weights[999][500]
	if (weight < 0.00021318) || (weight > 0.00021338) {
		fmt.Printf("Wrong weight. Got %1.8f Expected 0.00021328\n", weight)
		t.FailNow()
	}

	//test that max weight occurs at correct index
	for i := range s.Weights {
		var max float64 = 0.0
		maxpoint := 0
		for j := range s.Weights[250] {
			if s.Weights[i][j] > max {
				max = s.Weights[i][j]
				maxpoint = j
			}
		}
		if maxpoint != i {
			fmt.Printf("Maximum weight found at wrong index! Got %d Expected %d\n", maxpoint, i)
			t.FailNow()
		} else if (max < 0.00022153) || (max > 0.00022173) {
			fmt.Printf("Wrong maximum weight found! Got %1.8f Expected 0.00022163\n", max)
			t.FailNow()
		}
	}
}
