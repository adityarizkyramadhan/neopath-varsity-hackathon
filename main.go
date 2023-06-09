package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/config"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/controllers"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/middlewares"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/repositories"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/usecase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}
	r := gin.New()

	r.Use(middlewares.CORS())

	r.Use(middlewares.TimeoutMiddleware())

	r.GET("health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "New Deployment v0.0.2")
	})

	cfgDb, err := config.NewDatabase()
	if err != nil {
		panic(err.Error())
	}

	db, err := config.MakeConnectionDatabase(
		cfgDb,
		new(models.Student),
		new(models.School),
		new(models.ApptitudeTest),
		new(models.Mentor),
		new(models.MetaLearningPath),
		new(models.DataLearningPath),
		new(models.StudentProgress),
		new(models.Question),
		new(models.Answer),
	)
	if err != nil {
		panic(err.Error())
	}

	repGeneral := repositories.NewGeneralRepositoryImpl(db)

	ucStudent := usecase.NewStudentUsecase(repGeneral)
	ctrlStudent := controllers.NewStudentController(ucStudent)
	routeStudent := r.Group("student")
	routeStudent.POST("login", ctrlStudent.Login)
	routeStudent.POST("register", ctrlStudent.Register)
	routeStudent.PUT("profile", middlewares.ValidateJWToken(), ctrlStudent.Update)
	routeStudent.GET("profile", middlewares.ValidateJWToken(), ctrlStudent.Profile)

	ucMentor := usecase.NewMentorUsecase(repGeneral)
	ctrlMentor := controllers.NewMentorController(ucMentor)
	routeMentor := r.Group("mentor")
	routeMentor.POST("register", ctrlMentor.Register)
	routeMentor.GET("details/:id", middlewares.ValidateJWToken(), ctrlMentor.GetById)
	routeMentor.GET("all", middlewares.ValidateJWToken(), ctrlMentor.GetAll)

	ucCourse := usecase.NewCourseUsecase(repGeneral)
	ctrlCourse := controllers.NewCourseController(ucCourse)
	routeCourse := r.Group("course")
	routeCourse.GET("meta", middlewares.ValidateJWToken(), ctrlCourse.GetAllMeta)
	routeCourse.GET("data/:meta_id", middlewares.ValidateJWToken(), ctrlCourse.GetAllData)
	routeCourse.PUT("meta/:meta_id", middlewares.ValidateJWToken(), ctrlCourse.DoneMeta)
	//Dummy
	routeCourse.POST("meta", ctrlCourse.CreateMeta)
	routeCourse.POST("data", ctrlCourse.CreateData)

	repoReflection := repositories.NewReflectionRepository(db)
	ucReflection := usecase.NewReflectionUsecase(repGeneral, repoReflection)
	ctrlReflection := controllers.NewReflectionController(ucReflection)
	routeReflection := r.Group("reflection")
	routeReflection.POST("answer", middlewares.ValidateJWToken(), ctrlReflection.AnswerPost)
	routeReflection.PUT("answer", middlewares.ValidateJWToken(), ctrlReflection.AnswerUpdate)
	routeReflection.GET("question", middlewares.ValidateJWToken(), ctrlReflection.Question)
	routeReflection.GET("evaluation", middlewares.ValidateJWToken(), ctrlReflection.EvaluationGet)
	//Dummy
	routeReflection.POST("question", ctrlReflection.CreateQuestion)

	repoApptitude := repositories.NewApptitudeRepository(db)
	ucApptitude := usecase.NewApptitudeTestUsecase(repoApptitude)
	ctrlApptitude := controllers.NewApptitudeTestController(ucApptitude)
	routeApptitude := r.Group("apptitude")
	routeApptitude.POST("answer", middlewares.ValidateJWToken(), ctrlApptitude.SaveAnswer)
	routeApptitude.GET("test/:section", middlewares.ValidateJWToken(), ctrlApptitude.GetApptitudeTestBySection)

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
