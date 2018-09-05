# jquery特殊效果
*	fadeIn() 淡入
*	fadeOut() 淡出
*	fadeToggle() 切换淡入淡出
*	hide() 隐藏元素
*	show() 显示元素
*	toggle() 切换元素的可见状态
*	slideDown() 向下展开
*	slideUp() 向上卷起
*	slideToggle() 依次展开或卷起某个元素

```
$btn.click(function(){

        $('#div1').fadeIn(1000,'swing',function(){
            alert('done!');
        });

    });
```

