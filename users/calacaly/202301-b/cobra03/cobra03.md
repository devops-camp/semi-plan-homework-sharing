
---
title: "作业: cobra - 03 交互式命令"
categories:
  - semi-plan 
tags:
  - semi-plan-202301-b 
  - cobra03

date: "2023-01-13T15:19:07+08:00"
lastmod: "2023-01-13T15:19:07+08:00"
toc: true
draft: false
hiddenFromHomePage: false


#  提交作业修改一下内容
pinTop: true
originAuthor: homework
originLink: ""
---


# 作业: cobra - 03 交互式命令

要求:

1.  使用 [https://github.com/spf13/cobra](https://github.com/spf13/cobra "https://github.com/spf13/cobra") 实现命令工具
2.  使用 [https://github.com/go-survey/survey](https://github.com/go-survey/survey "https://github.com/go-survey/survey") 实现交互式命令
3.  实现 Demo 效果

![](https://camo.githubusercontent.com/fd7c6e39ecf076e5da86a26ecfaa92d4e2ec3b56effe53bb804ed5cd7bb5f895/68747470733a2f2f7468756d62732e6766796361742e636f6d2f56696c6c61696e6f757347726163696f75734b6f75707265792d73697a655f726573747269637465642e676966)

### 解题思路

1.  学习survey
    > Input
    ```go
    name := ""
    prompt := &survey.Input{
        Message: "ping",
    }
    survey.AskOne(prompt, &name)
    ```
    > Select
    ```go
    color := ""
    prompt := &survey.Select{
        Message: "Choose a color:",
        Options: []string{"red", "blue", "green"},
    }
    survey.AskOne(prompt, &color)
    ```
2.  具体实现参考cmd/root.go
