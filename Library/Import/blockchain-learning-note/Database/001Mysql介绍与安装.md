# Mysql介绍与安装
## 了解MySQL
>	MySQL是目前最为流行的**<mark>开放源码**的数据库管理系统，是完全网络化的**<mark>跨平台**的**<mark>关系型数据库**系统，它是由瑞典的MySQL AB公司开发的，由MySQL的初始开发人员David Axmark和Michael "Monty" Widenius于1995年建立，目前属于Oracle公司。它的象征符号是一只名为Sakila的海豚，代表着MySQL数据库和团队的速度、能力、精确和优秀本质。 	MySQL数据库可以称得上是**<mark>目前运行速度最快**的SQL语言数据库。除了具有许多其他数据库所不具备的功能和选择之外，MySQL数据库还是一种完全免费的产品，用户可以直接从网上下载使用，而不必支付任何费用。

##		MySQL服务器的安装和配置

###	先下载
 ```
 sudo apt-get install mysql-server
 sudo apt install mysql-client
 sudo apt install libmysqlclient-dev
 ```


###	进入管理员模式
`su`(这一步会让你输入root密码，默认是没有密码，可以用`sudo passwd`设置密码)
输入密码进到root模式

###	进入mysql
输入`mysql`指令，进入mysql

###	设置mysql的root密码为123456
```
update user set authentication_string=PASSWORD("123456") where User='root';
```

###	退出mysql
`exit`

###	推出root模式
`exit`

###	登陆mysql
```
mysql -uroot -p123456
```



