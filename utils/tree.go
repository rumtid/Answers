package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"unsafe"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Aton(s string) unsafe.Pointer {
	root, _ := construct(s)
	return unsafe.Pointer(root)
}

func construct(s string) (*TreeNode, string) {
	if s[0] != '[' {
		panic("missing '['")
	}
	if s[1] == ']' {
		return nil, s[2:]
	}
	i := 1
	for ; s[i] == '-' || (s[i] >= '0' && s[i] <= '9'); i++ {
	}
	root := new(TreeNode)
	root.Val, _ = strconv.Atoi(s[1:i])
	if s = s[i:]; s[0] == ' ' {
		root.Left, s = construct(s[1:])
		root.Right, s = construct(s[1:])
	}
	return root, s[1:]
}

func Ntoa(root interface{}) string {
	n := (*TreeNode)(unsafe.Pointer(
		reflect.ValueOf(root).Pointer(),
	))
	return fmt.Sprint(traverse(n))
}

func traverse(root *TreeNode) (tree []interface{}) {
	if root != nil {
		tree = append(tree, root.Val)
		if root.Left != nil || root.Right != nil {
			tree = append(tree, traverse(root.Left))
			tree = append(tree, traverse(root.Right))
		}
	}
	return
}
