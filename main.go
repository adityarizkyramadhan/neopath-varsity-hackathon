package main

import (
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/config"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/controllers"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/repositories"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	cfgDb, err := config.NewDatabase()
	if err != nil {
		panic(err.Error())
	}

	db, err := config.MakeConnectionDatabase(cfgDb, new(models.Student))
	if err != nil {
		panic(err.Error())
	}

	repGeneral := repositories.NewGeneralRepositoryImpl(db)

	ucStudent := usecase.NewStudentUsecase(repGeneral)
	ctrlStudent := controllers.NewStudentController(ucStudent)
	routeStudent := r.Group("student")
	routeStudent.POST("login", ctrlStudent.Login)
	routeStudent.POST("register", ctrlStudent.Register)
}
