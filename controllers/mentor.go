package controllers

import (
	"net/http"
	"strconv"

	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/usecase"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/utils"
	"github.com/gin-gonic/gin"
)

type MentorController struct {
	ucMentor *usecase.MentorUsecase
}

func NewMentorController(ucMentor *usecase.MentorUsecase) *MentorController {
	return &MentorController{ucMentor: ucMentor}
}

func (mc *MentorController) Register(c *gin.Context) {
	var input models.MentorRegister
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}
	if err := mc.ucMentor.Register(&input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success add mentor", nil))
}

func (mc *MentorController) GetById(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	mentor, err := mc.ucMentor.GetById(uint(idInt))
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success get by id mentor", mentor))
}

func (mc *MentorController) GetAll(c *gin.Context) {
	mentors, err := mc.ucMentor.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success get all mentor", mentors))
}
