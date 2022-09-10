package algorithm

// 把每一层的节点看成一个整体，traverse(root.Left, root.Right)要做的就是将当前根节点的左子节点的next指针指向其右子节点

func connect(root *Node) *Node {
    if root == nil {return nil}
    dfs(root.Left, root.Right)
    return root
}

func dfs(node1 *Node, node2 *Node) {
    if node1 == nil || node2 == nil {return}
    
    node1.Next = node2
    
    dfs(node1.Left, node1.Right)
    dfs(node2.Left, node2.Right)
    
    dfs(node1.Right, node2.Left)
}