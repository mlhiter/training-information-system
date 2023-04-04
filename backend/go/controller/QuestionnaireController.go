package controller

import (
	//需要用到的结构体
	"backend/go/entity"
	"backend/go/request"

	//gin框架的依赖
	"github.com/gin-gonic/gin"
	//http连接包
	"net/http"
	//service层方法
	"backend/go/service"
)

func CreateQuestionnaire(c *gin.Context) {
	var createQuestionnaireRequest request.CreateQuestionnaireRequest
	c.BindJSON(&createQuestionnaireRequest)
	err := service.CreateQuestionnaire(&createQuestionnaireRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func GetQuestionnaireList(c *gin.Context) {
	questionnaireList, err := service.GetAllQuestionnaire()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": questionnaireList,
		})
	}
}

func DeleteQuestionnaire(c *gin.Context) {
	var questionnaire entity.Questionnaire
	c.BindJSON(&questionnaire)
	err := service.DeleteQuestionnaireById(questionnaire.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}
