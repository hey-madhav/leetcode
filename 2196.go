package main

func createBinaryTree(descriptions [][]int) *TreeNode {
    nodes := make(map[int]*TreeNode)
    isChildren := make(map[int]bool)

    for _, d := range descriptions {
        parent := d[0]
        child := d[1]
        isLeft := d[2]

        if _, exists := nodes[parent]; !exists {
            nodes[parent] = &TreeNode{Val: parent}
        }

        if _, exists := nodes[child]; !exists {
            nodes[child] = &TreeNode{Val: child}
        }

        if isLeft == 1 {
            nodes[parent].Left = nodes[child]
        } else {
            nodes[parent].Right = nodes[child]
        }

        isChildren[child] = true
    }

    for value, node := range nodes {
        if !isChildren[value] {
            return node
        }
    }

    return nil
}