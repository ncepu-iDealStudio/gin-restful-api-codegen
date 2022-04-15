# 代码生成器
###### 作者：李瑞阳

## 项目介绍
本项目提供两种代码生成方式：

1. 根据提供的模板进行字符串替换用以生成项目
2. 根据数据库的库表结构和提供的代码模板进行代码的生成

这样使用的目的是为了便于大众使用自己的模板进行代码生成，而不是固定于一个模板。此方法生成代码不限语言，可以根据开发者的需要自己进行配置。

### 项目背景

现有的代码生成器大多是通过代码开发者的模板进行生成，无法自己进行代码的配置。我的代码生成器的完成工作是根据个人提供的代码模板进行代码的相关生成，提高工作效率节约时间。

## 项目使用说明

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