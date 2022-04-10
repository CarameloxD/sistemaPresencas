package main

import (
	"sistemaPresencas/model"
	"sistemaPresencas/routes"
	"sistemaPresencas/services"

	"github.com/gin-gonic/gin"
	_ "gorm.io/driver/postgres"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var identityKey = "id"

func init() {
	services.OpenDatabase()
	services.Db.AutoMigrate(&model.Admin{})
	services.Db.AutoMigrate(&model.Student{})
	services.Db.AutoMigrate(&model.Course{})
	services.Db.AutoMigrate(&model.Teacher{})
	services.Db.AutoMigrate(&model.Subject{})
	services.Db.AutoMigrate(&model.Classroom{})
	services.Db.AutoMigrate(&model.Class{})
	services.Db.AutoMigrate(&model.Schedule{})
	services.Db.AutoMigrate(&model.Subscription{})
	services.Db.AutoMigrate(&model.Attendance{})
}

func main() {

	services.FormatSwagger()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	admin := router.Group("/api/v1/admin")
	{
		admin.POST("/login", routes.LoginAdmin)
	}

	student := router.Group("/api/v1/student")
	{
		student.GET("/:id", routes.GetStudentByNumber)
		student.GET("/", routes.GetAllStudents)
		student.POST("/", routes.InsertStudent)
		student.DELETE("/:id", routes.DeleteStudent)
		student.GET("getSchedulesByStudent/:id", routes.GetSchedulesByStudent)
	}

	teacher := router.Group("/api/v1/teacher")
	{
		teacher.GET("/:id", routes.GetTeacherInfo)
		teacher.GET("/", routes.GetAllTeachers)
		teacher.POST("/", routes.InsertTeacher)
	}

	subscription := router.Group("/api/v1/subscription")
	{
		subscription.POST("/", routes.InsertSubscription)
	}

	class := router.Group("/api/v1/class")
	{
		class.POST("/", routes.InsertClass)
		class.GET("/", routes.GetAllClasses)
	}

	classroom := router.Group("/api/v1/classroom")
	{
		classroom.POST("/", routes.InsertClassroom)
		classroom.GET("/", routes.GetAllClassrooms)
	}

	subject := router.Group("/api/v1/subject")
	{
		subject.POST("/", routes.InsertSubject)
		subject.GET("/", routes.GetAllSubjects)
	}

	course := router.Group("/api/v1/course")
	{
		course.GET("/", routes.GetAllCourses)
		course.POST("/", routes.InsertCourse)
	}

	schedule := router.Group("/api/v1/schedule")
	{
		schedule.POST("/", routes.InsertSchedule)
		schedule.GET("/", routes.GetAllSchedules)
		schedule.GET("/:id", routes.GetStudentsBySchedule)
	}

	attendance := router.Group("/api/v1/attendance")
	{
		attendance.POST("/", routes.InsertAttendance)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
