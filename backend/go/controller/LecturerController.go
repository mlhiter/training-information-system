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

func CreateLecturer(c *gin.Context) {
	var lecturer entity.Lecturer
	c.BindJSON(&lecturer)
	err := service.CreateLecturer(&lecturer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func GetLecturerList(c *gin.Context) {
	lecturerList, err := service.GetAllLecturer()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": lecturerList,
		})
	}
}

func DeleteLecturer(c *gin.Context) {
	var lecturer entity.Lecturer
	c.BindJSON(&lecturer)
	err := service.DeleteLecturerById(lecturer.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}
