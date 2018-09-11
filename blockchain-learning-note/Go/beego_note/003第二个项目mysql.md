# beego操作mysql

1. 配置路由router.go

	```
	func init() {
		beego.Router("/mysql/create",&controllers.MysqlController{},"get:CreateMysql")
		beego.Router("/mysql/insert",&controllers.MysqlController{},"get:InsertMysql")
		beego.Router("/mysql/query",&controllers.MysqlController{},"get:QueryMysql")
	}
	```

2. 创建数据表 goMysql.go
	
	```
	//首先创建一个MysqlController继承自beego.Controller
	type MysqlController struct {
		beego.Controller
	}
	```
	
	```
	func (this *MysqlController)CreateMysql(){
		//1. 打开数据库连接
			//参数 1. 驱动名称
			//参数 2. 连接字符串
		conn,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
		if err != nil {
			beego.Info("sql open err:=",err)
			return
		}
		//3. 关闭连接
		defer conn.Close()
		//2. 执行语句
		res,err:=conn.Exec("create table userInfo(id int,name varchar(30))")
		fmt.Println(res)
		if err != nil {
			beego.Info("conn.Exec err:=",err)
			return
		}
		//4. 必须给页面响应数据，否则报错
		this.Ctx.WriteString("创建成功")
	}
	```
	

3. 插入数据
	
	```
	func (this *MysqlController)InsertMysql(){
		//1. 打开数据库连接
		//参数 1. 驱动名称
		//参数 2. 连接字符串
		conn,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
		if err != nil {
			fmt.Println("sql.Open",err)
			return
		}
		//3. 关闭连接
		defer conn.Close()
		//2. 执行语句
		_,err=conn.Exec("insert userInfo(id,name) values (?,?)",1,"zhangsan")
		if err != nil {
			beego.Info("插入失败")
			return
		}
		this.Ctx.WriteString("插入成功")
	}
	```

4. 查询数据
	
	```
	func (this *MysqlController)QueryMysql(){
		//1. 打开数据库连接
		//参数 1. 驱动名称
		//参数 2. 连接字符串
		conn,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
		if err != nil {
			beego.Info("sql open err:",err)
			return
		}
		//3. 关闭连接
		defer conn.Close()
		//2. 执行语句
		row,err:=conn.Query("select id from userInfo")
		if err != nil {
			beego.Info("conn.Query err:",err)
			return
		}
		var ids string
		//遍历查询结果
		for row.Next(){
			var id int
			row.Scan(&id)
			ids+= strconv.Itoa(id)+","
		}
		this.Ctx.WriteString("查询成功id:"+ids)
	}
	```
	
5. 删除数据
	
	```
	func (this *MysqlController)DeleteMysql(){
		conn,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
		if err != nil {
			beego.Info("sql open err:",err)
			return
		}
		defer conn.Close()
		//获取要删除的用户的id
		id:=this.GetString(":id")
		_,err=conn.Exec("delete from userInfo where id = ?",id)
		if err != nil {
			beego.Info("conn exec err:",err)
			return
		}
		this.Ctx.WriteString("删除成功")
	}
	```