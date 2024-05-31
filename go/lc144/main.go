package main
 
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
 
func preorderTraversal(root *TreeNode) []int {
    out := make([]int, 0)
    if root == nil {
        return out
    }
    out = append(out, root.Val)
    out = append(out, preorderTraversal(root.Left)...)
    out = append(out, preorderTraversal(root.Right)...)
    return out
}

func main(){
    r3 := TreeNode{3, nil, nil}
    r2 := TreeNode{1, nil, &r3}
    root := TreeNode{2, &r2, nil}
    fmt.Println(preorderTraversal(root)
}