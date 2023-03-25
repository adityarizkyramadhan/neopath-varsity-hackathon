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

	r.GET("health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "API health 100%")
	})

	cfgDb, err := config.NewDatabase()
	if err != nil {
		panic(err.Error())
	}

	db, err := config.MakeConnectionDatabase(cfgDb,
		new(models.Student),
		new(models.School),
		new(models.Test),
		new(models.StudentSurveyScore),
		new(models.SurveyQuestion),
		new(models.SurveyResponse),
		new(models.Mentor),
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

	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
