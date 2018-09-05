# jquery动画

通过animate方法可以设置元素某属性值上的动画，可以设置一个或多个属性值，动画执行完成后会执行一个函数。

```
/*
    animate参数：
    参数一：要改变的样式属性值，写成字典的形式
    参数二：动画持续的时间，单位为毫秒，一般不写单位
    参数三：动画曲线，默认为‘swing’，缓冲运动，还可以设置为‘linear’，匀速运动
    参数四：动画回调函数，动画完成后执行的匿名函数

*/

$('#div1').animate({
    width:300,
    height:300
},1000,'swing',function(){
    alert('done!');
});
```
## 练习 
动画：一个200*200的正方形，向右变大，向下移动，恢复原来大小。

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
            width:200px;
            height:200px;
            background:gold;
        }    
    </style>
    <script src="js/jquery-1.12.4.min.js"></script>
    <script>
        $(function(){
            var $btn = $('#btn');
            var $box = $('.box');
            
            $btn.click(function fnMove(){
                //$box.css({'width':1000});
                // $box.animate({'width':1000});
                /* animate的参数
                   参数一：要做动画的样式属性，写成字典形式
                   参数二：动画持续的时间，默认是400毫秒，设置时不写单位
                   参数三：动画曲线，默认是'swing',缓冲运动，还有'linear' 匀速运动
                   参数四：动画的回调函数，它是一个匿名函数，在动画结束时会自动调用
                */
                $box.animate({'width':1000},1000,'swing',function(){                    
                    $box.animate({'margin-top':400},1000,function(){                        
                        $box.animate({'width':200,'margin-top':0},1000,function(){
                            fnMove();
                        })
                    })
                });                
            })
        })    
    </script>
</head>
<body>
    <input type="button" value="动画" id="btn">
    <div class="box"></div>
</body>
</html>
```