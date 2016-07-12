package treap

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	ErrNotFound = errors.New("Key not found")
)

// Init seeds the PRNG
func Init() {
	rand.Seed(time.Now().UnixNano())
}

type Node struct {
	Priority int
	Key      int
	Value    int
	Left     *Node
	Right    *Node
}

func NewNode(key, value, p int) *Node {
	n := &Node{
		Priority: p,
		Key:      key,
		Value:    value,
		Left:     nil,
		Right:    nil,
	}

	return n
}

func (n *Node) String() string {
	leftKey := "<nil>"
	if n.Left != nil {
		leftKey = fmt.Sprint(n.Left.Key)
	}
	rightKey := "<nil>"
	if n.Right != nil {
		rightKey = fmt.Sprint(n.Right.Key)
	}

	return fmt.Sprintf("(Key=%d, value=%d, Left=%s, Right=%s)", n.Key, n.Value, leftKey, rightKey)
}

// Get the keys from the treap as is (non-ordered / breadth-first)
func (n *Node) RawList() []int {
	if n == nil {
		return []int{}
	}

	leftKeys := n.Left.SortedList()
	rightKeys := n.Right.SortedList()

	withLeftSide := append([]int{n.Key}, leftKeys...)
	return append(withLeftSide, rightKeys...)
}

// Get the keys from the treap in a sorted order
func (n *Node) SortedList() []int {
	if n == nil {
		return []int{}
	}

	leftKeys := n.Left.SortedList()
	rightKeys := n.Right.SortedList()

	withLeftSide := append(leftKeys, n.Key)
	return append(withLeftSide, rightKeys...)
}

func (n *Node) RotateLeft() *Node {
	nl := &Node{
		Priority: n.Left.Priority,
		Key:      n.Left.Key,
		Value:    n.Left.Value,
		Left:     n.Left.Left,
		Right:    n,
	}

	n.Left = n.Left.Right

	return nl
}

func (n *Node) RotateRight() *Node {
	nr := &Node{
		Priority: n.Right.Priority,
		Key:      n.Right.Key,
		Value:    n.Right.Value,
		Left:     n,
		Right:    n.Right.Right,
	}

	n.Right = n.Right.Left

	return nr
}

// Given a key, find the corresponding value
func (n *Node) Get(key int) (int, error) {
	if n == nil {
		return 0, ErrNotFound
	}

	if key < n.Key {
		return n.Left.Get(key)
	}

	if key > n.Key {
		return n.Right.Get(key)
	}

	return n.Value, nil
}

// Given a key, find if the treap contains it
func (n *Node) Contains(key int) bool {
	if n == nil {
		return false
	}

	if key < n.Key {
		return n.Left.Contains(key)
	}

	if key > n.Key {
		return n.Right.Contains(key)
	}

	return true
}

func (n *Node) Set(key, val int) *Node {
	return n.set(key, val, -1)
}

func (n *Node) set(key, val, p int) *Node {
	if p == -1 {
		p = rand.Int()
	}

	if n == nil {
		// Just generate a new node
		return NewNode(key, val, p)
	}

	if key < n.Key {
		// Recurse leftwards, then do a left-rotation if necessary
		n.Left = n.Left.set(key, val, p)
		if n.Left.Priority < n.Priority {
			return n.RotateLeft()
		}
		return n
	}

	if key > n.Key {
		// Recurse rightwards, then do a right-rotation if necessary
		n.Right = n.Right.set(key, val, p)
		if n.Right.Priority < n.Priority {
			return n.RotateRight()
		}
		return n
	}

	// Update existing node
	n.Value = val
	return n
}

// Get the height of the treap
func (n *Node) Height() int {
	if n == nil {
		return 0
	}

	leftHeight := n.Left.Height()
	rightHeight := n.Right.Height()

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

// Split the treap at the key (key assumed to not be in the treap)
func (n *Node) Split(key int) (leftTreap, rightTreap *Node) {
	x := n.set(key, 0, -1)
	return x.Left, x.Right
}

// Merge two treaps into one
func Merge(left, right *Node) *Node {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	if left.Priority < right.Priority {
		left.Right = Merge(left.Right, right)
		return left
	}

	right.Left = Merge(right.Left, left)
	return right
}

func (n *Node) Delete(key int) (*Node, error) {
	if n == nil {
		return nil, ErrNotFound
	}

	if key < n.Key {
		tmp, err := n.Left.Delete(key)
		n.Left = tmp
		return n, err
	}

	if key > n.Key {
		tmp, err := n.Right.Delete(key)
		n.Right = tmp
		return n, err
	}

	return Merge(n.Left, n.Right), nil
}
