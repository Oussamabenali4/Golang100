package main

import (
	"fmt"
	piscine "tests/Quest12"
)

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	fmt.Println(piscine.BTreeLevelCount(root))
}
