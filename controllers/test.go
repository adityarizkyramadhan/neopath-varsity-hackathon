package controllers

import (
	"net/http"
	"strconv"

	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/usecase"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/utils"
	"github.com/gin-gonic/gin"
)

type ApptitudeTestController struct {
	usecase *usecase.ApptitudeTestUsecase
}

func NewApptitudeTestController(usecase *usecase.ApptitudeTestUsecase) *ApptitudeTestController {
	return &ApptitudeTestController{
		usecase: usecase,
	}
}

func (ac *ApptitudeTestController) GetApptitudeTestBySection(c *gin.Context) {
	sectionStr := c.Param("section")
	section, err := strconv.Atoi(sectionStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error()))
		return
	}

	tests, err := ac.usecase.GetApptitudeTestBySection(section)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success", tests))
}

func (ac *ApptitudeTestController) SaveAnswer(c *gin.Context) {
	var answers []*models.ApptitudeAnswer
	if err := c.ShouldBindJSON(&answers); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error()))
		return
	}

	if err := ac.usecase.SaveAnswer(answers); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success", answers))
}
