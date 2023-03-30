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
	var student entity.Student
	c.BindJSON(&student)
	err := service.CreateStudent(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func GetStudentList(c *gin.Context) {
	studentList, err := service.GetAllStudent()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": studentList,
		})
	}
}

func DeleteStudent(c *gin.Context) {
	var student entity.Student
	c.BindJSON(&student)
	err := service.DeleteStudentById(student.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}
