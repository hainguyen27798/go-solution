package utils

import "fmt"

type TNode struct {
	leaf     bool
	keys     []int
	children []*TNode
}

type BTree struct {
	root *TNode
	t    int // minimum degree
}

func NewBTree(t int) *BTree {
	return &BTree{root: nil, t: t}
}

func (n *TNode) Search(k int) bool {
	i := 0
	for i < len(n.keys) && k > n.keys[i] {
		i++
	}

	if i < len(n.keys) && n.keys[i] == k {
		return true
	}

	if n.leaf {
		return false
	}

	return n.children[i].Search(k)
}

func (bt *BTree) Insert(k int) {
	if bt.root == nil {
		bt.root = &TNode{leaf: true, keys: []int{k}}
		return
	}

	if len(bt.root.keys) == 2*bt.t-1 {
		newRoot := &TNode{}
		newRoot.children = append(newRoot.children, bt.root)
		bt.splitChild(newRoot, 0)
		bt.root = newRoot
	}

	bt.insertNonFull(bt.root, k)
}

func (bt *BTree) insertNonFull(n *TNode, k int) {
	i := len(n.keys) - 1

	if n.leaf {
		n.keys = append(n.keys, 0)
		for i >= 0 && k < n.keys[i] {
			n.keys[i+1] = n.keys[i]
			i--
		}
		n.keys[i+1] = k
	} else {
		for i >= 0 && k < n.keys[i] {
			i--
		}
		i++

		if len(n.children[i].keys) == 2*bt.t-1 {
			bt.splitChild(n, i)
			if k > n.keys[i] {
				i++
			}
		}

		bt.insertNonFull(n.children[i], k)
	}
}

// CRITICAL FIX: Extract middle key BEFORE truncating child.keys
func (bt *BTree) splitChild(parent *TNode, index int) {
	t := bt.t
	child := parent.children[index]
	newChild := &TNode{leaf: child.leaf}

	midKey := child.keys[t-1] // Extract middle key first

	// Update parent with the middle key
	parent.keys = append(parent.keys, 0)
	copy(parent.keys[index+1:], parent.keys[index:])
	parent.keys[index] = midKey

	// Split keys and children AFTER extracting the middle key
	newChild.keys = append(newChild.keys, child.keys[t:]...)
	child.keys = child.keys[:t-1]

	if !child.leaf {
		newChild.children = append(newChild.children, child.children[t:]...)
		child.children = child.children[:t]
	}

	parent.children = append(parent.children, nil)
	copy(parent.children[index+2:], parent.children[index+1:])
	parent.children[index+1] = newChild
}

// Print the B-tree structure
func (bt *BTree) Print() {
	if bt.root == nil {
		fmt.Println("🌳 Tree is empty")
		return
	}
	bt.root.print("", true)
	fmt.Println("\nLegend:")
	fmt.Println("├── : Has sibling to the right")
	fmt.Println("└── : Last child in parent node")
	fmt.Println("│   : Vertical connector")
}

func (n *TNode) print(prefix string, isTail bool) {
	fmt.Printf("%s", prefix)
	connector := "├── "
	if isTail {
		connector = "└── "
	}
	fmt.Printf("%s%v\n", connector, n.keys)

	childPrefix := prefix
	if isTail {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	for i, child := range n.children {
		isLast := i == len(n.children)-1
		child.print(childPrefix, isLast)
	}
}
