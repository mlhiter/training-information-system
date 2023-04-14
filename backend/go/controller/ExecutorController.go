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

func CreateExecutor(c *gin.Context) {
	var executor entity.Executor
	c.BindJSON(&executor)
	err := service.CreateExecutor(&executor)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func GetExecutorList(c *gin.Context) {
	executorList, err := service.GetAllExecutor()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": executorList,
		})
	}
}

func DeleteExecutor(c *gin.Context) {
	var executor entity.Executor
	c.BindJSON(&executor)
	err := service.DeleteExecutorById(executor.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}
