# Revel
一个高生产力的 Go 语言 Web 框架，中文官网：http://www.gorevel.cn/

官网手册：http://revel.github.io/manual/tool.html

***热编译***
编辑, 保存, 和 刷新时，Revel自动编译代码和模板，如果代码编译错误，会给出一个 错误提示，同时捕捉 运行期错误。
***全栈功能***
Revel 支持： 路由, 参数解析, 验证, session/flash, 模板, 缓存, 计划任务, 测试, 国际化 等功能。
***高性能***
Revel 基于 Go HTTP server 构建。
***同步***
Go HTTP server 对于每个请求都运行在 goroutine上。Write simple callback-free code without guilt。
***无状态***
Revel 提供了保持Web层无状态的可预知的扩展。例如，会话数据被存储在用户的cookie中，缓存由memcached集群提供支持。
***模块化***
Revel框架由被称为 过滤器 的中间件组成,它实现了几乎所有的请求处理功能。开发者可以自由地使用自定义的过滤器，比如如自定义的路由器，用以替换Revel默认的路由过滤器。

## 1.安装

切换到工作目录

安装 Revel 框架

go get github.com/revel/revel

安装 Revel 命令行工具

go get github.com/revel/cmd/revel

确保$ GOPATH / bin目录在PATH中

通过命令查看安装是否成功

revel version

## 2.命令

### 创建Revel工程
revel new [import_path] [skeleton]  

复制skeleton到import_path，skeleton可选参数，如果GOPATH有多个工作空间，Revel会根据GOPATH检查当前工作空间，并创建工程
### 运行工程
revel run [import_path] [run_mode] [port]

run_mode: dev 开发 prod 线上 ，开发模式支持热部署，prod通过修改conf/app.conf配置文件watch = true支持

port端口,默认9000
### 编译工程
revel build [import_path] [target_path] [run_mode]
### 打包
revel package [import_path] [run_mode]
### 删除编译文件
revel clean [import_path]

Deletes the app/tmp directory.
Deletes the app/routes directory.

### 测试
revel test [import_path] [run_mode] [suite.method]
Run all tests for the Revel app named by the given import path.

## 3.项目结构

gocode                  GOPATH 目录
  src                   GOPATH src 目录
    revel               Revel 安装目录
      ...
    sample              Revel应用程序根目录
      app               MVC目录
        controllers     控制器
          init.go       
        models          模型
        routes          
        views           模板
      tests             测试
      conf              
        app.conf        默认配置
        routes          路由定义
      messages          国际化
      public            静态文件
        css             CSS files
        js              Javascript files
        images          Image files
## 4.