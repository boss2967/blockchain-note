# 导入MySQL驱动

在Go中操作mysql数据库，需要引入第三方驱动，首先导入以下包名，如果不识别`_ "github.com/go-sql-driver/mysql"`，需要先在github上下载驱动，下载和配置详细过程可以参考附录。

```
import (
   "database/sql"
   _ "github.com/go-sql-driver/mysql"
)
```
# 连接数据库

通过调用sql.Open函数返回一个sql.DB指针; sql.Open函数原型如下:

`func Open(driverName, dataSourceName string) (*DB, error)`

driverName: 使用的驱动名. 这个名字其实就是数据库驱动注册到 database/sql 时所使用的名字.

dataSourceName: 数据库连接信息，这个连接包含了数据库的用户名, 密码, 数据库主机以及需要连接的数据库名等信息.

**sql.Open并不会立即建立一个数据库的网络连接, 也不会对数据库链接参数的合法性做检验, 它仅仅是初始化一个sql.DB对象. 当真正进行第一次数据库查询操作时, 此时才会真正建立网络连接;**

**语法：sql.Open(驱动名，"用户名:密码@[连接方式]（主机名称:端口号）/连接的数据库名")**

```
db,err :=sql.Open("mysql","root:1@tcp(192.168.110.3:3306)/huoying");

defer db.Close()
```

<mark>sql.DB表示操作数据库的抽象接口的对象，但不是所谓的数据库连接对象，sql.DB对象只有当需要使用时才会创建连接。</mark>

如果想立即验证连接，需要用Ping()方法;

```
err = db.Ping()
if err != nil{
   fmt.Println("ping fail")
   log.Fatal(err)
}else{
   fmt.Println("ping success")
}
```
DB的主要方法有：

1. **Query**执行数据库的Query操作，例如一个Select语句，返回*Rows
2. **QueryRow** 执行数据库至多返回1行的Query操作，返回*Row
3. **PrePare** 准备一个数据库query操作，返回一个*Stmt，用于后续query或执行。这个Stmt可以被多次执行，或者并发执行
4. **Exec** 执行数不返回任何rows的据库语句，例如delete操作

# 插入数据

## 单行插入

在go中输入代码：

```
ret, err := db.Exec("INSERT INTO hero VALUES(390000,'水影','上忍',null,'2000-10-10',22000,null,null)")

if err != nil {
   fmt.Printf("insert data error: %v\n", err)
   return
}
if RowsAffected, err := ret.RowsAffected(); nil == err {
   fmt.Println("影响的行数 -- RowsAffected:", RowsAffected)
}
```

可以通过然后结果ret中的RowsAffected()函数获取影响的行数

## 多行插入

```
str :=  `
INSERT INTO hero (
   heroId,
   heroName,
   job
)
VALUES
   (
      '111',
      'aaa',
      '下忍'
   ),
   (
      '222',
      'bbb',
      '下忍'
   );
`
ret, err := db.Exec(str)
if err != nil {
   fmt.Printf("insert data error: %v\n", err)
   return
}
if RowsAffected, err := ret.RowsAffected(); nil == err {
   fmt.Println("影响的行数 -- RowsAffected:", RowsAffected)
}

```

## 预编译插入

使用预编译插入可以后期指定插入的数据；

代码如下：

```
stmt, _ := db.Prepare(`INSERT INTO hero(heroId,heroName,job)VALUES(?,?,?);`)
ret, err := stmt.Exec(333,"ccc","上忍")

if err != nil {
   fmt.Printf("insert data error: %v\n", err)
   return
}
if RowsAffected, err := ret.RowsAffected(); nil == err {
   fmt.Println("影响的行数 -- RowsAffected:", RowsAffected)
}
```

# 查询数据 

## 单行查询

利用QueryRow实现单行查询，将插入结果放入到.Scan()的参数中

```
var heroId int
var heroName string

err =db.QueryRow("SELECT heroId,heroName FROM hero WHERE heroName = '漩涡鸣人';").Scan(&heroId,&heroName)

if err != nil{
   log.Fatal(err)
}

fmt.Println(heroId,heroName)
```

## 多行查询

```
rows, err := db.Query("SELECT heroId,heroName FROM hero;")

for rows.Next(){
   var heroId,heroName string
   err := rows.Scan(&heroId,&heroName);

   if err != nil{
      log.Fatal(err)
   }
   fmt.Println(heroId)
   fmt.Println(heroName)
}
```
## 预编译查询

```
var heroId,heroName string
stmt, _ := db.Prepare("SELECT heroId,heroName FROM hero where heroid = ?;");
ret, err := stmt.Query(90516);
//注意这里需要Next()下，不然下面取不到值
ret.Next();
ret.Scan(&heroId, &heroName);
fmt.Println("heroId =" , heroId, "heroName = ", heroName);
```
