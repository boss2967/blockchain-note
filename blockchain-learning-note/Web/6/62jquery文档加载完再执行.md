# jquery文档加载完再执行

将获取元素的语句写到页面头部，会因为元素还没有加载而出错，jquery提供了ready方法解决这个问题，它的速度比原生的 window.onload 更快。

```
<script type="text/javascript">

$(document).ready(function(){

     ......

});

</script>
```
可以简写为：

```
<script type="text/javascript">

$(function(){

     ......

});

</script>
```