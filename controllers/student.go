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
	var input models.StudentRegister
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}
	if err := sc.studentUsecase.Register(&input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success add student", nil))
}

func (sc *StudentController) Update(c *gin.Context) {
	id := c.MustGet("id").(uint)
	var input models.StudentUpdate
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}
	if err := sc.studentUsecase.UpdateProfile(id, &input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success update student", nil))
}

func (sc *StudentController) Profile(c *gin.Context) {
	id := c.MustGet("id").(uint)
	student, err := sc.studentUsecase.GetProfile(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success get student", student))
}
