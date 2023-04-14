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

func CreateStudentCourse(c *gin.Context) {
	var studentCourse entity.StudentCourse
	c.BindJSON(&studentCourse)
	err := service.CreateStudentCourse(&studentCourse)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func GetStudentCourseList(c *gin.Context) {
	studentCourseList, err := service.GetAllStudentCourse()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": studentCourseList,
		})
	}
}

func GetStudentCourseListByCourseName(c *gin.Context) {
	var getStudentCourseRequest request.GetStudentCourseRequest
	c.BindJSON(&getStudentCourseRequest)
	course, _ := service.GetCourseByCourseName(getStudentCourseRequest.CourseName)
	studentCourseList, err := service.GetStudentCourseListByCourseId(course.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": studentCourseList,
		})
	}
}

func DeleteStudentCourse(c *gin.Context) {
	var studentCourse entity.StudentCourse
	c.BindJSON(&studentCourse)
	err := service.DeleteStudentById(studentCourse.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}
