package cmd

import (
	"encoding/json"
	"fmt"
	survey2 "github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

/*
@Auth: menah3m
@Desc:
*/
var profile string

var profileAnswers struct {
	AccessKeyID         string `survey:"access_key_id"`
	AccessKeySecret     string `survey:"access_key_secret"`
	DefaultRegionID     string `survey:"default_region_id"`
	DefaultOutputFormat string `survey:"default_output_format"`
	DefaultLanguage     string `survey:"default_language"`
}

var configureCmd = &cobra.Command{
	Use:   "configure",
	Long:  "profile configure cli by interactive cli demo",
	Short: "profile configure cli",
	Run:   Configure,
}

var profileQuestionSet = []*survey2.Question{
	{
		Name:     "access_key_id",
		Prompt:   &survey2.Input{Message: "Access Key ID[]"},
		Validate: survey2.Required,
	},
	{
		Name:     "access_key_secret",
		Prompt:   &survey2.Input{Message: "Access Key Secret[]"},
		Validate: survey2.Required,
	},
	{
		Name:     "default_region_id",
		Prompt:   &survey2.Input{Message: "Default Region ID[]"},
		Validate: survey2.Required,
	},
	{
		Name:     "default_output_format",
		Prompt:   &survey2.Input{Message: "Default Output Format[json]", Default: "json (Only upport json)"},
		Validate: survey2.Required,
	},
	{
		Name:     "default_language",
		Prompt:   &survey2.Input{Message: "Default Language?[zh|en]"},
		Validate: survey2.Required,
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
	configureCmd.Flags().StringVarP(&profile, "profile", "p", "", "target ")
}

func Configure(cmd *cobra.Command, args []string) {
	fmt.Printf("Configuring profile '%s' in authenticate mode...\n", profile)
	err := survey2.Ask(profileQuestionSet, &profileAnswers)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	jsonByte, err := json.MarshalIndent(profileAnswers, "", "\t")

	err = ioutil.WriteFile(profile+".profile", jsonByte, os.ModePerm)
	if err != nil {
		fmt.Println("write profile err:", err)
		os.Exit(1)
	}
}
