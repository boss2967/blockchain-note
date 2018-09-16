# gitbook如何静态加载

**问题：**

通过gitbook build生成的 _book中的html文件，点击左侧的目录没有反应。

但是通过gitbook serve启动的是可以正常访问的

**原因**

从 3.0.0 版起, gitbook build 生成的 website 就不支持 local 打开了, 必需要 webserver 开启;

实在要完全静态的, 就安装 2.6.7 版吧( 在有些浏览器下估计不太完美 )

**解决方法一：**

```
gitbook build --gitbook=2.6.7
```

**解决方法二：**

在_book 下启动一个webserver


```
gitbook build
```
```
cd _book 
```
```
python -m SimpleHTTPServer 4001
```
<mark>大小写区分，端口自定义