/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package student

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var itemCmd = &cobra.Command{
	Use:   "item",
	Short: "item",
	Long:  `item`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("PersistentPreRun in item")
	},
	Run: func(cmd *cobra.Command, args []string) {
		n, _ := cmd.Flags().GetString("name")
		a, _ := cmd.Flags().GetInt("age")
		g, _ := cmd.Flags().GetInt("grade")
		s := Student{
			Name:  n,
			Age:   a,
			Grade: g,
		}
		fmt.Println(s, "in item")
	},
}

func init() {
	getCmd.AddCommand(itemCmd)
}
