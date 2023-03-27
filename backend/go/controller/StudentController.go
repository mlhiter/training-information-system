package controller

import (
	//需要用到的结构体
	"backend/go/entity"
	//gin框架的依赖
	"github.com/gin-gonic/gin"
	//http连接包
	"net/http"
	//service层方法
	"backend/go/service"
)

func CreateStudent(c *gin.Context) {
	//定义一个User变量
	var student entity.Student
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&student)
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err := service.CreateStudent(&student)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": student,
		})
	}
}
