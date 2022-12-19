package labtest

import (
	"math"
	"testing"
	"time"
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

func FailAfter(t testing.TB, d time.Duration, f func()) {
	t.Helper()
	done := make(chan struct{}, 1)

	go func() {
		f()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Error("timed out")
	case <-done:
	}
}

func WalkSlice[T any](s []T, f func(elem T)) {
	length := len(s)
	for i := 0; i < length; i++ {
		f(s[i])
	}
}