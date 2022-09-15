package labtest

import (
	"math"
	"testing"
)

func AssertEqual[T comparable](t *testing.T, expected, actual T) {
	t.Helper()
	if expected != actual {
		t.Errorf("expected %v, but got %v", expected, actual)
	}
}

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Got unexpected error: %v", err)
	}
}

func AssertEqualFloat64(t *testing.T, receivedResult, expectedResult float64) {
	const equalityThreshold = 1e-7
	if !(math.Abs(receivedResult-expectedResult) < equalityThreshold) {
		t.Errorf("got result %f but want %f", receivedResult, expectedResult)
	}
}
