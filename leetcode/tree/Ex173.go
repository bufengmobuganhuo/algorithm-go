package tree

type BSTIterator struct {
	inorder []int
}

func Constructor(root *TreeNode) (it BSTIterator) {
	it.inorderTravel(root)
	return it
}

func (this *BSTIterator) inorderTravel(root *TreeNode) {
	if root == nil {
		return
	}
	this.inorderTravel(root.Left)
	this.inorder = append(this.inorder, root.Val)
	this.inorderTravel(root.Right)
}

func (this *BSTIterator) Next() int {
	val := this.inorder[0]
	this.inorder = this.inorder[1:]
	return val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.inorder) > 0
}
