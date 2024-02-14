package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	newnode := &TreeNode{Data: data}

	if root == nil {
		root = newnode
		return root
	}

	if data >= root.Data {
		if root.Right == nil {
			newnode.Parent = root
			root.Right = newnode
			return newnode
		} else {
			return BTreeInsertData(root.Right, data)
		}
	} else {
		if root.Left == nil {
			newnode.Parent = root
			root.Left = newnode
			return newnode
		} else {
			return BTreeInsertData(root.Left, data)
		}
	}
}
