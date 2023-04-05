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

func CreateTrainingApplication(c *gin.Context) {
	var trainingApplication entity.TrainingApplication
	c.BindJSON(&trainingApplication)
	err := service.CreateTrainingApplication(&trainingApplication)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func GetTrainingApplicationList(c *gin.Context) {
	trainingApplicationList, err := service.GetAllTrainingApplication()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": trainingApplicationList,
		})
	}
}

func UpdateTrainingApplication(c *gin.Context) {
	var trainingApplication entity.TrainingApplication
	c.BindJSON(&trainingApplication)
	err := service.UpdateTrainingApplication(&trainingApplication)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func DeleteTrainingApplication(c *gin.Context) {
	var trainingApplication entity.TrainingApplication
	c.BindJSON(&trainingApplication)
	err := service.DeleteStudentById(trainingApplication.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}
