package main

import (
	"fmt"

	"mengyu.com/algorithm/leetcode/tree"
)

func main() {
	root := tree.TreeNode{
		Left: &tree.TreeNode{
			Val: 3,
			Left: &tree.TreeNode{
				Val: 2,
			},
			Right: &tree.TreeNode{
				Val: 4,
			},
		},
		Right: &tree.TreeNode{
			Val: 6,
			Right: &tree.TreeNode{
				Val: 7,
			},
		},
		Val: 5,
	}
	fmt.Println(tree.DeleteNode(&root, 3))
}
