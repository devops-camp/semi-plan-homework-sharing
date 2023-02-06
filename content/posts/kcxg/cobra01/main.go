package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var root = &cobra.Command{
	Use:   "person",
	Short: "打招呼",
	Run: func(cmd *cobra.Command, args []string) {
		Person(name, age)
	},
}

var (
	name = ""
	age  = 0
)

func init() {
	root.Flags().StringVarP(&name, "name", "n", "", "姓名")
	root.Flags().IntVarP(&age, "age", "a", 20, "年龄")
}

func Person(name string, age int) {
	fmt.Printf("%s 你好， 今年 %d 岁\n", name, age)
}

func Execute() error {
	return root.Execute()
}

func main() {
	err := Execute()
	if err != nil {
		log.Fatal(err)
	}
}
