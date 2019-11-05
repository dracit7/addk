package algorithms

import (
	"math/rand"
	"testing"
	"time"

	"github.com/dracit7/algorithms/algo"
)

// Define a brute force algo for topk problem
// to generate standard solutions.
func topkBF(E []algo.Element, k int) algo.Element {
	shell := append([]algo.Element{}, E...)
	algo.InsertSort(shell)
	return shell[k-1]
}

func genTestCase(size, max int) ([]algo.Element, int) {

	var testcase []algo.Element

	// Initialize the randnum generator
	rand.Seed(time.Now().UnixNano())

	// Generate the test case
	for i := 0; i < size; i++ {
		testcase = append(testcase, element(rand.Intn(max)))
	}

	return testcase, rand.Intn(size)

}

func TestKthSmallest(t *testing.T) {

	testcase, k := genTestCase(10000, 10000)

	// Solve the problem using BruteForce and our topK algo separately
	std := topkBF(testcase, k)
	mysol := algo.KthSmallestCustom(testcase, k, 5)

	// Check the correctness of solution
	if std != mysol {
		t.Errorf(`
Testcase: %dth smallest in %v
Result should be: %d
Your solution is: %d
`, k, testcase, std, mysol)
	}

}

func BenchmarkWithPartSize3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testcase, k := genTestCase(10000, 10000)
		algo.KthSmallestCustom(testcase, k, 3)
	}
}

func BenchmarkWithPartSize5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testcase, k := genTestCase(10000, 10000)
		algo.KthSmallestCustom(testcase, k, 5)
	}
}

func BenchmarkWithPartSize7(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testcase, k := genTestCase(10000, 10000)
		algo.KthSmallestCustom(testcase, k, 7)
	}
}

func BenchmarkWithPartSize9(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testcase, k := genTestCase(10000, 10000)
		algo.KthSmallestCustom(testcase, k, 9)
	}
}

func BenchmarkWithPartSize11(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testcase, k := genTestCase(10000, 10000)
		algo.KthSmallestCustom(testcase, k, 11)
	}
}

func BenchmarkWithPartSize27(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testcase, k := genTestCase(10000, 10000)
		algo.KthSmallestCustom(testcase, k, 27)
	}
}
func BenchmarkWithPartSize31(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testcase, k := genTestCase(10000, 10000)
		algo.KthSmallestCustom(testcase, k, 31)
	}
}

func BenchmarkWithPartSize36(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testcase, k := genTestCase(10000, 10000)
		algo.KthSmallestCustom(testcase, k, 36)
	}
}

func BenchmarkWithPartSize50(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testcase, k := genTestCase(10000, 10000)
		algo.KthSmallestCustom(testcase, k, 50)
	}
}
