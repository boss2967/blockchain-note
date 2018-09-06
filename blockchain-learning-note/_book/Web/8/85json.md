# json

>	json是 JavaScript Object Notation 的首字母缩写，单词的意思是javascript对象表示法，这里说的json指的是类似于javascript对象的字符串，它同时是一种数据格式，目前这种数据格式比较流行，逐渐替换掉了传统的xml数据格式。

json数据类似于JavaScript中的对象，但是它的键对应的值里面是没有函数方法的，值可以是普通变量，但不支持undefined，值还可以是数组或者JavaScript对象。

json写法需要注意的是，json中的属性名称和字符串值需要用双引号引起来，用单引号或者不用引号会导致读取数据错误。

json格式的数据：

```
{
    "name":"tom",
    "age":18
}
```
json的另外一个数据格式是数组，和javascript中的数组字面量相同。

`["tom",18,"programmer"]`

还可以是更复杂的数据机构：

```
{
    "name":"jack",
    "age":29,
    "hobby":["reading","travel","photography"]
    "school":{
        "name":"Merrimack College",
        "location":'North Andover, MA'
    }
}
```

json本质上是字符串，如果在js中操作json数据，可以将json字符串转化为JavaScript对象，转化的方式如下：

```
 	var sJson = '{"name":"tom","age":18}';
    var oPerson = JSON.parse(sJson);

    // 操作属性
    alert(oPerson.name);
    alert(oPerson.age);
```
