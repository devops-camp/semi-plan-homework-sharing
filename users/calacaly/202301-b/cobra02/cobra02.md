# cobra02

title: "作业: cobra - 02 读取配置配置文件"
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
pinTop: true
originAuthor: calacaly
originLink: ""

***

# 作业: cobra - 02 读取配置配置文件

要求:

1.  使用 [https://github.com/spf13/cobra](https://github.com/spf13/cobra "https://github.com/spf13/cobra") 实现命令工具
2.  命令具有以下参数
    1.  `--config` , `-c` 配置文件

**配置文件如下**

```yaml
# config.yml
name: zhangsan
age: 20
```

1.  将配置文件保存为 `JSON` 格式

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

### 解题思路

1.  学习带有配置文件的参数绑定
    > Bind Flags with Config
    ```go
    //You can also bind your flags with viper:
    var author string

    func init() {
      rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
      viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
    }

    //In this example, the persistent flag author is bound with viper. Note: the variable author will not be set to the value from config, when the --author flag is provided by user.
    //使用viper和cobra进行参数持久绑定，可以关联手动配置和配置文件的参数，手动配置的参数将会覆盖配置文件配置的参数
    ```
    > viper库
    > viper支持多种配置文件格式读入和导出，非常好用的配置文件库，使用方法参阅文档
    <https://github.com/spf13/viper>
2.  将导出为json
    ```go
    viper.WriteConfigAs("config.json")
    ```
3.  具体实现参照代码cmd/root.go
