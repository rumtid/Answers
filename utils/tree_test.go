package utils

import (
	"fmt"
	"testing"
)

func TestConstruct(t *testing.T) {
	var root *TreeNode
	root, _ = construct("[]")
	if root != nil {
		t.Error(`construct("[]") =`, root, "; want <nil>")
	}
	root, _ = construct("[2]")
	if root == nil || root.Val != 2 || root.Left != nil || root.Right != nil {
		t.Error(`construct("[2]") =`, root, "; want [2]")
	}
	root, _ = construct("[3 [] [4]]")
	if root == nil || root.Val != 3 || root.Left != nil || root.Right == nil ||
		root.Right.Val != 4 || root.Right.Left != nil || root.Right.Right != nil {
		t.Error(`construct("[3 [] [4]]") =`, root, "; want [3 [] [4]]")
	}
}

func TestTraverse(t *testing.T) {
	var tree string
	tree = fmt.Sprint(traverse(nil))
	if tree != "[]" {
		t.Error(`traverse(nil) =`, tree, "; want []")
	}
	tree = fmt.Sprint(traverse(&TreeNode{Val: 2}))
	if tree != "[2]" {
		t.Error(`traverse("[2]") =`, tree, "; want [2]")
	}
	tree = fmt.Sprint(traverse(&TreeNode{Val: 3,
		Right: &TreeNode{Val: 4},
	}))
	if tree != "[3 [] [4]]" {
		t.Error(`traverse("[3 [] [4]]") =`, tree, "; want [3 [] [4]]")
	}
}
