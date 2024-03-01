package piscine

import "fmt"

func DrawTree(root *TreeNode, indent string) {
	if root == nil {
		return
	}

	fmt.Printf("%s-> %s\n", indent, root.Data)

	// Recursively print the right subtree with indentation
	if root.Right != nil {
		fmt.Printf("%s|\n", indent)
		DrawTree(root.Right, indent+"|   ")
	}

	// Recursively print the left subtree with indentation
	if root.Left != nil {
		DrawTree(root.Right, indent)
	}
}
