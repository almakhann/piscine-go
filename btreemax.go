package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	return cur
}
