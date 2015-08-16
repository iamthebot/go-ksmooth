# Golang Kernel Smoothing

##Introduction
This package implements functions for fast computation of the Nadaraya-Watson Kernel Estimator. You can use these functions for easy smoothing of go slices.

Currently only a gaussian kernel is implemented.

##Usage
```go
//make a new gaussian nadaraya-watson smoother
//bandwidth is the kernel bandwidth to use
//density is what two-tailed % of the kernel we want to use for smoothing.
s, err := NewNWGaussianSmoother(2000.0, 0.95)
if err != nil {
    //do something
}
```
