# Gocore

## Quote

```
>config
github.com/Terry-Mao/goconf
>db
github.com/go-sql-driver/mysql

```

## used
### config 


```
> cat .env
[dev]
PORT 8033

[database]

DB_HOST 127.0.0.1:3306
DB_DATABASE db_name
DB_USERNAME root
DB_PASSWORD root

import "github.com/swirldawn/gocore"

	db_user := GetConfig("database", "DB_USERNAME")
```
### utils
```
import "github.com/swirldawn/gocore"
	//string md5 value
  	gocore.Md5("123456")
	//file md5 value
	gocore.FileMd5("/data/code/123456.jpg")

```

### db
```
import "github.com/swirldawn/gocore"
	//初始化mysql连接
	gocore.InitMysql()
	gocore.Insert("insert into user (id,name) values(1,'root')")
	gocore.Exec("delete from user where id = 100")
	gocore.FetchOne("select name from user where id = 1") //string "root"
	gocore.FetchRow("select * from user where id = 1") 
	gocore.FetchAll("select * from user limit 10") 
	//获取分页数据
	var params map[string]string /*创建集合 */
	params = make(map[string]string)
	params["page"] = "2"
	params["size"] = "3"
	params["id"] = ">= '3'"
	params["orderby"] = "order by id desc"
	
	articles := gocore.TablePaginator("articles", params)
	

```