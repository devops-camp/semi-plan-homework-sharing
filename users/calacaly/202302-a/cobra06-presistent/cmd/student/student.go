/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package student

import (
	"fmt"

	"github.com/spf13/cobra"
)

var student = Student{}

type Student struct {
	Name  string
	Age   int
	Grade int
}

// studentCmd represents the student command
var StudentCmd = &cobra.Command{
	Use:   "student",
	Short: "student get",
	Long:  `student get`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentPreRun in student")
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PreRun in student")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(student, "in student")
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PostRun in student")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentPostRun in student")
	},
}

func init() {
	StudentCmd.PersistentFlags().StringVarP(&student.Name, "name", "n", "alice", "设置学生名称")
	StudentCmd.PersistentFlags().IntVarP(&student.Age, "age", "a", 0, "设置学生年龄")
	StudentCmd.PersistentFlags().IntVarP(&student.Grade, "grade", "g", 0, "设置学生成绩")
}
