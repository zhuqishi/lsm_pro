/*
 * @Author: your name
 * @Date: 2021-08-29 20:32:11
 * @LastEditTime: 2021-08-29 20:34:42
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /lsm_tree_pro/cmd/root.go
 */
package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsm_tree_pro",
	Short: "a demo for lsm_tree",
	Long:  "a demo for lsm_tree",
}
