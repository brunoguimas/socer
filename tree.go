package socer

// Tree represents a simple generic binary search tree.
// The zero value of a BinaryTree is a empty binary tree ready to use.
type Tree[T any] struct {
	// root node of the BST.
	root *TreeNode[T]
}

// TreeNode represents a node in a binary tree.
type TreeNode[T any] struct {
	// Value of the node.
	Value T
	// Left child node.
	Left *TreeNode[T]
	// Right child node.
	Right *TreeNode[T]
}

// Insert value v in t. Computes in O(n).
func (t *Tree[T]) Insert(v T) *TreeNode[T] {
	n := &TreeNode[T]{Value: v}

	if t.root == nil {
		t.root = n
		return t.root
	}

	var q Queue[*TreeNode[T]]
	q.Enqueue(t.root)

	for !q.IsEmpty() {
		v, _ := q.Dequeue()

		if v.Left == nil {
			v.Left = n
			return v.Left
		}

		q.Enqueue(v.Left)

		if v.Right == nil {
			v.Right = n
			return v.Right
		}

		q.Enqueue(v.Right)
	}

	return nil
}

// Root returns the root of t.
func (t *Tree[T]) Root() *TreeNode[T] {
	return t.root
}

// Iter returns a new TreeIterator.
func (t *Tree[T]) Iter() *TreeIterator[T] {
	var q Queue[*TreeNode[T]]

	q.Enqueue(nil)
	enqueueTreeNode(t.root, &q)

	return &TreeIterator[T]{
		remaining: &q,
	}
}

// enqueueTreeNode enqueues n and its childs in q.
func enqueueTreeNode[T any](n *TreeNode[T], q *Queue[*TreeNode[T]]) {
	if n == nil {
		return
	}

	q.Enqueue(n)
	enqueueTreeNode(n.Left, q)
	enqueueTreeNode(n.Right, q)
}

// TreeIterator represents a iterator for a tree.
//
// To iterate over a stack (where s is a *Stack):
//
//	it := s.Iter()
//	for it.Next() {
//		// do something with it.Value()
//	}
type TreeIterator[T any] struct {
	remaining *Queue[*TreeNode[T]]
}

func (t *TreeIterator[T]) Next() bool {
	t.remaining.Dequeue()

	n, _ := t.remaining.Front()
	return n != nil
}

func (t *TreeIterator[T]) Value() T {
	n, _ := t.remaining.Front()

	return n.Value
}

func (t *TreeIterator[T]) Node() *TreeNode[T] {
	n, _ := t.remaining.Front()

	return n
}
