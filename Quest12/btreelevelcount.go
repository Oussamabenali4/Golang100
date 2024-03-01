package piscine

func BTreeLevelCount(root *TreeNode) int {
	// return the hight of the tree
	if root == nil {
		return 0
	}
	left := BTreeLevelCount(root.Right) + 1
	right := BTreeLevelCount(root.Left) + 1
	if left > right {
		return left
	}
	return right
}
