package algorithm

// 把每一层的节点看成一个整体，traverse(root.Left, root.Right)要做的就是将当前根节点的左子节点的next指针指向其右子节点

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	dfsConnect(root.Left, root.Right)
	return root
}

func dfsConnect(Node1 *Node, Node2 *Node) {
	if Node1 == nil || Node2 == nil {
		return
	}

	Node1.Next = Node2

	dfsConnect(Node1.Left, Node1.Right)
	dfsConnect(Node2.Left, Node2.Right)

	dfsConnect(Node1.Right, Node2.Left)
}
