# CSS权重

>	CSS权重指的是样式的优先级，有两条或多条样式作用于一个元素，权重高的那条样式对元素起作用,权重相同的，后写的样式会覆盖前面写的样式。

## 权重的等级

可以把样式的应用方式分为几个等级，按照等级来计算权重

1. 内联样式，如：style=””，权重值为1000
2. ID选择器，如：#content，权重值为100
3. 类，伪类，如：.content、:hover 权重值为10
4. 标签选择器，如：div、p 权重值为1 

## 权重的计算实例

实例一：

```
<style type="text/css">
    div{
        color:red;
    }        
</style>
......
<div style="color:blue">这是一个div元素</div>
<!-- 
两条样式同时作用一个div，上面的样式权重值为1，下面的行间样式的权重值为1000，
所以文字的最终颜色为blue 
-->
```

实例二：

```
<style type="text/css">

    body #content .main_content h2{
        color:red;    
    }

    #content .main_content h2{
        color:blue;
    }
</style>
......
<div id="content">
    <div class="main_content">
        <h2>这是一个h2标题</h2>
    </div>
</div>
<!-- 
第一条样式的权重计算： 1+100+10+1，结果为112；
第二条样式的权重计算： 100+10+1，结果为111；
h2标题的最终颜色为red
-->
```