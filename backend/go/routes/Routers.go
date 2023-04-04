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
		//修改某个User
		//userGroup.PUT("/users/:id",controller.UpdateUser)
		//删除某个User
		//userGroup.DELETE("/users/:id",controller.DeleteUserById)
	}
	/*
		权限认证
		前台界面
	*/
	frontEndGroup := r.Group("api")
	{
		// 权限认证
		// 登录
		frontEndGroup.POST("/login")
		// 注册
		frontEndGroup.POST("/signup")

		// 前台界面
		// 个人报名
		frontEndGroup.POST("/enroll/individual")
		// 团体报名
		frontEndGroup.POST("/enroll/organisation")
		// 获取缴费账单
		frontEndGroup.POST("/pay/individual")
		// 获取团体缴费账单
		frontEndGroup.POST("/pay/organisation")
		// 团体上传报名名单
		frontEndGroup.POST("/api/enroll/organisation/excel")
		// 签到
		frontEndGroup.POST("/inplace")
		// 问卷
		frontEndGroup.POST("/questionnaire")
	}
	/*
		后台界面 - 执行人界面
	*/
	// 报名
	enrollGroup := r.Group("api/enroll")
	{
		enrollGroup.GET("/list")
		enrollGroup.POST("/review")
	}
	// 讲师
	lecturerGroup := r.Group("api/lecturer")
	{
		lecturerGroup.GET("/list")
		lecturerGroup.POST("/add")
		lecturerGroup.DELETE("/delete")
	}
	// 讲师
	trainingGroup := r.Group("api/training")
	{
		trainingGroup.GET("/list")
		trainingGroup.POST("/add")
		trainingGroup.DELETE("/delete")
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
		executorGroup.GET("/list")
		executorGroup.POST("/add")
		executorGroup.DELETE("/delete")
	}
	// 课程收入
	incomeGroup := r.Group("api/income")
	{
		incomeGroup.GET("/list")
	}
	// 培训审核
	enrollOrganisationGroup := r.Group("api/enroll/organisation")
	{
		enrollOrganisationGroup.GET("/list")
		enrollOrganisationGroup.POST("/review")
	}

	/*
		后台界面 - 现场工作人员
	*/
	// 现场学员名单
	sceneGroup := r.Group("/api/scene")
	{
		sceneGroup.GET("/list")
	}
	return r
}
