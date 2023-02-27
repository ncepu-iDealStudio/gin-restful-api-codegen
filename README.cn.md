# Go Gin框架接口项目代码生成器
###### By：华北电力大学iDeal工作室

## 项目介绍
本项目提供两种代码生成方式：

1. 根据提供的模板进行字符串替换用以生成项目
2. 根据数据库的库表结构和提供的代码模板进行代码的生成

这样使用的目的是为了便于大众使用自己的模板进行代码生成，而不是固定于一个模板。此方法生成代码不限语言，可以根据开发者的需要自己进行配置。

### 项目背景

现有的代码生成器大多是通过代码开发者的模板进行生成，无法自己进行代码的配置。我的代码生成器的完成工作是根据个人提供的代码模板进行代码的相关生成，提高工作效率节约时间。

## 项目使用说明
### 接口文档
https://console-docs.apipost.cn/preview/fcc88d2d841b6cd5/a11380b5bd0e7014

### 模板语法

1. 变量

   模板内内嵌的语法支持，全部需要加{{}}来标记。
   在模板文件内， . 代表了当前变量，即在非循环体内，.就代表了传入的那个变量
   假设我们传入一个字典：

   ```json
   {
       ArticleId: 1,
       ArticleContent: "这是一个示例内容"
   }
   ```

   那么我们在模板内可以通过

   ```html
   <p>{{.ArticleContent}}<span>{{.ArticleId}}</span></p>
   ```

   当然，我们有时候需要定义变量，比如我们需要定义一个article变量，同时将其初始化为”hello”，那么我们可以这样写：

   ```
   {{$article := "hello"}}
   ```

   假设我们想要把传入值的内容赋值给article，则可以这样写：

   ```
   {{$article := .ArticleContent}}
   ```

   这样我们只要使用{{$article}}则可以获取到这个变量的内容。

2. 判断

   golang的模板也支持if的条件判断，当前支持最简单的bool类型和字符串类型的判断

   ```
   {{if .condition}}
   {{end}}
   ```

   当.condition为bool类型的时候，则为true表示执行，当.condition为string类型的时候，则非空表示执行。

   当然也支持else ， else if嵌套

   ```
   {{if .condition1}}
   {{else if .contition2}}
   {{end}}
   ```

   假设我们需要逻辑判断，比如与或、大小不等于等判断的时候，我们需要一些内置的模板函数来做这些工作，目前常用的一些内置模板函数有：

   - not 非

     {{if not .condition}}
     {{end}}

   - and 与

     {{if and .condition1 .condition2}}
     {{end}}

   - or 或

     {{if or .condition1 .condition2}}
     {{end}}

   - eq 等于

     {{if eq .var1 .var2}}
     {{end}}

   - ne 不等于

     {{if ne .var1 .var2}}
     {{end}}

   - lt 小于 (less than)

     {{if lt .var1 .var2}}
     {{end}}

   - le 小于等于

     {{if le .var1 .var2}}
     {{end}}

   - gt 大于

     {{if gt .var1 .var2}}
     {{end}}

   - ge 大于等于

     {{if ge .var1 .var2}}
     {{end}}

3. 循环

   golang的template支持range循环来遍历map、slice内的内容，语法为：

   ```
   {{range $i, $v := .slice}}
   {{end}}
   ```

   在这个range循环内，我们可以通过iv来访问遍历的值，还有一种遍历方式为：

   ```
   {{range .slice}}
   {{end}}
   ```

   这种方式无法访问到index或者key的值，需要通过.来访问对应的value

   ```
   {{range .slice}}
   {{.field}}
   {{end}}
   ```

   当然这里使用了.来访问遍历的值，那么我们想要在其中访问外部的变量怎么办？(比如渲染模板传入的变量)，在这里，我们需要使用$.来访问外部的变量

   ```
   {{range .slice}}
   {{$.ArticleContent}}
   {{end}}
   ```

### 传入结构体

```go
type tablesCodeDict struct {
	TablesInfo []tableModel
	Dict       map[string]string
}

type tableModel struct {
	TableName   string
	StructName  string
	PackageName string
	Columns     []columnModel
	NaturalKey  []string
}

type columnModel struct {
	Field      string
	Type       string
	GoType     string
	Collation  *string
	Null       bool
	Key        string
	Default    *string
	Extra      string
	Privileges []string
	Comment    string
	NaturalKey bool
}
```

可以直接通过{{.}}的方式调用tablesCodeDict对象内容

## 项目思路

详见项目wiki

## 开发进度

- [x] 提供命令行工具的使用方式
  - [x] 生成项目代码功能
  - [x] 生成数据库代码功能
  - [x] 动态读取模板和配置文件
- [ ] 提供web的可视化生成工具
  - [x] 搭建Api后端框架
  - [x] 个人账号的登录等
  - [x] 项目管理
  - [x] 用户上传模板
  - [ ] 用户配置代码生成配置
  - [ ] 生成代码

# 项目设计思路交流
## cmd命令行方式设置

直接运行——按照config文件进行代码生成

-c 设置config文件名
- cmd: gen生成代码\
  args:
    - p——生成项目代码
    - d——生成数据库代码

## v1.0思路
1. 生成项目：
    1. 读取要生成的项目模板
    2. 读取要替换的字符串对
    3. 生成项目model
    4. 写入生成目标目录
2. 生成数据库结构：
    1. 读取配置文件，
        1. 获取数据库信息，
        2. 获取生成的代码层及模板
    2. 获取数据库信息，生成对应库表model
    3. 通过template生成代码
## V2.0思路
    # 代码生成器v2.0规划

## 一、界面（展示模块）

- [ ] 基于layui的纯web界面（基于embed封装进go项目中）
- [ ] 基于vue的前端界面

## 二、后端（核心模块）

### 1. 后端程序

1. 架构设计

    - [ ] 基于gin框架编写

    - [ ] 数据储存使用session

    - [ ] 数据持久化使用mysql

    - [ ] 用户登录/管理

2. 模块设计

    1. 用户登录
        1. SSO单点登录
        2. OAuth2登录
    2. 用户的项目代码模板
        1. 存储方式
            1. 本地存储
            2. 阿里云OSS对象存储
        2. 替换方式
            1. 字典替换
        3. 选配不生成文件
    3. 用户的数据库映射文件

### 2. 项目生成

后端读取项目文件思路：

1. 在上传时就读取
2. 在上传后通过定时任务将zip文件转化为一个大型的json文件
3. 在生成代码的时候解析zip文件形成json文件。

项目模板导入优化：

- [ ] 基于压缩包的代码导入
- [ ] 基于git的代码

项目生成规则：

- [ ] 复制项目模板
- [ ] 找到更好的替换字符串的方式
- [ ] 基于Key-Value字典进行替换
- [ ] 基于template代码生成

### 3. 数据库映射文件

1. 基于单表关系生成
    - [ ] 找出业务主键
    - [x] 基于template进行代码生成
2. 基于视图生成
    - [ ] 研究基于视图的增删改查方法
3. 基于多表关系生成
    - [ ] 外键
    - [ ] 相同名称
    - [ ] 相同备注

### 4. 代码模板

1. go-gin项目代码
    - [ ] 研究更好操作的update方法
    - [ ] 完善**增删改**的**存在/不存在**的判断方法
2. python项目代码
    - [ ] 移植ideal工作室的代码python-flask代码模板
    - [ ] 完善相关数据库代码生成

## 配置文件设计
1. 核心配置文件
    1.
2. 数据库信息配置
    1. 数据库类型
        1. host
        2. port
        3. username
        4. password
        5. database
3. 项目配置
    1. 项目名称
    2. 项目模板路径
    3. 项目生成目标路径
    4.
4. 根据数据库生成代码配置
    1. 