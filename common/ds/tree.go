package ds

type TreeNode[T any] struct {
	Data T

	Parent   *TreeNode[T]
	Children []*TreeNode[T]
}

func NewTreeNode[T any](data T, parent *TreeNode[T]) *TreeNode[T] {
	return &TreeNode[T]{
		Data:   data,
		Parent: parent,
	}
}

func (tn *TreeNode[T]) Insert(data T) *TreeNode[T] {
	node := NewTreeNode(data, tn)

	tn.Children = append(tn.Children, node)

	return node
}

type Tree[T any] struct {
	Root *TreeNode[T]

	TraversalMethod uint8
}

func NewTree[T any](root *TreeNode[T]) *Tree[T] {
	return &Tree[T]{
		Root: root,
	}
}

func (t *Tree[T]) TraverseDF(cb func(node *TreeNode[T])) {
	t.recursiveDF(cb, t.Root)
}

func (t *Tree[T]) recursiveDF(cb func(node *TreeNode[T]), currentNode *TreeNode[T]) {
	for _, child := range currentNode.Children {
		t.recursiveDF(cb, child)
	}

	cb(currentNode)
}

func (t *Tree[T]) TraverseBF(cb func(node *TreeNode[T])) {
	queue := NewQueue[*TreeNode[T]]()

	queue.Enqueue(t.Root)

	currentNode, err := queue.Dequeue()
	if err != nil {
		panic(err)
	}

	for currentNode != nil {
		for _, node := range currentNode.Children {
			queue.Enqueue(node)
		}

		cb(currentNode)
		currentNode, err = queue.Dequeue()

		if _, ok := err.(*QueueEmptyException); err != nil && !ok {
			panic(err)
		}
	}
}
