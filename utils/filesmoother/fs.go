package main

import (
	"bufio"
	"fmt"
	"github.com/blacklabcapital/go-ksmooth"
	"log"
	"os"
	"strconv"
)

func ReadFloats(filename string) ([]float64, error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []float64
	for scanner.Scan() {
		x, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	r.Close()
	return result, scanner.Err()
}

func main() {
	log.Println("Loading data file")
	data, err := ReadFloats("input.csv")
	if err != nil {
		log.Fatalf("Couldn't read data file: %s\n", err.Error())
	}
	log.Println("Loaded data file")
	of, err := os.OpenFile("smooth.csv", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Couldn't open output file for writing: %s\n", err.Error())
	}
	log.Println("Creating smoother")
	s, err := ksmooth.NewNWGaussianSmoother(11000, 2500, 0.80)
	if err != nil {
		log.Fatalf("Couldn't create smoother: %s\n", err.Error())
	}
	log.Println("Created smoother")

	log.Println("Smoothing")
	for i := 0; i < 11000; i++ {
		p, _ := s.SmoothPoint(i, data, 11000)
		fmt.Fprintf(of, "%.4f\n", p)
	}
	log.Println("Done")
	of.Close()
}
