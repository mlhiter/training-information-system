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

func CreateTrainingNotice(c *gin.Context) {
	var trainingNotice entity.TrainingNotice
	c.BindJSON(&trainingNotice)
	err := service.CreateTrainingNotice(&trainingNotice)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func GetTrainingNoticeList(c *gin.Context) {
	trainingNoticeList, err := service.GetAllTrainingNotice()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": trainingNoticeList,
		})
	}
}

func DeleteTrainingNotice(c *gin.Context) {
	var trainingNotice entity.TrainingNotice
	c.BindJSON(&trainingNotice)
	err := service.DeleteStudentById(trainingNotice.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}
