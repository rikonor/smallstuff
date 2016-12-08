package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"github.com/gonum/matrix/mat64"
)

func main() {
	rand.Seed(time.Now().Unix())

	// Generate a list of numbers
	nFeatureVectors := 30

	trainingSet := []*FeatureVector{}
	for i := 0; i < nFeatureVectors; i++ {
		data := fmt.Sprintf("Training Vector %d", i)
		trainingSet = append(trainingSet, randomFeatureVector(data))
	}
	// fmt.Printf("%+v\n", trainingSet)
	fmt.Println("Initial training set:", trainingSet)

	// Sort the numbers by their distance from origin
	sort.Sort(ByOriginDistance(trainingSet))
	fmt.Println("\nTraining set sorted by distance from origin:", trainingSet)

	// Start with one candidate // TODO later more candidate to amortize initial sorting cost
	// Find the nearest neighbor
	c := randomFeatureVector("Candidate Vector")
	fmt.Println("\nCandidate Vector:", c)

	FindNearestNeighbor(c, trainingSet)
}

func FindNearestNeighbor(c *FeatureVector, td []*FeatureVector) *FeatureVector {
	// Binary Search for the vector with the closest distance to origin as c
	i := sort.Search(len(td), func(i int) bool {
		distA := mat64.Norm(c.Features, 2)
		distB := mat64.Norm(td[i].Features, 2)

		return distA < distB
	})

	// We know that c is between then td[i-1] and td[i]
	// We'll start sucking up the fv from left and right and check their distance to the candidate
	// Then we'll keep a tally of the closest one

	// setup aux data
	var distCVec mat64.Vector
	var distC, distOrigin, distOriginDiff float64

	lIdx, rIdx := i-1, i
	minDist := math.MaxFloat64
	var nearestNeighbor *FeatureVector
	for {
		// check left side
		if lIdx >= 0 {
			// distance to candidate
			distCVec.Reset()
			distCVec.SubVec(c.Features, td[lIdx].Features)
			distC = mat64.Norm(&distCVec, 2)

			// distance to origin
			distOrigin = mat64.Norm(td[lIdx].Features, 2)

			// difference in origin distances
			distOriginDiff = math.Abs(distOrigin - mat64.Norm(c.Features, 2))

			fmt.Printf("i=%d, dC=%.2f, dOrigin=%.2f, dOriginDiff=%.2f\n", lIdx, distC, distOrigin, distOriginDiff)
			if distOriginDiff > minDist {
				lIdx = 0
			} else {
				if distC < minDist {
					minDist = distC
					nearestNeighbor = td[lIdx]
				}
			}

			// set up the next left element
			lIdx--
		}

		// check right side
		if rIdx < len(td) {
			// distance to candidate
			distCVec.Reset()
			distCVec.SubVec(c.Features, td[rIdx].Features)
			distC = mat64.Norm(&distCVec, 2)

			// distance to origin
			distOrigin = mat64.Norm(td[rIdx].Features, 2)

			// difference in origin distances
			distOriginDiff = math.Abs(distOrigin - mat64.Norm(c.Features, 2))

			fmt.Printf("i=%d, dC=%.2f, dOrigin=%.2f, dOriginDiff=%.2f\n", rIdx, distC, distOrigin, distOriginDiff)
			if distOriginDiff > minDist {
				rIdx = len(td)
			} else {
				if distC < minDist {
					minDist = distC
					nearestNeighbor = td[rIdx]
				}
			}

			// set up the next right element
			rIdx++
		}

		// should we keep going?
		if lIdx < 0 && rIdx >= len(td) {
			break
		}
	}

	fmt.Println("nearest neighbor:", nearestNeighbor)

	return nil
}

func randomFeatureVector(data string) *FeatureVector {
	dim := 3
	var minVal, maxVal float64 = -10, 10

	vs := []float64{}
	for i := 0; i < dim; i++ {
		val := (maxVal-minVal)*rand.Float64() + minVal
		vs = append(vs, val)
	}

	return &FeatureVector{
		Data:     data,
		Features: mat64.NewVector(dim, vs),
	}
}

type FeatureVector struct {
	Data     string
	Features *mat64.Vector
}

func (fv FeatureVector) String() string {
	return fmt.Sprintf("Data: %s, Features: %v", fv.Data, *fv.Features)
}

type ByOriginDistance []*FeatureVector

func (a ByOriginDistance) Len() int      { return len(a) }
func (a ByOriginDistance) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByOriginDistance) Less(i, j int) bool {
	// sort by distance from origin ascending
	distA := mat64.Norm(a[i].Features, 2)
	distB := mat64.Norm(a[j].Features, 2)

	return distA < distB
}

/*
// Round 1 - Normal Nearest Neighbor

i Data: 3, Features:  0.72
  Data: 1, Features: -1.25
  Data: 7, Features:  2.16
  Data: 2, Features: -3.45
  Data: 8, Features: -3.97
  Data: 5, Features:  4.04
  Data: 0, Features: -4.46
  Data: C, Features: -4.94
  Data: 6, Features:  7.34
  Data: 9, Features: -8.33
j Data: 4, Features:  9.88

1.  compare with i (vec 0): distance 0.48. minimum distance is 0.48 for vec 0
2.  compare with j (vec 6): distance is 12.28
3.  compare with i (vec 5): distance is 8.98
4.  compare with j (vec 9): distance is 3.389
5.  compare with i (vec 8): distance is 0.97
6.  compare with j (vec 4): distance is 14.82
7.  compare with i (vec 2): distance is 1.49
8.  no more vectors on the right
9.  compare with i (vec 7): distance is 7.10
10. compare with i (vec 1): distance is 3.69
11. compare with i (vec 3): distance is 5.66
12. no more vectors on the left

nearest neighbor is vec 0 with distance 0.48

// Round 2 - Nearest Neighbor and stop in the middle

  Data: 3, Features:  0.72, dC=5.66,  OrigDist: 0.72, dOriginDiff=4.22
  Data: 1, Features: -1.25, dC=3.69,  OrigDist: 1.25, dOriginDiff=3.69
  Data: 7, Features:  2.16, dC=7.10,  OrigDist: 2.16, dOriginDiff=2.78
  Data: 2, Features: -3.45, dC=1.49,  OrigDist: 3.45, dOriginDiff=1.49
  Data: 8, Features: -3.97, dC=0.97,  OrigDist: 3.97, dOriginDiff=0.97
  Data: 5, Features:  4.04, dC=8.98,  OrigDist: 4.04, dOriginDiff=0.90
i Data: 0, Features: -4.46, dC=0.48,  OrigDist: 4.46, dOriginDiff=0.48
  Data: C, Features: -4.94, dC=0.00,  OrigDist: 4.94, dOriginDiff=0.00
j Data: 6, Features:  7.34, dC=12.28, OrigDist: 7.34, dOriginDiff=2.39
  Data: 9, Features: -8.33, dC=3.38,  OrigDist: 8.33, dOriginDiff=3.38
  Data: 4, Features:  9.88, dC=14.82, OrigDist: 9.88, dOriginDiff=4.94

1. compare with i (vec 0). their distance is 0.48 -> closest one we found so far
2. compare with j (vec 6). their origdist diff is 2.39, meaning at the least their diff is 2.39, so no use in continuing on the right side
3. compare with i (vec 5). their origdist diff is 0.90, meaning at the least their diff is 0.90, so no use in continuing on the left side

// 3 steps instead of 12

// Round 3 - Again with a different candidate

  Data: 3, Features:  0.72, dC=0.72,  OrigDist: 0.72, dOriginDiff=0.72
i Data: 1, Features: -1.25, dC=1.25,  OrigDist: 1.25, dOriginDiff=1.25
  Data: C, Features:  0.00, dC=0.00,  OrigDist: 0.00, dOriginDiff=0.00
j Data: 7, Features:  2.16, dC=2.16,  OrigDist: 2.16, dOriginDiff=2.16
  Data: 2, Features: -3.45, dC=3.45,  OrigDist: 3.45, dOriginDiff=3.45
  Data: 8, Features: -3.97, dC=3.97,  OrigDist: 3.97, dOriginDiff=3.97
  Data: 5, Features:  4.04, dC=4.04,  OrigDist: 4.04, dOriginDiff=4.04
  Data: 0, Features: -4.46, dC=4.46,  OrigDist: 4.46, dOriginDiff=4.46
  Data: 6, Features:  7.34, dC=7.34,  OrigDist: 7.34, dOriginDiff=7.34
  Data: 9, Features: -8.33, dC=8.33,  OrigDist: 8.33, dOriginDiff=8.33
  Data: 4, Features:  9.88, dC=9.88,  OrigDist: 9.88, dOriginDiff=9.88

1. compare with i (vec 1). their distance 1.25 -> closest one we found so far
2. compare with j (vec 7). their origdist is 2.16, meaning at the least their diff is 2.16, so no use in continuing on the right side
3. compare with i (vec 3). their origdist is 0.72, and their dist is 0.72 -> closest one we found so far
4. no more i

*/
