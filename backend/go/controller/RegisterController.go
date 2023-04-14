package controller

import (
	//需要用到的结构体
	"backend/go/entity"
	"fmt"
	"time"

	//gin框架的依赖
	"github.com/gin-gonic/gin"
	//http连接包
	"net/http"
	//service层方法
	"backend/go/service"
)

func CreateRegister(c *gin.Context) {
	// register
	var register entity.Register
	c.BindJSON(&register)
	err := service.CreateRegister(&register)

	student, _ := service.GetStudentByStudentName(register.Name)
	course, _ := service.GetCourseByCourseName(register.CourseName)

	fmt.Println("end register")

	// course
	course.SelectedNumber += 1
	service.UpdateCourse(course)

	fmt.Println("end course")

	// student_course
	var studentCourse entity.StudentCourse
	studentCourse.CourseId = course.ID
	studentCourse.CourseName = course.Name
	studentCourse.StudentId = student.ID
	studentCourse.StudentName = student.Name
	service.CreateStudentCourse(&studentCourse)

	fmt.Println("end sc")

	// payment_record
	var paymentRecord entity.PaymentRecord
	paymentRecord.CourseId = course.ID
	paymentRecord.CourseName = course.Name
	paymentRecord.StudentId = student.ID
	paymentRecord.StudentName = student.Name
	paymentRecord.PaymentAmount = course.Price
	paymentRecord.Payee = student.Name
	paymentRecord.PaymentTime = time.Now().Format("2006-01-02 15:04:05")
	paymentRecord.PayAccount = "6222031614003722522"
	service.CreatePaymentRecord(&paymentRecord)

	fmt.Println("end pr")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func GetRegisterList(c *gin.Context) {
	registerList, err := service.GetAllRegister()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
			"data": registerList,
		})
	}
}

func UpdateRegister(c *gin.Context) {
	var register entity.Register
	c.BindJSON(&register)
	err := service.UpdateRegister(&register)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}

func DeleteRegister(c *gin.Context) {
	var register entity.Register
	c.BindJSON(&register)
	err := service.DeleteStudentById(register.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "success",
		})
	}
}
