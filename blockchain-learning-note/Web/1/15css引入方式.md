# css引入方式
css引入页面的方式有三种：

1. 内联式：通过标签的style属性，在标签上直接写样式。
```
<div style="width:100px; height:100px; background:red ">......</div>
```
2. 嵌入式：通过style标签，在网页上创建嵌入的样式表。
```
<style type="text/css">
    div{ width:100px; height:100px; background:red }
    ......
</style>
```
3. 外链式：通过link标签，链接外部样式文件到页面中。
```
<link rel="stylesheet" type="text/css" href="css/main.css">
```