---
title: "Cobra07 - 命令行自动补全 （简单）"
subtitle: "Cobra05 Completion"
categories:
  - 半月刊
tags:
  - 半月刊202302A
  - 简单
date: "2023-01-19T22:47:01+08:00"
lastmod: "2023-01-19T22:47:01+08:00"
toc: true
draft: false
hiddenFromHomePage: false
pinTop: true
---



# Cobra07 - 命令行自动补全 （简单）

## 提示

文档在官网

## 作业要求


## 作业要求

模拟 `kubectl` 实现一个带有 **多子命令** 及 **多参数** 的命令。 并根据官网提示， 在自己的命令终端实现 **补全** 功能

> 提示: 不用实现具体功能， 在 `Run()` 中打印一句话就行了。

![kubectl-completion](./kubectl-completion.gif)


## 解题思路


用cobra生成的命令行程序，默认有completion功能提示，查看提示

![completion](./render1675780961867.gif)

按提示执行后打开新的终端按执行命令，按tab键就可补全

![completion1](./render1675781011678.gif)