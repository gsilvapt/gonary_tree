package bintree

// Node represents a bit of the binary tree that takes integer values.
// The left reference to a new node represents inferior values, whereas the
// right reference contains the bigger nodes.
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// private method to abstract the necessary logic to add or create a new node
// in the tree. If the node does not exist, create one. If it does, then
// recursively call the Insert method to figure out where to put this new node
// (left or right).
func insertInNode(n *Node, k int) *Node {
	if n == nil {
		return &Node{Key: k}
	} else {
		return n.Insert(k)
	}
}

// Insert will try add k in this tree.
// If higher, then does right. Otherwise, tries inserting in the left.
// If Left or Right are not empty, it will call Insert recursively in those nodes until in finds a fitting place
func (n *Node) Insert(k int) *Node {
	if n.Key > k {
		node := insertInNode(n.Left, k)
		n.Left = node
	}

	if n.Key < k {
		node := insertInNode(n.Right, k)
		n.Right = node
	}

	return n
}

// Search will attempt to search k in this Node.
// If k is not equal to the value, then it will search right or left, if k is higher or lower (respectively).
// Returns true if there is a node with that value.
func (n *Node) Search(k int) bool {
	if n == nil {
		return false
	}

	if n.Key > k {
		return n.Left.Search(k)
	}

	if n.Key < k {
		return n.Right.Search(k)
	}

	return true
}
