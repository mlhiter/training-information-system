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

func CreatePaymentRecord(c *gin.Context) {
	var paymentRecord entity.PaymentRecord
	c.BindJSON(&paymentRecord)
	err := service.CreatePaymentRecord(&paymentRecord)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func GetPaymentRecordList(c *gin.Context) {
	paymentRecordList, err := service.GetAllPaymentRecord()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": paymentRecordList,
		})
	}
}

func UpdatePaymentRecord(c *gin.Context) {
	var paymentRecord entity.PaymentRecord
	c.BindJSON(&paymentRecord)
	err := service.UpdatePaymentRecord(&paymentRecord)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func DeletePaymentRecord(c *gin.Context) {
	var paymentRecord entity.PaymentRecord
	c.BindJSON(&paymentRecord)
	err := service.DeleteStudentById(paymentRecord.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}
