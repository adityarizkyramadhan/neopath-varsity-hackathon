package controllers

import (
	"net/http"

	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/usecase"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/utils"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	studentUsecase *usecase.StudentUsecase
}

func NewStudentController(studentUsecase *usecase.StudentUsecase) *StudentController {
	return &StudentController{studentUsecase: studentUsecase}
}

func (sc *StudentController) Login(c *gin.Context) {
	var input models.StudentLogin
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}
	token, err := sc.studentUsecase.Login(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success login", gin.H{"Token": token}))
}

func (sc *StudentController) Register(c *gin.Context) {
	var input models.StudentLogin
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}
	if err := sc.studentUsecase.Register(&input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success login", nil))
}
