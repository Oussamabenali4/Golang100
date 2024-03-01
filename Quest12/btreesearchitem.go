package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return root
	} else if elem > root.Data {
		return BTreeSearchItem(root.Right, elem)
	} else {
		BTreeSearchItem(root.Left, elem)
	}
	return nil
}
