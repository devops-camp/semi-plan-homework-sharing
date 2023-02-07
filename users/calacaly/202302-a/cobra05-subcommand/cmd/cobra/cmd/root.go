/*
Copyright © 2023 calacaly

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var person = Person{}

type Person struct {
	Name    string
	Age     uint8
	Gender  string /*可以自定义性别*/
	Account int
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cobra",
	Short: "cobra 子命令及子命令参数",
	Long:  `cobra 子命令及子命令参数`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		/* Without pointer
		age, err := cmd.Flags().GetUint8("age")
		if err != nil {
			fmt.Println(err)
			return
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println(err)
			return
		}*/
		if person.Name == "zhangsan" {
			person.Account = 0
			fmt.Printf("烂赌鬼 %s 账户余额为 %d .\n", person.Name, person.Account)
		} else {
			fmt.Printf("%s 你好, 今年 %d 岁, 性别是 %s , 账户余额为 %d .\n", person.Name, person.Age, person.Gender, person.Account)
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra01.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	/*With pointer function*/
	rootCmd.Flags().Uint8VarP(&person.Age, "age", "a", 20, "年龄（default 20）")
	rootCmd.Flags().StringVarP(&person.Name, "name", "n", "", "用户名")

	rootCmd.Flags().StringVarP(&person.Gender, "gender", "g", "", "性别")
	rootCmd.Flags().IntVar(&person.Account, "account", 100, "账户余额（default 100）")

	/*Without pointer function*/
	//rootCmd.Flags().Uint8P("age", "a", 20, "set your age")
	//rootCmd.Flags().StringP("name", "n", "", "set your name")
}
