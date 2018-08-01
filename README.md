# XSQL2

基于XSQL的深度优化，颠覆性操作习惯，支持了联表查询


第一步：在config_sql.json里面配置数据库属性

```json
{
	"sqlName":"root",//数据库用户名
	"sqlPassword":"root",//数据库密码
	"ip":"192.168.1.102",//数据库IP 
	"port":"3306",//数据库端口
	"Database":"Message",//数据库名称
	"path":"./"//生成文件所在目录
}
```

第二步：执行XSqlTools，生成与数据库匹配的Golang文件

第三步开始调用：

Select示例：

```go
package main

import (
   "fmt"
   "github.com/NiuStar/xsql2"
   "XSqlTools/Message"
)

func main() {

   var xs = xsql2.InitSql("root","root","192.168.1.102","3306","Message","utf8")
   message := Message.CreateMessages()//创建messages库的映射对象
   test := Message.CreateTest()//创建test库的映射对象
   abc := Message.Createabc()
   results := xs.Table(message,test).LeftJoin(abc).OnParam(abc.ID, "=", db2.ID).Field(test.ID.,message.ID.AS("MsgID"),abc.ID,test.NAME).Where(message.ID,"=","1").LL().OR().Where(message.ID,"=","1").RR().WhereParam( message.ID , "=" , test.ID).Select()

  // results := xs.Table(test).Field(test.ID,test.NAME,test.AGE).Where(test.ID,">","1").OrderByDESC(test.ID).OrderByASC(test.NAME).Select()
   xs.Table(test).Where(test.ID,">","1").Delete()

   for _,result := range results {
      fmt.Println(result)
   }
   return
}
```

Update示例：

```go
package main

import (
   "github.com/NiuStar/xsql2"
   "XSqlTools/Message"
)

func main() {

   var xs = xsql2.InitSql("root","root","192.168.1.102","3306","Message","utf8")
   message := Message.CreateMessages()//创建messages库的映射对象
   test := Message.CreateTest()//创建test库的映射对象

   xs.Table(message,test).Field(message.ID,message.IDD).SET(1,"??????").Where(message.ID,"=","1").WhereParam( message.ID , "=" , test.ID).Update()
   return
}
```

Delete示例：

```go
package main

import (
   "github.com/NiuStar/xsql2"
   "XSqlTools/Message"
)

func main() {
   var xs = xsql2.InitSql("root","root","192.168.1.102","3306","Message","utf8")
   test := Message.CreateTest()//创建test库的映射对象
   xs.Table(test).Where(test.ID,">","1").Delete()
   return
}
```

Insert示例：

```go
package main

import (
   "github.com/NiuStar/xsql2"
   "XSqlTools/Message"
)

func main() {
   var xs = xsql2.InitSql("root","root","192.168.1.102","3306","Message","utf8")
   test := Message.CreateTest()//创建test库的映射对象
   xs.Table(test).Field(test.NAME,test.AGE).Add("李海燕1","11").Add("李海燕2","12").Add("李海燕3","13").Add("李海燕4","14").Insert()
   return
}
```

