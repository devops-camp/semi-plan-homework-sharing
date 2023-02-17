package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

/*
@Auth: menah3m
@Desc:
*/
var name string
var age int

// rootCmd 根命令定义
var rootCmd = &cobra.Command{
	// 命令的相关文本信息
	Use:   "greeting",
	Short: "give me your name and age,I will greet you",
	Long:  `this is a CLI demo for devops camp homework.`,

	// 该命令的执行逻辑
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s 你好，今年 %v 岁\n", name, age)
	},
}

func init() {
	// 绑定标志
	rootCmd.Flags().IntVar(&age, "age", 20, "add info: age")
	rootCmd.Flags().StringVar(&name, "name", "", "add info: name")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("execute err:", err)
		os.Exit(1)
	}
}
