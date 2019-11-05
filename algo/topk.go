package algo

const partSize = 5

// KSmallest returns the kth smallest element
func KSmallest(E []Element, k int) Element {

	var M []Element
	var v Element

	// Recursive exit
	if len(E) <= partSize {
		InsertSort(&E)
		return E[k-1]
	}

	for i := 0; i < len(E); i += partSize {
		var m []Element
		for j := i; j < len(E) && j < i+partSize; j++ {
			m = append(m, E[j])
		}
		InsertSort(&m)
		M = append(M, m[len(m)/2])
	}

	if (len(E)/partSize)%2 == 1 {
		v = KSmallest(M, (1 + (len(E)/partSize)/2))
	} else {
		v = KSmallest(M, (len(E)/partSize)/2)
	}

	var j = partition(E, v)

	if k == j {
		return v
	} else if k < j {
		var A1 []Element
		for i := 0; i <= j-1; i++ {
			A1 = append(A1, E[i])
		}
		return KSmallest(A1, k)
	} else {
		var A2 []Element
		for i := j - 1; i != len(E); i++ {
			A2 = append(A2, E[i])
		}
		return KSmallest(A2, k-j+1)
	}

}
