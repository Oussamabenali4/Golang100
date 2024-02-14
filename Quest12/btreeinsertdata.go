package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root.Parent == nil {
		root.Parent = root
	}
	newnode := TreeNode{Data: data}
	if data > root.Data {
		if root.Right == nil {
			root.Right = &newnode
			return root
		} else {
			return BTreeInsertData(root.Right, data)
		}

	} else {
		if root.Left == nil {
			root.Left = &newnode
			return root
		} else {
			return BTreeInsertData(root.Left, data)
		}

	}
}
