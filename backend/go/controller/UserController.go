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

func CreateUser(c *gin.Context) {
	//定义一个User变量
	var user entity.User
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&user)
	if u, _ := service.GetUserByUserName(user.UserName); u != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "userName重复",
		})
		return
	}
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	err := service.CreateUser(&user)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": user,
		})
	}
}

func LoginUser(c *gin.Context) {
	//定义一个User变量
	var user entity.User
	//将调用后端的request请求中的body数据根据json格式解析到User结构变量中
	c.BindJSON(&user)
	//将被转换的user变量传给service层的CreateUser方法，进行User的新建
	u, err := service.GetUserByUserName(user.UserName)
	//判断是否异常，无异常则返回包含200和更新数据的信息
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if u.Password == user.Password && u.Role == user.Role {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "success",
				"data": user,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "密码或职责错误",
			})
		}
	}
}

func GetUserList(c *gin.Context) {
	todoList, err := service.GetAllUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": todoList,
		})
	}
}
