# html标签入门

## 标签语
学习html语言就是学习标签的用法，html常用的标签有20多个，学会这些标签的使用，就基本上学会了HTML的使用。
## 标签的的使用方法

```
<!-- 1、成对出现的标签：-->

<h1>h1标题</h1>
<div>这是一个div标签</div>
<p>这个一个段落标签</p>


<!-- 2、单个出现的标签： -->
<br>
<img src="images/pic.jpg" alt="图片">

<!-- 3、带属性的标签，如src、alt 和 href等都是属性 -->
<img src="images/pic.jpg" alt="图片">
<a href="http://www.baidu.com">百度网</a>

<!-- 4、标签的嵌套 -->
<div>
    <img src="images/pic.jpg" alt="图片">
    <a href="http://www.baidu.com">百度网</a>
</div>
```

## 块元素标签(行元素)和内联元素标签(行内元素)

标签在页面上会显示成一个方块。除了显示成方块，它们一般分为下面两类：

*	**块元素**：在布局中默认会独占一行，宽度默认等于父级的宽度，块元素后的元素需换行排列。
*	**内联元素**：元素之间可以排列在一行，设置宽高无效，它的宽高由内容撑开，元素之间默认有小间距，而且是基线对齐(文字底部对齐)。

## 常用块元素标签

1. **标题标签，表示文档的标题，除了具有块元素基本特性外，还含有默认的外边距和字体大小**

	```
<h1>一级标题</h1>
<h2>二级标题</h2>
<h3>三级标题</h3>
<h4>四级标题</h4>
<h5>五级标题</h5>
<h6>六级标题</h6>
```

2. **段落标签，表示文档中的一个文字段落，除了具有块元素基本特性外，还含有默认的外边距**

	```
<p>本人叫张山，毕业于某大学计算机科学与技术专业，今年23岁，本人性格开朗、
稳重、待人真诚、热情。有较强的组织能力和团队协作精神，良好的沟通能力和社
交能力，善于处理各种人际关系。能迅速适应环境，并融入其中。</p>
<p>本人热爱研究技术，热爱编程，希望能在努力为企业服务的过程中实现自身价值。</p>
```

3. **通用块容器标签，表示文档中一块内容，具有块元素基本特性，没有其他默认样式**

	```
<div>这是一个div元素</div>
<div>这是第二个div元素</div>
<div>
    <h3>自我介绍</h3>
    <p>本人叫张山，毕业于某大学计算机科学与技术专业，今年23岁，本人性格开朗、
稳重、待人真诚、热情。有较强的组织能力和团队协作精神，良好的沟通能力和社
交能力，善于处理各种人际关系。能迅速适应环境，并融入其中。</p>
</div>
```

## 常用内联元素标签

1. **超链接标签，链接到另外一个网页，具有内联元素基本特性，默认文字蓝色，有下划线**

	```
<a href="02.html">第二个网页</a>
<a href="http://www.baidu.com">百度网</a>
<a href="http://www.baidu.com"><img src="images/logo.png" alt="logo"></a>
<a href="#">默认链接</a>
```
2.	**通用内联容器标签，具有内联元素基本特性，没有其他默认样式**

	```
<p>这是一个段落文字，段落文字中有<span>特殊标志或样式</span>的文字</p>
```
3.	**图片标签，在网页中插入图片，具有内联元素基本特性，但是它支持宽高设置**

	`<img src="images/pic.jpg" alt="图片" />`

## 其他常用功能标签

1. 换行标签
	`<p>这是一行文字，<br>这是一行文字</p>`
2. html注释

	html文档代码中可以插入注释，注释是对代码的说明和解释，注释的内容不会显示在页面上，html代码中插入注释的方法是：

	`<!-- 这是一段注释  -->`

## 常用html字符实体

代码中成段的文字，如果文字间想空多个空格，在代码中空多个空格，在渲染成网页时只会显示一个空格，如果想显示多个空格，可以使用空格的字符实体,代码如下：

```
<!--  在段落前想缩进两个文字的空格，使用空格的字符实体：&nbsp;   -->
<p>
&nbsp;&nbsp;一个html文件就是一个网页，html文件用编辑器打开显示的是文本，可以用<br />
文本的方式编辑它，如果用浏览器打开，浏览器会按照标签描述内容将文件<br />
渲染成网页，显示的网页可以从一个网页链接跳转到另外一个网页。</p>
```

在网页上显示 “<” 和 “>” 会误认为是标签，想在网页上显示“<”和“>”可以使用它们的字符实体，比如：

```
<!-- “<” 和 “>” 的字符实体为 &lt; 和 &gt;  -->
<p>
    &lt;div&gt;是一个html的一个标签<br>
    3 &lt; 5 <br>
    10 &gt; 5
</p>
```