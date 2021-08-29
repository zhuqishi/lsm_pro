/*
 * @Author: your name
 * @Date: 2021-08-29 20:54:45
 * @LastEditTime: 2021-08-30 00:00:38
 * @LastEditors: Please set LastEditors
 * @Description: Tree-related data struct
 * @FilePath: /lsm_tree_pro/commons/tree.go
 */
package lsm

import "fmt"

/*******************************************************************************************************
rbtree start
*******************************************************************************************************/

const (
	RED   = true
	BLACK = false
)

type RBNode struct {
	Val                 int64
	Color               bool
	Left, Right, Parent *RBNode
}

type RBTree struct {
	Root *RBNode
}

func (rbnode *RBNode) GetParentNode() *RBNode {
	if rbnode == nil {
		return nil
	}
	return rbnode.Parent
}

func (rbnode *RBNode) GetBorNode() *RBNode {
	if rbnode == nil {
		return nil
	}
	parent := rbnode.Parent
	if parent == nil {
		return nil
	}
	if rbnode == parent.Left {
		return parent.Right
	} else {
		return parent.Left
	}
}

func (rbnode *RBNode) GetUncleNode() *RBNode {
	return rbnode.GetParentNode().GetBorNode()
}

func (rbnode *RBNode) LeftRotate() (root *RBNode, err error) {
	if rbnode == nil || rbnode.Right == nil {
		err = fmt.Errorf("right child is nil")
		return
	}
	root = rbnode.Right
	rbnode.Right = root.Left
	root.Left = rbnode

	return
}

func (rbnode *RBNode) RightRotate() (root *RBNode, err error) {
	if rbnode == nil || rbnode.Left == nil {
		err = fmt.Errorf("left child is nil")
		return
	}
	root = rbnode.Left
	rbnode.Left = root.Right
	root.Right = rbnode

	return
}

func (rbtree *RBTree) InsertNode(pnode *RBNode, data int64) {
	if pnode.Val >= data {
		if pnode.Left != nil {
			rbtree.InsertNode(pnode.Left, data)
		} else {
			newNode := &RBNode{}
			newNode.Parent = pnode
			pnode.Left = newNode
			rbtree.CheckInsertNode(newNode)
		}
	} else {
		if pnode.Right != nil {
			rbtree.InsertNode(pnode.Right, data)
		} else {
			newNode := &RBNode{}
			newNode.Parent = pnode
			pnode.Right = newNode
			rbtree.CheckInsertNode(newNode)
		}
	}
}

func (rbtree *RBTree) CheckInsertNode(checknode *RBNode) {
	return
}

/*******************************************************************************************************
rbtree  end
*******************************************************************************************************/
