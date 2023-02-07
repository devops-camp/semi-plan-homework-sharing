/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var address = Address{}

type Address struct {
	Email string
	Home  string
}

// addressCmd represents the address command
var addressCmd = &cobra.Command{
	Use:   "address",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("email is: %s, home at : %s .\n", address.Email, address.Home)
	},
}

func init() {
	rootCmd.AddCommand(addressCmd)

	addressCmd.Flags().StringVar(&address.Email, "email", "", "邮箱")
	addressCmd.Flags().StringVar(&address.Home, "home", "", "家庭住址")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addressCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addressCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
