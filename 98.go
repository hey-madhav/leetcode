package main

func isValidBST(root *TreeNode) bool {
    return isValidBSTHelper(root, nil, nil)
}

func isValidBSTHelper(root *TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}

	if (min != nil && root.Val <= *min) || (max != nil && root.Val >= *max) {
		return false
	}

	return isValidBSTHelper(root.Left, min, &root.Val) && isValidBSTHelper(root.Right, &root.Val, max)
}