# cobra01

title: "作业: cobra - 01 实现编译与参数绑定"
categories:
&#x20; \- 半月刊
tags:
&#x20; \- 半月刊202301下
&#x20; \- code
date: "2023-01-13T15:19:07+08:00"
lastmod: "2023-01-13T15:19:07+08:00"
toc: true
draft: false
hiddenFromHomePage: false

\#  提交作业修改一下内容
pinTop: false
originAuthor: calacaly
originLink: ""

***

# 作业: cobra - 01 实现编译与参数绑定

要求:

1.  使用 [https://github.com/spf13/cobra](https://github.com/spf13/cobra "https://github.com/spf13/cobra") 实现命令工具
2.  命令具有以下参数
    1.  `--name` 姓名
    2.  `--age` 年龄
3.  如果年龄为空， 默认为 20 岁。
4.  完成交叉编译脚本， 编译其他平台的二进制文件

```纯文本
-rwxr-xr-x  1 franktang  staff  4220672 Jan 13 15:35 greeting-darwin-amd64
-rwxr-xr-x  1 franktang  staff  4203442 Jan 13 15:35 greeting-darwin-arm64
-rwxr-xr-x  1 franktang  staff  4215010 Jan 13 15:35 greeting-linux-amd64
-rwxr-xr-x  1 franktang  staff  4157892 Jan 13 15:35 greeting-linux-arm64
```

1.  执行输出效果如下

```bash
$ ./out/greeting-darwin-arm64
 你好, 今年 20 岁

$ ./out/greeting-darwin-arm64 --age 30 --name zhangsan
zhangsan 你好, 今年 30 岁
```

## 解题思路

1.  查阅github文档学习cobra的使用

    <https://github.com/spf13/cobra>
    ```bash
    #安装
    go get -u github.com/spf13/cobra@latest
    #引用 import "github.com/spf13/cobra"

    #官方推荐使用cobra-cli工具，安装cobra-cli工具
    go install github.com/spf13/cobra-cli@latest

    ```
2.  使用cobra-cli快速生成

    <https://github.com/spf13/cobra-cli/blob/main/README.md>
    ```bash
    #进入到个人作业目录下，使用cobra-cli工具
    cobra-cli init

    ```
3.  Flags学习

    <https://github.com/spf13/cobra/blob/main/user_guide.md>
    > Persistent Flags
    ```go
    // A flag can be 'persistent', meaning that this flag will be available to the command it's assigned to as well as every command under that command. 
    // For global flags, assign a flag as a persistent flag on the root.

    //持久flags
    rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
    ```
    > Local Flags
    ```go
    //A flag can also be assigned locally, which will only apply to that specific command.

    //本地flags
    localCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")

    ```
    > Required Flags
    ```go
    //Flags are optional by default. If instead you wish your command to report an error when a flag has not been set, mark it as required:

    //标记为required的flags必须赋值，否则会报错
    rootCmd.Flags().StringVarP(&Region, "region", "r", "", "AWS region (required)")
    rootCmd.MarkFlagRequired("region")

    ```
    > Flags Groups
    ```go
    //If you have different flags that must be provided together (e.g. if they provide the --username flag they MUST provide the --password flag as well) then Cobra can enforce that requirement:

    //标记为requiredTogether的一组Flags需要同时赋值
    rootCmd.Flags().StringVarP(&u, "username", "u", "", "Username (required if password is set)")
    rootCmd.Flags().StringVarP(&pw, "password", "p", "", "Password (required if username is set)")
    rootCmd.MarkFlagsRequiredTogether("username", "password")

    //You can also prevent different flags from being provided together if they represent mutually exclusive options such as specifying an output format as either --json or --yaml but never both:

    //标记为MutuallyExclusive的一组Flags不能同时赋值

    rootCmd.Flags().BoolVar(&u, "json", false, "Output in JSON")
    rootCmd.Flags().BoolVar(&pw, "yaml", false, "Output in YAML")
    rootCmd.MarkFlagsMutuallyExclusive("json", "yaml")

    ```
    > 使用Flags实现参数绑定
    ```go
    // 官方给出的案例都使用的是指针保存参数的值，

    //Uint8Var有指针，Uint8无指针
    //我们也可以使用无指针的方式进行绑定参数，
    //需要将参数取出来的时候，通过cmd.Flags().Getxxx("key")

    //Uint8Var无短flag，Uint8VarP有短flag

    ```
4.  具体实现方式参考cmd/root.go代码
5.  交叉编译
    ```powershell
    #GOOS 设置操作系统类型 GOARCH 设置系统架构
    #windows环境操作示例
    $Env:GOOS="windows";$Env:GOARCH="amd64";
    #使用go env可以查看当前的go环境变量设置
    go build -o filename
    #修改GOOS和GOARCH，进行交叉编译
    ```
