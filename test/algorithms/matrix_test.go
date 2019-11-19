package algorithms

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/dracit7/algorithms/algo"
	chart "github.com/wcharczuk/go-chart"
)

func BenchmarkMatrix(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testcase := genMatrixTestCase(1000, 100)
		algo.MatrixChainMultiOrder(testcase)
	}
}

func TestMatrix(t *testing.T) {

	var testcases [][]int
	for i := 1; i <= 100; i++ {
		testcases = append(testcases, genMatrixTestCase(10*i, 400))
	}

	points := chart.ContinuousSeries{
		XValues: []float64{},
		YValues: []float64{},
	}

	ite := time.Now()
	for i, tc := range testcases {
		algo.MatrixChainMultiOrder(tc)
		// t.Logf("%d %d\n", 10*(i+1), time.Now().Sub(ite).Nanoseconds())
		points.YValues = append(points.YValues, float64(time.Now().Sub(ite).Microseconds()))
		points.XValues = append(points.XValues, float64(10*(i+1)))
		ite = time.Now()
	}

	graph := chart.Chart{
		Series: []chart.Series{
			points,
		},
	}
	buf := bytes.NewBuffer([]byte{})
	graph.Render(chart.PNG, buf)
	ioutil.WriteFile("../../instances/img/test_matrix.png", buf.Bytes(), os.ModePerm)

}

func genMatrixTestCase(size, max int) []int {

	var testcase []int

	// Initialize the randnum generator
	rand.Seed(time.Now().UnixNano())

	// Generate the test case
	for i := 0; i < size; i++ {
		testcase = append(testcase, rand.Intn(max))
	}

	return testcase

}
