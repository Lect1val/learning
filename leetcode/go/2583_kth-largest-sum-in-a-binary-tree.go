package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
    if root == nil {
        return -1
    }

    queue := list.New()
    queue.PushBack(root)

    var levelSums []int

    for queue.Len() > 0 {
        levelSize := queue.Len()
        levelSum := 0

        for i := 0; i < levelSize; i++ {
            node := queue.Remove(queue.Front()).(*TreeNode)
            levelSum += node.Val

            if node.Left != nil {
                queue.PushBack(node.Left)
            }
            if node.Right != nil {
                queue.PushBack(node.Right)
            }
        }

        levelSums = append(levelSums, levelSum)
    }

    sort.Slice(levelSums, func(i, j int) bool {
        return levelSums[i] > levelSums[j]
    })

    if k > len(levelSums) {
        return -1
    }

    return int64(levelSums[k-1])
}
