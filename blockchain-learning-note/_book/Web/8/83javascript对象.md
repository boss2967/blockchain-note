# javascript对象

javascript中的对象，可以理解成是一个键值对的集合，键是调用每个值的名称，值可以是基本变量，还可以是函数和对象。

创建javascript对象有两种方法，一种是通过顶级Object类来实例化一个对象，然后在对象上面添加属性和方法：

```
var person = new Object();

// 添加属性：
person.name = 'tom';
person.age = '25';

// 添加方法：
person.sayName = function(){
    alert(this.name);
}

// 调用属性和方法：
alert(person.age);
person.sayName();
```
还可以通过对象直接量的方式创建对象：

```
var person2 = {
    name:'Rose',
    age: 18,
    sayName:function(){
        alert('My name is' + this.name);
    }
}

// 调用属性和方法：
alert(person2.age);
person2.sayName();

```