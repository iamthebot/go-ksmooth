package ksmooth

import (
	"fmt"
	"testing"
)

func TestNewNWGaussianSmoother(t *testing.T) {
	s, _ := NewNWGaussianSmoother(1000, 3600.0, 0.80)
	if s.Radius != 4613 {
		fmt.Printf("Got wrong radius for gaussian smoother! Got %d Expected 4613\n", s.Radius)
	}

	//test mean
	weight := s.Weights[999][500]
	if (weight < 0.00010965) || (weight > 0.00010985) {
		fmt.Printf("Wrong weight. Got %1.8f Expected 0.00010975\n", weight)
		t.FailNow()
	}

	//test that max weight occurs at correct index
	for i := range s.Weights {
		var max float32 = 0.0
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
		} else if (max < 0.00011072) || (max > 0.00011092) {
			fmt.Printf("Wrong maximum weight found! Got %1.8f Expected 0.00011082\n", max)
			t.FailNow()
		}
	}
}
