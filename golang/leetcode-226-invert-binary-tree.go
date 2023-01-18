// https://leetcode.com/problems/invert-binary-tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTreeImpl(root *TreeNode) *TreeNode {
    if (root == nil) {
        return root
    }
    t := invertTreeImpl(root.Left)
    root.Left = invertTreeImpl(root.Right)
    root.Right = t
    return root
}

func invertTree(root *TreeNode) *TreeNode {
    // recursive solution
    return invertTreeImpl(root)
}
