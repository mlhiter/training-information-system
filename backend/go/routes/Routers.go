package routes

import (
	"backend/go/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	/*
		测试用 - 用户User路由组
	*/
	userGroup := r.Group("user")
	{
		//增加用户User
		userGroup.POST("/users", controller.CreateUser)
		//查看所有的User
		userGroup.GET("/users", controller.GetUserList)
	}
	/*
		权限认证
		前台界面
	*/
	frontEndGroup := r.Group("api")
	{
		// 权限认证
		// 登录
		frontEndGroup.POST("/login", controller.LoginUser)
		// 注册
		frontEndGroup.POST("/signup", controller.CreateUser)

		// 前台界面
		// 个人报名
		frontEndGroup.POST("/enroll/individual", controller.CreateRegister)
		// 团体报名
		frontEndGroup.POST("/enroll/organisation", controller.CreateTrainingApplication)
		// 获取缴费账单
		frontEndGroup.GET("/pay/individual", controller.GetPaymentRecordListByStudentName)
		// 签到
		frontEndGroup.POST("/inplace", controller.CreateCheckInRecordByStudentNameAndCourseName)
		// 问卷
		frontEndGroup.POST("/questionnaire", controller.CreateQuestionnaire)
	}
	/*
		后台界面 - 执行人界面
	*/
	// 报名
	enrollGroup := r.Group("api/enroll")
	{
		enrollGroup.GET("/list", controller.GetRegisterList)
		enrollGroup.PUT("/review", controller.UpdateRegister)
	}
	// 讲师
	lecturerGroup := r.Group("api/lecturer")
	{
		lecturerGroup.GET("/list", controller.GetLecturerList)
		lecturerGroup.POST("/add", controller.CreateLecturer)
		lecturerGroup.DELETE("/delete", controller.DeleteLecturer)
	}
	// 课程
	trainingGroup := r.Group("api/training")
	{
		trainingGroup.GET("/list", controller.GetCourseList)
		trainingGroup.POST("/add", controller.CreateCourse)
		trainingGroup.DELETE("/delete", controller.DeleteCourse)
	}
	// 学生
	studentGroup := r.Group("api/student")
	{
		studentGroup.GET("/list", controller.GetStudentList)
		studentGroup.POST("/add", controller.CreateStudent)
		studentGroup.DELETE("/delete", controller.DeleteStudent)
	}
	// 问卷
	questionnaireGroup := r.Group("api/questionnaire")
	{
		questionnaireGroup.GET("/list", controller.GetQuestionnaireList)
		questionnaireGroup.POST("/add", controller.CreateQuestionnaire)
		questionnaireGroup.DELETE("/delete", controller.DeleteQuestionnaire)
	}
	/*
		后台界面 - 经理界面
	*/
	// 执行人
	executorGroup := r.Group("api/executor")
	{
		executorGroup.GET("/list", controller.GetExecutorList)
		executorGroup.POST("/add", controller.CreateExecutor)
		executorGroup.DELETE("/delete", controller.DeleteExecutor)
	}
	// 课程收入
	incomeGroup := r.Group("api/income")
	{
		incomeGroup.GET("/list", controller.GetIncomeList)
	}
	// 培训审核
	enrollOrganisationGroup := r.Group("api/enroll/organisation")
	{
		enrollOrganisationGroup.GET("/list", controller.GetTrainingApplicationList)
		enrollOrganisationGroup.PUT("/review", controller.UpdateTrainingApplication)
	}

	/*
		后台界面 - 现场工作人员
	*/
	// 现场学员名单
	sceneGroup := r.Group("/api/scene")
	{
		sceneGroup.GET("/list", controller.GetStudentCourseListByCourseName)
	}
	return r
}
