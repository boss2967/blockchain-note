# css选择器一
## 标签选择器

标签选择器，此种选择器影响范围大，一般用来做一些通用设置，或用在层级选择器中。

举例：

```
div{color:red} 
......
<div>这是第一个div</div>   <!-- 对应以上样式 -->
<div>这是第二个div</div>   <!-- 对应以上样式 -->
```
## 类选择器

通过类名来选择元素，一个类可应用于多个元素，一个元素上也可以使用多个类，应用灵活，可复用，是css中应用最多的一种选择器。

举例：

```
.blue{color:blue}
.big{font-size:20px}
.box{width:100px;height:100px;background:gold} 
......
<div class="blue">....</div>
<h3 class="blue big box">....</h3>
<p class="blue box">....</p>
```
## 层级选择器

主要应用在标签嵌套的结构中，层级选择器，是结合上面的两种选择器来写的选择器,它可与标签选择器结合使用，减少命名，同时也可以通过层级，限制样式的作用范围。

举例：

```
.con{width:300px;height:80px;background:green}
.con span{color:red}
.con .pink{color:pink}
.con .gold{color:gold}
......
<div class="con">
    <span>....</span>
    <a href="#" class="pink">....</a>
    <a href="#" class="gold">...</a>
</div>
<span>....</span>
<a href="#" class="pink">....</a>
```