package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

/*
@Auth: menah3m
@Desc:
*/
var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "transform config file layout from yml to json",
	Short: "import configuration from a yml file,then export to a json file",
	Long:  `this is a CLI demo for devops camp homework.`,
	Run:   yml2json,
}

func init() {
	// 执行命令前会先执行该函数
	cobra.OnInitialize(initConfig)
	// 定义参数
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "add target config file ")
	// 将参数设置为必需
	rootCmd.MarkFlagRequired("config")

}

func initConfig() {
	// 使用viper库 读取配置
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		if err := viper.ReadInConfig(); err != nil {
			fmt.Println("can not read config from file:", err)
			os.Exit(1)
		}
	}
}

func yml2json(cmd *cobra.Command, args []string) {
	// 将viper获取到的配置转为json格式
	jsonByte, err := json.MarshalIndent(viper.AllSettings(), "", "\t")
	if err != nil {
		fmt.Println("transform err:", err)
	}
	// 将转换之后的内容写入文件
	err = ioutil.WriteFile("config.json", jsonByte, os.ModePerm)
	if err != nil {
		fmt.Println("write file err:", err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("execute err:", err)
		os.Exit(1)
	}
}
