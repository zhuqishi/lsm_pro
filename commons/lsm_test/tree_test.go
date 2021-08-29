/*
 * @Author: your name
 * @Date: 2021-08-30 00:29:28
 * @LastEditTime: 2021-08-30 00:33:50
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /lsm_tree_pro/commons/lsm_test/tree_test.go
 */
package lsm_test

import (
	"lsm_tree_pro/commons/lsm"
	"testing"
)

func TestInit(t *testing.T) {
	a := make([]*lsm.RBNode, 10)
	b := "a"
	t.Log(a)
	t.Log(b)
}
