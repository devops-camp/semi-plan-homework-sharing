package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

/*
   @Auth: menah3m
   @Desc:
*/

var addressCmd = &cobra.Command{
	Use:   "address",
	Long:  "set email address and home address",
	Short: "set address",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is address command")
	},
}

func init() {
	rootCmd.AddCommand(addressCmd)
	addressCmd.Flags().StringVar(&person.Email, "emails", "", "邮箱地址")
	addressCmd.Flags().StringVar(&person.Address, "home", "", "家庭住址")
}
