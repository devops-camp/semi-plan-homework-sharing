---
title: "作业: 翻译 《Dockerfile 最佳实践》"
subtitle: "Dockerfile Best Practices"
categories:
  - semi-plan 
tags:
  - semi-plan-202301-b 
  - documents 
date: "2023-01-13T20:58:18+08:00"
lastmod: "2023-01-13T20:58:18+08:00"
toc: true
draft: false
hiddenFromHomePage: false

#  提交作业修改一下内容
pinTop: false
originAuthor: menah3m
originLink: "https://github.com/Menah3m"
---



# Dockerfile Best Practices

阅读并翻译 dockerfile 最佳实践

> https://docs.docker.com/develop/develop-images/dockerfile_best-practices/

1. 本文有很多命令， 从上下文可以推测文档内容。 因此相对简单。
2. 本文有很多关键字的知识点。

## 译文

这篇文章涵盖了高效地构建镜像的最佳实践和推荐方法。

Dockerfile 是一个按顺序包含了构建给定镜像所需的全部指令的文本文件，Docker 通过读取 Dockerfile 中的指令来自动构建镜像。
Dockerfile 遵循特定的格式和指令集，相关内容可以在 [Dockerfile 相关](https://docs.docker.com/engine/reference/builder/) 中找到。

每个 Docker 镜像都包含多个只读层，每一个只读层都对应 Dockerfile 中的一个指令。这些层堆叠在一起，每一层都是前一层的增量变化，以下是一个 Dockerfile 文件的例子：
```dockerfile
# syntax=docker/dockerfile:1
FROM ubuntu:18.04
COPY . /app
RUN make /app
CMD python /app/app.py
```
每条指令都会创建一层：
- `FROM` 使用 `ubuntu:18.04` Docker 镜像来创建一层
- `COPY` 从你的 Docker 客户端的当前目录添加文件
- `RUN` 使用 `make` 命令来构建你的应用
- `CMD` 指明容器内需要运行的命令

当你运行一个镜像来生成容器时，会产生一个新的可写层，这个可写层也被称为容器层，它位于所有下层的顶部。对正在运行的容器做出的所有改动，如写入新文件、编辑文件、删除文件等操作，都会写入这个可写层中。

关于更多镜像层和 Docker 如何构建和存储镜像的信息，可以参考[这里](https://docs.docker.com/storage/storagedriver/)

### 常用指南和推荐
#### 创建临时容器
通过 Dockerfile 定义的镜像应产生尽可能短暂的容器。短暂意味着容器可以被停止和清除，同时只需要极少的设置就可以重建和替换。
请参考 *12因素应用* 方法中的流程来理解使用这种无状态方式运行容器的动机。
#### 理解构建上下文
查看 [构建上下文](https://docs.docker.com/build/building/context/) 来了解更多
#### 通过标准输入编排 Dockerfile
Docker 拥有通过使用本地或远程标准输入编排的 Dockerfile 来构建镜像的能力。通过标准输入来编排 Dockerfile 对于不用将 Dockerfile 写入磁盘就可以构建一次性容器或者临时生成 Dockerfile 而后续不做持久化的场景是很有用的。
> 为了方便，下面的示例使用了[这里的文档](https://tldp.org/LDP/abs/html/here-docs.html) ,但通过标准输入提供 Dockerfile 的任何方法都是可以使用的。
> <br>比如，下面的命令是等价的：
> ```shell
> echo -e 'FROM busybox\nRUN echo "hello world"' | docker build -
> ```
> ```shell
> docker build -<<EOF
> FROM busybox
> RUN echo "hello world"
> EOF
> ```
> 你可以使用你喜欢的方式或最适合你使用情况的方式来替换示例。

#### 使用未传递构建上下文的标准输入创建的 Dockerfile 来构建镜像
这种语法可以使用一个没有传递额外文件作为构建上下文的标准输入来构建镜像。分隔符(`-`)处于`PATH`的位置,指导 Docker 从标准输入而不是目录去读取包含了 Dockerfile 的构建上下文。
```shell
docker build [OPTIONS] -
```
下面的示例使用通过标准输入传递的 Dockerfile 来构建镜像。可以看到，这里并没有文件作为构建上下文被传递给后台进程。
```shell
docker build -t myimage:latest -<<EOF
FROM busybox
RUN echo "hello world"
EOF
```
删除构建上下文在不需要复制文件到镜像的 Dockerfile 中是很有用的，因为并没有文件被发送到后台进程，所以它可以提高构建速度。<br>
如果你想通过从构建上下文中排除一些文件来提高构建速度，可以参考[这里](https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#exclude-with-dockerignore)
> 在使用不传递构建上下文的标准输入生成的 Dockerfile 来构建镜像时，如果使用了 `COPY` 和 `ADD` 指令，构建就会失败。<br>
> 如下所示：
> ```shell
> # create a directory to work in
> mkdir example
> cd example
>
> # create an example file
> touch somefile.txt
>
> docker build -t myimage:latest -<<EOF
> FROM busybox
> COPY somefile.txt ./
> RUN cat /somefile.txt
> EOF
>
> # observe that the build fails
> ...
> Step 2/3 : COPY somefile.txt ./
> COPY failed: stat /var/lib/docker/tmp/docker-builder249218248/somefile.txt: no such file or directory```
>


---
#### TODO
#### Build from a local build context, using a Dockerfile from stdin
