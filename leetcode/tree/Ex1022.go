package tree

func Run(root *TreeNode) int {
	return sumRootToLeaf(root)
}

func sumRootToLeaf(root *TreeNode) int {
	return sumRootToLeafDfs(root, 0)
}

func sumRootToLeafDfs(root *TreeNode, track int) int {
	track = (track << 1) | root.Val
	if root.Left == nil && root.Right == nil {
		return track
	}
	var left, right int
	if root.Left != nil {
		left = sumRootToLeafDfs(root.Left, track)
	}
	if root.Right != nil {
		right = sumRootToLeafDfs(root.Right, track)
	}

	return left + right
}
