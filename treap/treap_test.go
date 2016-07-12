package treap

import (
	"os"
	"testing"
)

func getIncTree(n int) *Node {
	var x *Node

	for i := 0; i < n; i++ {
		x = x.Set(i, i)
	}

	return x
}

func TestMain(m *testing.M) {
	Init()
	os.Exit(m.Run())
}

func TestCanAddToTreap(t *testing.T) {
	var x *Node

	x = x.Set(0, 0)

	if x.Key != 0 || x.Value != 0 {
		t.Error("Failed to add to treap")
	}
}

func TestHeightIsLogarithmic(t *testing.T) {
	N := 100

	h1 := getIncTree(N).Height()
	h2 := getIncTree(N * N).Height()

	r := float64(h2) / float64(h1)

	// r should get asymptotically close to 2 as N grows
	if r > 3.0 {
		t.Error("Treap height is not approximately logarithmic")
	}
}
