/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package student

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get student info",
	Long:  `get student info`,
	Run: func(cmd *cobra.Command, args []string) {
		n, _ := cmd.Flags().GetString("name")
		a, _ := cmd.Flags().GetInt("age")
		g, _ := cmd.Flags().GetInt("grade")
		s := Student{
			Name:  n,
			Age:   a,
			Grade: g,
		}
		fmt.Println(s, "in get")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentPostRun in get")
	},
}

func init() {
	StudentCmd.AddCommand(getCmd)
	getCmd.Flags().String("test", "test", "test")
}
