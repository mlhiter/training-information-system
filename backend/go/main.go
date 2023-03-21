package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-gin-gorm/go/dao"
	"go-gin-gorm/go/entity"
	"go-gin-gorm/go/routes"
)

func main() {
	//连接数据库
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库连接
	defer dao.Close()
	//绑定模型
	dao.SqlSession.AutoMigrate(&entity.User{})
	//注册路由
	r := routes.SetRouter()
	//启动端口为8081的项目
	r.Run(":8081")
}
