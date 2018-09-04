# JavaScript嵌入页面的方式
## 行间事件（主要用于事件,不推荐使用）

`<input type="button" name="" onclick="alert('ok!')">`

## 页面script标签嵌入（内嵌式，部分推荐）

```
<script >
        alert('hello js inner!')
</script>
```

## 把Js代码写到新的文件里
新建一个.js文件，写入

```
<script >
        alert('hello js outer!')
</script>
```
在html中引用

```
<script src="js/hello.js">
</script>
```