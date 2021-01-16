package main

import (
	"net/http"
	"strings"
)

func main() {
	var tripper http.RoundTripper = &RoundTripCounter{}
	r, _ := http.NewRequest(http.MethodGet, "http://pluralsight.com", strings.NewReader("test call"))
	_, _ = tripper.RoundTrip(r)
}

type RoundTripCounter struct {
	count int
}

func (rt *RoundTripCounter) RoundTrip(*http.Request) (*http.Response, error) {
	rt.count += 1
	return nil, nil
}

// func main() {
// 	sv := simplemath.NewSemanticVersion(1, 2, 3)
// 	sv.IncrementMajor()
// 	println((sv.String()))
// }

// func main() {
// 	answer, err := simplemath.Divide(6, 2)
// 	if err != nil {
// 		fmt.Printf("An error occured %s\n", err.Error())
// 	} else {
// 		fmt.Printf("%f\n", answer)
// 	}

// 	numbers := []float64{12.2, 14, 16, 22, 4}
// 	total := simplemath.Sum(numbers...)
// 	fmt.Printf("total of sum: %f\n", total)
// }
