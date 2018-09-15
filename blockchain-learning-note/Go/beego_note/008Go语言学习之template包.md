# Go语言学习之html/template包

>	官方介绍： 
Package template (html/template) implements data-driven templates for generating HTML output safe against code injection. It provides the same interface as package text/template and should be used instead of text/template whenever the output is HTML.

>	template包（html/template）实现了数据驱动的模板，用于生成可对抗代码注入的安全HTML输出。本包提供了和text/template包相同的接口，无论何时当输出是HTML的时候都应使用本包。

## 字段操作 
Go语言的模板通过{{}}来包含需要在渲染时被替换的字段，{{.}}表示当前的对象。

```
<title>{{ .Title }}</title>
```

## 输出嵌套字段内容

那么如果字段里面还有对象，我们可以使用{{with …}}…{{end}}和{{range …}}{{end}}来进行数据的输出。

```
{{ range array }}
    {{ . }}
{{ end }}
```
```
{{range $index, $element := array}}
    {{ $index }}
    {{ $element }}
{{ end }}
```

## 条件处理

在Go模板里面如果需要进行条件判断，那么我们可以使用和Go语言的if-else语法类似的方式来处理。 

*	if:

	```
	{{ if isset .Params "title" }}<h4>{{ index .Params "title" }}</h4>{{ end }}
	```
	```
	{{ if isset .Params1 .Params2 }}<h4>{{ index .Params "title" }}</h4>{{ end }}
	```
*	if …else:

	```
	{{ if isset .Params "alt" }}
    {{ index .Params "alt" }}
{{else}}
    {{ index .Params "caption" }}
{{ end }}
	```
*	and & or:

	```
	{{ if and (or (isset .Params "title") (isset .Params "caption")) (isset .Params "attr")}}
	```
*	with:
	
	```
	{{ with .Params.title }}<h4>{{ . }}</h4>{{ end }}
	```

# 支持pipe数据
	
```
{{. | html}}
```

```
{{ if isset .Params "caption" | or isset .Params "title" | or isset .Params "attr" }}
Stuff Here
{{ end }}
```

# 模板变量

```
{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
```
局部变量的作用域在end前。

# 模板函数 

## func Must

```
func Must(t *Template, err error) *Template
```
模板包里面有一个函数Must，它的作用是检测模板是否正确，例如大括号是否匹配，注释是否正确的关闭，变量是否正确的书写.

## func (*Template) Parse

```
func (t *Template) Parse(text string) (*Template, error)
```
进行解析 

Parse parses text as a template body for t. Named template definitions ({{define …}} or {{block …}} statements) in text define additional templates associated with t and are removed from the definition of t itself.

## func (*Template) Execute

```
func (t *Template) Execute(wr io.Writer, data interface{}) error
```
执行 

Execute applies a parsed template to the specified data object, writing the output to wr. If an error occurs executing the template or writing its output, execution stops, but partial results may already have been written to the output writer. A template may be executed safely in parallel.

## func (*Template) ParseFiles

```
func (t *Template) ParseFiles(filenames ...string) (*Template, error)
```
通过文件名指定模板:

```
t := template.New("hello")
template.Must(t.ParseFiles("hello.txt"))
template.Must(t.ParseFiles("world.txt"))
t.ExecuteTemplate(os.Stdout, "world.txt", nil)
```

# 简单应用

## 代码1：
```
package main

import (
    "os"
    "text/template"
    "fmt"
    )

type Person struct {
    Name string
    nonExportedAgeField string
}

func main() {
    p:= Person{Name: "Mary", nonExportedAgeField: "31"}

    t := template.New("nonexported template demo")
    t, _ = t.Parse("hello {{.Name}}! Age is {{.nonExportedAgeField}}.")
    err := t.Execute(os.Stdout, p)
    if err != nil {
        fmt.Println("There was an error"
    }
}
```

输出： 
`hello Mary! Age is There was an error`

是不是有些失望呢，这里想要提醒大家的是struct中字段的首字母大写，才可以导出！！！ 
只是把`nonExportedAgeField`改为`NonExportedAgeField`即可，得到正确输出： 
`hello Mary! Age is 31`.

## 代码2：

```
package main

import (
    "html/template"
    "log"
    "os"
)

func main() {
    const tpl = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>{{.Title}}</title>
    </head>
    <body>
        {{range .Items}}<div>{{ . }}</div>{{else}}<div><strong>no rows</strong></div>{{end}}
    </body>
</html>`

    check := func(err error) {
        if err != nil {
            log.Fatal(err)
        }
    }

    t, err := template.New("webpage").Parse(tpl)
    check(err)

    data := struct {
        Title string
        Items []string
    }{
        Title: "My page",
        Items: []string{
            "My pictures",
            "My dialog",
        },
    }
    err = t.Execute(os.Stdout, data)
    check(err)

    emptyItems := struct {
        Title string
        Items []string
    }{
        Title: "My white page",
        Items: []string{},
    }
    err = t.Execute(os.Stdout, emptyItems)
    check(err)

}
```

## 代码3： 

从文件读取html模板，并且通过服务器访问。 

首先写一个todos.html:

```
<h1>Todos</h1>
<ul>
    {{range .Todos}}
        {{if .Done}}
            <li><s>{{.Task}}</s></li>
        {{else}}
            <li>{{.Task}}</li>
        {{end}}
    {{end}}
</ul>
```

main.go:

```
package main

import (
    "html/template"
    "net/http"
)

type Todo struct {
    Task string
    Done bool
}

func main() {
    tmpl := template.Must(template.ParseFiles("todos.html"))
    todos := []Todo{
        {"Learn Go", true},
        {"Read Go Web Examples", true},
        {"Create a web app in Go", false},
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        tmpl.Execute(w, struct{ Todos []Todo }{todos})
    })

    http.ListenAndServe(":8080", nil)
}
```
