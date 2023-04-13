package main

import (
	"backend/go/dao"
	"backend/go/entity"
	"backend/go/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//连接数据库
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	//程序退出关闭数据库连接
	defer dao.Close()
	//绑定模型
	dao.SqlSession.AutoMigrate(&entity.User{})
	dao.SqlSession.AutoMigrate(&entity.CheckInRecord{})
	dao.SqlSession.AutoMigrate(&entity.Course{})
	dao.SqlSession.AutoMigrate(&entity.Executor{})
	dao.SqlSession.AutoMigrate(&entity.Lecturer{})
	
	dao.SqlSession.AutoMigrate(&entity.PaymentRecord{})
	dao.SqlSession.AutoMigrate(&entity.Questionnaire{})
	dao.SqlSession.AutoMigrate(&entity.Register{})
	dao.SqlSession.AutoMigrate(&entity.Student{})
	dao.SqlSession.AutoMigrate(&entity.StudentCourse{})

	dao.SqlSession.AutoMigrate(&entity.TrainingApplication{})
	dao.SqlSession.AutoMigrate(&entity.TrainingNotice{})
	//注册路由
	r := routes.SetRouter()

	url := ginSwagger.URL("http://localhost:8081/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//启动端口为8081的项目
	r.Run(":8081")
}
