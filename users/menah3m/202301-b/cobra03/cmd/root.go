package cmd

import (
	"fmt"
	survey2 "github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"os"
)

/*
   @Auth: menah3m
   @Desc:
*/

var demoAnswers = struct {
	Name          string
	FavoriteColor string `survey:"color"`
}{}

var rootCmd = &cobra.Command{
	Use:   "survey",
	Short: "An Interactive CLI demo",
	Long:  "An Interactive CLI demo",
	Run:   RunDemo,
}

var demoQuestionSet = []*survey2.Question{
	{
		Name:      "name",
		Prompt:    &survey2.Input{Message: "What is your name?"},
		Validate:  survey2.Required,
		Transform: survey2.Title,
	},
	{
		Name: "color",
		Prompt: &survey2.Select{
			Message: "Choose a color:",
			Options: []string{"red", "blue", "green"},
			Default: "red",
		},
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("command run failed,err:", err)
		os.Exit(1)
	}
}

func RunDemo(cmd *cobra.Command, args []string) {
	err := survey2.Ask(demoQuestionSet, &demoAnswers)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s's favorite color is %s.\n", demoAnswers.Name, demoAnswers.FavoriteColor)
}
