package ncpgraph

import (
	"testing"
)

// go test -v -run=TestCreateLineGraph
func TestCreateLineGraph(t *testing.T) {
	// go StartWebService("./images", ":8080")
	xvals := []float64{1, 2, 3, 4, 5, 6, 7}
	yvals := [][]float64{
		{0, 1, 2, 3, 4, 5, 6},
		{0, 3, 9, 8, 6, 7, 5},
		{0, 8, 10, 9, 3, 4, 6},
	}
	legends := []string{"havest aa", "bb", "cc"}
	CreateLineGraph("time", "windowsize", xvals, yvals, legends, "test")
	// ch := make(chan int)
	// <-ch
}
