package solution_search

import (
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

func SimpleMultipliers(x int) {
	FSimpleMultipliers(os.Stdout, x)
}

func FSimpleMultipliers(w io.Writer, x int) {
	var elems []float64
	for m := 0; m <= int(log(7, float64(x))); m++ {
		for l := 0; l <= int(log(5, float64(x))); l++ {
			for k := 0; k <= int(log(3, float64(x))); k++ {
				testElem := math.Pow(3, float64(k)) * math.Pow(5, float64(l)) * math.Pow(7, float64(m))
				if testElem > float64(x) {
					break
				}
				elems = append(elems, testElem)

			}
		}

	}
	sort.Float64s(elems)
	for _, ch := range elems {
		fmt.Fprintln(w, ch)
	}
}

func log(base float64, n float64) float64 {
	return math.Log(n) / math.Log(base)
}
