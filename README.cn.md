# Go Gin框架接口项目代码生成器
###### By：华北电力大学iDeal工作室

## 项目介绍
本项目开发初衷是基于数据库的表，生成Restful接口来简化开发，去除编写CRUD基础代码的时间，提高软件开发效率。使用本项目可以快速生成你的项目，缩短交付时间，你可以有更多的时间花在业务逻辑上，更多的时间测试！我们致力于开发出开箱即用的代码生成工具，欢迎使用与反馈！
- 本项目基于已有的数据库表结构，自动生成Go语言完整的基础接口项目
- 生成的目标项目基于Gin框架，接口符合Restful风格规范
- 项目架构满足分层设计规范，分为实体层，控制器层和资源层(接口层)，用户可以添加服务层，作为商业逻辑层
- 提供完整的文件结构，利于项目管理，节约时间
- 工具类完整，包含了日志管理、错误处理等
- 提供RPC接口示例与文件组织，可以将项目转变为微服务项目
- 目标项目包含基于Docker容器的部署脚本

## 使用Go-GinGodeGen
查看我们的文档 []()

1. 下载我们的最新发行版，进行解压。项目结构如下

![项目目录结构.png](https://s1.ax1x.com/2023/02/27/pp9ILf1.png)
2. 进入configs文件夹，修改config.yaml文件，配置你的数据库
   ~~~~config.yaml:
   database:
     driver : "mysql"
     host : "XXXX"
     port : "XXXX"
     username : "XXX"
     password : "XXX"
     database : "XXX"
   ~~~~
3. 回到上一级目录，运行main.go文件
    `go run main.go`
4. 代码生成在dist文件夹下(第一次运行会在主目录生成dist文件夹)。dist文件夹存放生成的项目

    ![生成的项目.png](https://s1.ax1x.com/2023/02/27/pp9oqHg.png)

## 项目详细说明
### 一、目标项目说明
1. 生成的项目文件结构
    
以一个简单的项目study_flask_api为例，生成时项目名与数据库名一致。项目结构如下
~~~~
    study_flask_api
       ├─apis                           接口层
       │  ├─api_1_0                         1.0版本接口层
       │  │  ├─student                          接口-student表相关接口定义
       │  │  │  ├─StudentResource.go                Resource层-完成参数解析、dao层或服务层调用、数据响应
       │  │  │  └─urls.go                           student相关路由定义
       │  │  └─vStudentCourseScore              接口-vStudentCourseScore视图相关接口定义
       │  │     ├─StudentResource.go                Resource层-完成参数解析、dao层或服务层调用、数据响应
       │  │     └─urls.go                           vStudentCourseScore路由定义
       │  ├─middlewares                     中间件
       │  └─routers                         路由层-初始化路由、使用中间件
       ├─assets
       │  └─proto
       ├─cmd                                命令行定义-通过命令行运行项目
       ├─configs                            配置文件夹-系统配置与数据库配置
       ├─deploy                             部署相关
       ├─init                               项目配置初始化
       ├─internal                           内部层（业务相关）
       │  ├─dao                                 CRUD封装层
       │  ├─globals                             全局工具配置
       │  │  ├─codes
       │  │  ├─database
       │  │  ├─extensions
       │  │  │  └─currentUser
       │  │  ├─jwt
       │  │  ├─parser
       │  │  ├─rsa
       │  │  ├─snowflake
       │  │  └─vipers
       │  ├─models                          数据库model层-基于gorm的ORM模型
       │  ├─remote                          RPC Clinet示例
       │  │  ├─httpReq
       │  │  └─rpcReq
       │  ├─rpcServer                       RPC Server示例
       │  │  ├─middlewares                  
       │  │  ├─pb
       │  │  └─service
       │  ├─services                        业务逻辑层-复杂业务
       │  └─settings                        工具- 数据库连接、gin启动
       ├─pkg                                第三方包
       ├─scripts                            
       └─utils                              工具层
           ├─errHelper                          错误处理
           ├─loggers                            日志管理
           ├─message                            消息管理
           ├─rsa                                密钥管理
           ├─snowflake                          
           ├─str
           └─structs
~~~~
2. 开发说明
- 开发时在internal/services文件夹下完善业务逻辑，调用dao层对数据库进行操作
- 在api_1_0文件夹下进行参数接收，定义路由
- 在configs文件夹下定义了两个配置文件，可以根据实际需要修改配置文件
  - 系统配置文件：`config.yaml`
  - 数据库配置文件：`database.yaml`

3. 启动
    1. 在文件主目录下运行`go mod tidy` 按照依赖包
    2. 在文件主目录运行`go run main.go` 或者在文件主目录运行`go build` 然后使用命令行方式运行

4. 测试
    - 启动项目之后，打开接口测试工具（postman等），测试请求 http://127.0.0.1:8000/api/version 测试返回版本号
    - 查看代码中的路由定义，测试其他业务接口

### 二、生成器项目的使用建议
#### 一、数据库满足以下的设计规范（建议）
1. 数据库表名称推荐全小写，如student；如果涉及多个描述词，可使用"_"连接。如：user_info
2. 数据库表的字段中，必须包含一个主键；且不能为复合主键
3. 数据库表的名称和表字段名称，不能是Golang的关键字，如：`var`、`range`、`int`都是不正确的
4. 建议表的字段名称使用"大驼峰"命名法。如：UserName
5. 建议设计一个timestamp类型的"CreateTime"字段，默认为当前时间戳(用来记录数据创建的时间)
6. 建议设计一个tinyint类型的"IsDelete"字段(用来实现记录的逻辑删除，0--有效，1--已删除)，默认为0


### 三、模板使用说明
本项目基于模板文件，生成了对应的代码结构。使用者可以根据自己的要求，修改模板文件来满足特定的需求。在本项目assets/template文件夹下定义了相关模板
- 接口层模板 `api.go.tpl`
- dao层模板 `dao.go.tpl`
- 模型层模板 `model.go.tpl`
- 路由层模板 `router.go.tpl`
- 服务层模板 `service.go.tpl`

根据实际需求修改模板文件即可，模板语法见文档 [模板语法]()