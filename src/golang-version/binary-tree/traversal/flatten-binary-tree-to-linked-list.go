package algorithm

// reference the illustration of labuladong's answer
func flatten(root *TreeNode)  {
    if root == nil {return}
    flatten(root.Left)
    flatten(root.Right)
    
    l := root.Left
    r := root.Right
    
    root.Left = nil
    root.Right = l
    
    p := root
    for p.Right != nil {
        p = p.Right
    }
    p.Right = r
}