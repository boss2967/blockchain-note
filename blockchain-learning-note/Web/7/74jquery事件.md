# jquery事件
*	blur() 元素失去焦点
*	focus() 元素获得焦点
*	change() 当表单元素的值发生改变时
*	click() 鼠标单击
*	mouseover() 鼠标进入（进入子元素也触发）
*	mouseout() 鼠标离开（离开子元素也触发）
*	mouseenter() 鼠标进入（进入子元素不触发）
*	mouseleave() 鼠标离开（离开子元素不触发）
*	ready() DOM加载完成
*	submit() 用户递交表单


focus()

```
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
    <script src="js/jquery-1.12.4.min.js"></script>
    <script>
        $(function(){
            var $input = $('#input01');
            // focus一般用于让元素起始就获取焦点
            $input.focus();
            //绑定change事件
            $input.change(function(){
                console.log($(this).val())
            })
        })
    </script>
</head>

<body>
    <input type="text" id = "input01">
</body>
</html>
```

mouseenter(),mouseleave()

```
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
    <style>
        .box{
            width: 200px;
            height: 200px;
            background: gold;
            margin: 50px auto;
        }
        .box2{
            width: 100px;
            height: 100px;
            background: green;
        }
    </style>
    <script src="js/jquery-1.12.4.min.js"></script>
    <script>
        // $(document).ready(function(){

        // })
        $(function(){
            var $box = $('.box')
            // 移到子元素上也会触发
            // $box.mouseover(function(){
            //     $(this).animate({'margin-top':100})
            // })
            // $box.mouseout(function(){
            //     $(this).animate({'margin-top':50})
            // })
            // 如果不希望移到子元素上也触发
            $box.mouseenter(function(){
                $(this).animate({'margin-top':100})
            })
            $box.mouseleave(function(){
                $(this).animate({'margin-top':50})
            })
        })
    </script>
</head>
<body>
    <!-- .box自动生成div class = box -->
    <div class="box">
        <div class="box2"></div>
    </div>
</body>
</html>
```

**W3C上的例子**


```
<html>

<head>
    <script type="text/javascript" src="../js/jquery-1.12.4.min.js"></script>
    <script type="text/javascript">
        x = 0;
        y = 0;
        $(document).ready(function () {
            $("div.over").mouseover(function () {
                $(".over span").text(x += 1);
            });
            $("div.enter").mouseenter(function () {
                $(".enter span").text(y += 1);
            });
        });
    </script>
</head>

<body>
    <p>不论鼠标指针穿过被选元素或其子元素，都会触发 mouseover 事件。</p>
    <p>只有在鼠标指针穿过被选元素时，才会触发 mouseenter 事件。</p>
    <div class="over" style="background-color:lightgray;padding:20px;width:40%;float:left">
        <h2 style="background-color:white;">被触发的 Mouseover 事件：<span></span></h2>
    </div>
    <div class="enter" style="background-color:lightgray;padding:20px;width:40%;float:right">
        <h2 style="background-color:white;">被触发的 Mouseenter 事件：<span></span></h2>
    </div>
</body>

</html>
```
