package tree

func DeleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	if root.Val < key {
		root.Right = DeleteNode(root.Right, key)
	} else if root.Val > key {
		root.Left = DeleteNode(root.Left, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		deletedNode := root
		root = findMin(deletedNode.Right)
		root.Right = deleteMin(deletedNode.Right)
		root.Left = deletedNode.Left
	}
	return root
}

func deleteMin(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return node.Right
	}
	node.Left = deleteMin(node.Left)
	return node
}

func findMin(root *TreeNode) *TreeNode {
	if root.Left == nil {
		return root
	}
	return findMin(root.Left)
}
