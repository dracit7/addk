package algorithms

import (
	"math/rand"
	"testing"
	"time"

	"github.com/dracit7/algorithms/algo"
)

const (
	maxint = 100
)

func topkBF(E []algo.Element, k int) algo.Element {
	algo.InsertSort(&E)
	return E[k-1]
}

func TestSparseInt(t *testing.T) {

	var testcase []algo.Element

	rand.Seed(time.Now().UnixNano())

	size := rand.Intn(maxint)
	for i := 0; i < size; i++ {
		testcase = append(testcase, element(rand.Intn(maxint*5)))
	}

	k := rand.Intn(size)

	std := topkBF(testcase, k)
	mysol := algo.KSmallest(testcase, k)

	if std != mysol {
		t.Errorf(`
Testcase: %dth smallest in %v
Result should be: %d
Your solution is: %d
`, k, testcase, std, mysol)
	}

}
