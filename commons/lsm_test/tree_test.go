/*
 * @Author: your name
 * @Date: 2021-08-30 00:29:28
 * @LastEditTime: 2021-09-01 23:13:55
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /lsm_tree_pro/commons/lsm_test/tree_test.go
 */
package lsm_test

import (
	"fmt"
	"lsm_tree_pro/commons/lsm"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func TestInit(t *testing.T) {
	a := make([]*lsm.RBNode, 10)
	b := "a"
	t.Log(a)
	fmt.Println(a)
	t.Log(b)
}

func TestRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		a := rand.Intn(10)
		fmt.Println(a)
	}

}

func TestParse(t *testing.T) {
	IntA, _ := strconv.ParseInt("003", 10, 64)
	fmt.Println(IntA)
}

func TestSplit(t *testing.T) {
	a := "1.20.001.23.120"
	contents := strings.Split(a, ".")
	fmt.Println(contents)
}

func TestDiv(t *testing.T) {

	a := 2234
	b := 1000
	c := a / b
	fmt.Println(c)
}
