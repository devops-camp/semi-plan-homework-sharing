package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

// the questions to ask
var qs = []*survey.Question{
	{
		// 1. Input 输入框
		Name: "id",
		Prompt: &survey.Input{
			Message: "Access Secret ID: ",
		},
		Validate: survey.Required,
	},
	{
		// 2. Password 密码输入框
		Name: "key",
		Prompt: &survey.Password{
			Message: "Access Secret Key: ",
		},
		Validate: survey.Required,
	},
	{
		// 3. Select 单选框
		Name: "region",
		Prompt: &survey.Select{
			Message: "Choose a region:",
			Options: []string{"cn-shanghai", "cn-hangzhou"},
			Default: "cn-hangzhou",
		},
	},
	{
		// 4. MultiSelect 多选框
		Name: "language",
		Prompt: &survey.MultiSelect{
			Message: "Supported Configure Language: ",
			Options: []string{"zh", "en", "jp"},
		},
	},
}

func start() {
	answers := struct {
		ID          string
		Key         string
		ChinaRegion string `survey:"region"`
		Language    []string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	dumpConfig(answers)
}

func dumpConfig(answer any) {
	// 将结构体解析成 []byte
	b, err := yaml.Marshal(answer)
	if err != nil {
		panic(err)
	}

	// os.ModePerm => folder 755, file 644
	err2 := os.WriteFile("config.yaml", b, os.ModePerm)
	if err2 != nil {
		panic(err)
	}
}

var root = &cobra.Command{
	Use:   "inter",
	Short: "交互",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func main() {
	err := root.Execute()
	if err != nil {
		return
	}

}
