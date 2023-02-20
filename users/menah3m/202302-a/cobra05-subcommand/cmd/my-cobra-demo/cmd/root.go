package cmd

import (
	"fmt"
	"github.com/menah3m/semi-plan-homework/users/menah3m/202302-a/cobra05-subcommand/pkg/greeting"
	"github.com/spf13/cobra"
	"os"
)

/*
@Auth: menah3m
@Desc:
*/

var person greeting.Person

// rootCmd 根命令定义
var rootCmd = &cobra.Command{
	// 命令的相关文本信息
	Use:   "greeting",
	Short: "give me your name and age,I will greet you",
	Long:  `this is a CLI demo for devops camp homework.`,

	// 该命令的执行逻辑
	Run: person.SetAccountTo0,
}

func init() {
	// 绑定标志
	rootCmd.Flags().IntVar(&person.Age, "age", 20, "年龄")
	rootCmd.Flags().StringVar(&person.Name, "name", "", "姓名")
	rootCmd.Flags().IntVar(&person.Account, "account", 100, "账户余额")
	rootCmd.Flags().StringVar(&person.Gender, "gender", "", "性别")

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("execute err:", err)
		os.Exit(1)
	}
}
