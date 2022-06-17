package main

import "fmt"

func main() {
	// Init a map for the integers
	ints := map[string]int64{
		"first":  50,
		"second": 10,
	}

	// Init a map of the floats
	floats := map[string]float64{
		"first":  21.73,
		"second": 7.44,
	}

	// Call the non-generic functions
	fmt.Printf("Using non-generic sums: [%v] and [%v]\n",
		SumInts(ints),
		SumFloats(floats))

}

// SumInts adds together the values of map m
func SumInts(inputMap map[string]int64) int64 {
	// sum variable holds the runnning total of the map elements
	var sum int64
	for _, element := range inputMap {
		sum += element
	}
	return sum
}

// SumFloats adds together the values of map m
func SumFloats(m map[string]float64) float64 {
	var sum float64
	for _, element := range m {
		sum += element
	}
	return sum
}
