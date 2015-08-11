package ksmooth

import (
	"fmt"
	"testing"
)

func TestKernelGaussian(t *testing.T) {
	//Case -1.96 with stdnormal
	p := KernelGaussian(-1.96, 1.0)
	if (p < 0.0584408) || (p > 0.0584410) {
		fmt.Printf("STDNORMPDF(-1.96) produced wrong value: %.7f expected: 0.584409", p)
		t.FailNow()
	}

	//Case 4.0 with stdnormal
	p = KernelGaussian(4.0, 1.0)
	if (p < 0.00013373) || (p > 0.00013393) {
		fmt.Printf("STDNORMPDF(4.0) produced wrong value: %.8f expected: 0.00013383", p)
		t.FailNow()
	}

	//Case 4000.0 with sigma=3600.0
	p = KernelGaussian(4000.0, 3600.0)
	if (p < 0.00005967) || (p > 0.00005987) {
		fmt.Printf("NORMPDF WITH SIGMA 3600.0 (4000.0) produced wrong value: %.8f expected: 0.00005977", p)
		t.FailNow()
	}

	//Case -14000.0 with sigma=8500.0
	p = KernelGaussian(-14000.0, 8500.0)
	if (p < 0.00001198) || (p > 0.00001218) {
		fmt.Printf("NORMPDF WITH SIGMA 8500.0 (-14000.0) produced wrong value: %.8f expected: 0.00001208", p)
		t.FailNow()
	}
}

func BenchmarkKernelGaussian(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = KernelGaussian(float64(i), float64(i))
	}
}
