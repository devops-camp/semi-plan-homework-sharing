---
title: "Cobra05 - 子命令及子命令参数 （简单）"
subtitle: "Cobra05 Subcommand"
categories:
  - semi-plan 
tags:
  - semi-plan-202302-a
  - cobra05-subcommand
date: "2023-01-17T22:54:13+08:00"
lastmod: "2023-01-17T22:54:13+08:00"
toc: true
draft: false
hiddenFromHomePage: false
pinTop: true

originAuthor: homework
originLink: ""
---



# Cobra05 - 子命令及子命令参数 （简单）

## 提示

考虑 **零值** 

## 要求

在 04 的基础上， 增加子命令 Address

1. 增加 Gender 和 Account 两个参数。 默认账户余额 Account 100。

```go
Usage:
  greeting [command]

Available Commands:
  address     
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
      --account int   账户余额 (default 100)
      --age int       年龄 (default 20)
      --gender        性别
  -h, --help          help for greeting
      --name string   用户名

Use "greeting [command] --help" for more information about a command.
```

2. **增加** 家庭住址子命令。

```
家庭住址

Usage:
  greeting address [flags]

Flags:
      --emails strings   邮箱
  -h, --help             help for address
      --home string      家庭住址
```

3. 已知 zhangsan 是个烂赌鬼， 将其账户余额由参数控制 Account。 现在输钱了， 将账户的值设为 0。 

并打印如下信息

```
烂赌鬼 zhangsan 账户余额为 0 。
```

## 解题思路

1. 添加子命令
    ```go
    rootCmd.AddCommand(addressCmd)
   
   ```
2. 给子命令添加flag
    ```go
   //添加flag
   addressCmd.Flags().StringVar(&person.Email, "emails", "", "邮箱地址")
   addressCmd.Flags().StringVar(&person.Address, "home", "", "家庭住址")
   ```
   
3. 给根命令添加 Account flag
   ```go
   ...
   rootCmd.Flags().IntVar(&person.Account, "account", 100, "账户余额")
   ...
   ```