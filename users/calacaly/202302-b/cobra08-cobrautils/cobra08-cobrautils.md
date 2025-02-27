---
title: "cobra08 - 使用第三方库 cobrautils 快速管理参数"
subtitle: "Cobra Third Party Library"
categories:
  - semi-plan 
tags:
  - semi-plan-202302-b
  - cobra08-cobrautils
date: "2023-02-17T22:34:14+08:00"
lastmod: "2023-02-17T22:34:14+08:00"
toc: true
draft: false
hiddenFromHomePage: false
pinTop: true

originAuthor: calacaly
originLink: "https://github.com/calacaly"
---

# Cobra08 - 使用第三方库 cobrautils 快速管理参数

## 提示

[`cobrautils`](https://github.com/go-jarvis/cobrautils) 是我写的一个 cobra 扩展库， 目的是为了能在 **结构体** 中 **方便** 的管理所有 flag 和默认值。

之前我们绑定参数时， 有多少参数就要写多少个 flag， 非常的麻烦。

```go
type Person struct {
    Name string
    Age int
}

var p=&Person{}

func init(){
    root.Flags.StringVarP(&p.Name, "name","n","zhangsan","user name")
    root.Flags.IntVarP(&p.Int, "age","a",20,"user age")
}
```

但是用了 `cobrautils` 后， 所有的定义都在 **结构体的标签中**， 就如同 `YAML` 和 `JSON` 管理参数。

## 作业要求

访问 https://github.com/go-jarvis/cobrautils ， 阅读 README.md 

实现以下效果。

```
go run .
student info

Usage:
  student [flags]

Flags:
      --addr.home string       home address (default "Sichuan,China")
      --addr.school string     school address
  -a, --age int                student age
      --contcat.email string   email address
      --contcat.phone string   phone number (default "138-8888-8888")
      --gender                 
  -h, --help                   help for student
      --name string            student name (default "zhangsan")
```

## 解题思路

第三方库的学习使用

+ 参数绑定依赖结构体和附加结构体标签

+ 结构体标签flag指定绑定的参数名称

+ usage指定绑定的绑参数说明

+ 参数默认值需要初始化结构体设置

+ 

```go
##配置形如示例的结构体，实现上面的效果需要嵌套使用结构体

type Student struct {
	Age     int     `flag:"age" usage:"student age"`
	Name    string  `flag:"name" usage:"student name"`
}

##初始化的值会绑定到参数的默认值中
student := &Student{Name: "zhangsan"}
cobrautils.BindFlags(rootCmd,student)
```
