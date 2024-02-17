package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	current := root
	if current != nil {
		BTreeApplyInorder(current.Left, f)
		f(current.Data)
		BTreeApplyInorder(current.Right, f)
	}
}
