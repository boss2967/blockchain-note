# JSX语法

jsx语法是一种类似于html标签的语法，它的作用相当于是让我们在JavaScript代码中直接写html代码，但是jsx不完全是html，它是 JavaScrip 的一种扩展语法，它具有 JavaScript 的全部能力，我们还可在jsx代码中插入变量或者表达式，用jsx语法写出来的语句是一个对象，我们可以将它存为一个变量，这个变量作为ReactDOM对象的render方法的第一个参数。

```
let el = <h1>Hello world!</h1>;
ReactDOM.render(
    el,
    document.getElementById('root')
)
```
jsx的结构还可以写得更复杂，可以是嵌套结构，如果是嵌套结构，需要有唯一的一个外层标签。标签中如果是单个的标签，在结尾要加“/”,在jsx中可以通过“{}”插入变量，表达式或者函数调用。

```

<script type="text/babel">
    let iNum01  = 10;
    let sTr = 'abc123456';
    let ok = true;
    let url = 'http://www.baidu.com;;
    function fnRev(s){
        return s.split('').reverse().join('');
    }    
    let el = (
        <div>
            <h3>jsx语法</h3>
            {/* 插入变量及运算 */}
            <p>{ iNum01+5 }</p>
            {/* 插入表达式 */}
            <p>{ sTr.split('').reverse().join('') }</p>
            {/* 插入函数调用 */}
            <p>{ fnRev(sTr) }</p>
            {/* 插入三元运算表达式 */}
            <p>{ ok?'YES':'NO' }</p>   
            <a herf={url} className="style01">百度</a>             
        </div>
    );

    ReactDOM.render(
        el,
        document.getElementById('root')
    )

</script>
```
jsx中指定标签的属性值建议用双引号，不能不用引号，属性名建议用驼峰式，其中class属性需要写成className，属性值如果是可变的，也可以写成“{}”的形式，里面可以和上面写法一样。 标签如果是单个的，在结尾一定要加“/”

```
{/* 定义class */}
<p className="sty01">使用样式</p>
{/* 单个标签，结尾要加“/” */}
<img src={user.avatarUrl} />
```