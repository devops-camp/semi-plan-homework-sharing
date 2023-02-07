package main

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var config string

func init() {
	root.Flags().StringVarP(&config, "config", "c", "config.yml", "配置文件")
}

type Person struct {
	Name string `yaml:"name,omitempty" json:"name,omitempty"`
	Age  int    `yaml:"age,omitempty" json:"age,omitempty"`
}

func readConfig() *Person {
	person := &Person{}

	// 1. 读取文件
	b, err := os.ReadFile(config)
	if err != nil {
		panic(err)
	}

	// 2. 绑定参数
	err2 := yaml.Unmarshal(b, person)
	if err2 != nil {
		panic(err)
	}

	return person
}

// dumpConfig 保存文件
func dumpConfig(person *Person) {
	// 将结构体解析成 []byte
	b, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}

	// os.ModePerm => folder 755, file 644
	err2 := os.WriteFile("config.json", b, os.ModePerm)
	if err2 != nil {
		panic(err)
	}
}

func ReadWrite() {
	person := readConfig()
	dumpConfig(person)
}

var root = &cobra.Command{
	Use:   "rw",
	Short: "转换",
	Run: func(cmd *cobra.Command, args []string) {
		ReadWrite()
	},
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
