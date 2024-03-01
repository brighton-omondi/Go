package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	newNode := &TreeNode{Data: data}

	// If the tree is empty, make the new node the root
	if root == nil {
		return newNode
	}

	// Traverse the tree to find the appropriate position to insert the new node
	current := root
	var parent *TreeNode
	for current != nil {
		parent = current
		if data < current.Data {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	// Set the parent of the new node
	newNode.Parent = parent
	// Insert the new node as a child of the appropriate parent
	if data < parent.Data {
		parent.Left = newNode
	} else {
		parent.Right = newNode
	}

	return root
}
