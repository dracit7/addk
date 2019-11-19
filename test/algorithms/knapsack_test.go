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

func TestKnapSack01(t *testing.T) {

	var points []chart.ContinuousSeries
	for i := 0; i < 30; i++ {
		points = append(points, chart.ContinuousSeries{
			XValues: []float64{},
			YValues: []float64{},
		})
	}

	graph := chart.Chart{
		Series: []chart.Series{},
	}

	for w := range points {

		var testcases [][]algo.Item
		for i := 1; i <= 100; i++ {
			testcases = append(testcases, genKnapsackTestCase(10*i, (w+1)*10, 100))
		}

		ite := time.Now()
		for i, tc := range testcases {
			algo.Knapsack01(tc, (w+1)*1000)
			// t.Logf("%d %d\n", 10*(i+1), time.Now().Sub(ite).Nanoseconds())
			points[w].YValues = append(points[w].YValues, float64(time.Now().Sub(ite).Microseconds()))
			points[w].XValues = append(points[w].XValues, float64(10*(i+1)))
			ite = time.Now()
		}

		graph.Series = append(graph.Series, points[w])

	}

	buf := bytes.NewBuffer([]byte{})
	graph.Render(chart.PNG, buf)
	ioutil.WriteFile("../../instances/img/test_knapsack.png", buf.Bytes(), os.ModePerm)

}

func genKnapsackTestCase(size, maxW, maxV int) []algo.Item {

	var testcase []algo.Item

	// Initialize the randnum generator
	rand.Seed(time.Now().UnixNano())

	// Generate the test case
	for i := 0; i < size; i++ {
		testcase = append(testcase, item{
			rand.Intn(maxV),
			rand.Intn(maxW),
		})
	}

	return testcase

}
