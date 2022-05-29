package scalar

import (
	"fmt"
	"testing"
)

var (
	ints = map[string]int64{
		"first":  34,
		"second": 12,
	}

	floats = map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
)

// go test -v .\scalar\ -run TestNonGeneric
func TestNonGeneric(t *testing.T) {
	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),
		SumFloats(floats))
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// go test -v .\scalar\ -run TestGeneric
func TestGeneric(t *testing.T) {
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats),
	)
}
