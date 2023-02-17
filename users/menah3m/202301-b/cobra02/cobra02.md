---
title: "作业: cobra - 02 读取配置配置文件"
categories:
  - semi-plan 
tags:
  - semi-plan-202301-b 
  - cobra02
date: "2023-01-13T15:19:07+08:00"
lastmod: "2023-01-13T15:19:07+08:00"
toc: true
draft: false
hiddenFromHomePage: false

#  提交作业修改一下内容
pinTop: false
originAuthor: menah3m
originLink: "https://github.com/Menah3m"
---


# 作业: cobra - 02 读取配置配置文件

要求:

1. 使用 https://github.com/spf13/cobra 实现命令工具
2. 命令具有以下参数
    1. `--config` , `-c` 配置文件

**配置文件如下**

```yaml
# config.yml
name: zhangsan
age: 20
```

3. 将配置文件保存为 `JSON` 格式 

```bash
$ cat config.json
```

输出结果

```json
{
    "name":"zhangsan",
    "age": 20
}
```

## 解题思路

### 关键逻辑
> 此部分是command.Run需要实现的内容
1. 读取配置文件中的内容 
   - 使用 github.com/spf13/viper 实现
2. 将配置文件中的内容解析成 json 格式后写入 文件
   - 使用 标准库 encoding/json 实现 json 序列化 
   - 使用 标准库 io/ioutil 实现 写入文件 


### 核心代码
使用 cobra 读取用户提交的 config file 参数<br>同时在执行对应的命令前先初始化（即先使用viper读取配置文件内容）
```go
func init() {
	// 执行命令前会先执行该函数
	cobra.OnInitialize(initConfig)
	// 定义参数
	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "add target config file ")
	// 将参数设置为必需
	rootCmd.MarkFlagRequired("config")

}
```
viper 库会读取配置文件中的内容，并保存在 viper 对象中
```go
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
```
将命令执行的核心逻辑抽象成 **yml2json** 函数：<br>
由于viper对象中保存的是一个结构体数据，所以使用 `encoding/json`的 `MarshalIndent` 方法进行序列化，将结构体转换成 json 格式数据
，然后使用 `io/ioutil` 库中的`WriteFile` 方法将数据写入文件中
```go
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
```


### 需要注意
- flag 的缩略形式(如 `--config 的缩略形式是 -c`) 是通过使用 `command.Flags().xxxxxP()` 方法定义 
  ```go
   rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "add target config file ")
   ```
- json 序列化时使用 `MarshalIndent` 方法可以美化 json 数据

### 参考文档
- [Cobra 文档]()
- [Viper 文档]()
- [Go 中 json 数据的序列化与反序列化]()