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


