package controllers

import (
	"net/http"

	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/usecase"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/utils"
	"github.com/gin-gonic/gin"
)

type ReflectionController struct {
	ucReflection *usecase.ReflectionUsecase
}

func NewReflectionController(ucReflection *usecase.ReflectionUsecase) *ReflectionController {
	return &ReflectionController{ucReflection: ucReflection}
}

func (rc *ReflectionController) AnswerPost(c *gin.Context) {
	var input []*models.AnswerInput

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}

	if err := rc.ucReflection.PostAnswer(input, c); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success add answer", nil))

}

func (rc *ReflectionController) AnswerUpdate(c *gin.Context) {
	var input []*models.AnswerInput

	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}

	if err := rc.ucReflection.UpdateAnswer(input, c); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success add answer", nil))

}

func (rc *ReflectionController) GetQuestion(c *gin.Context) {
	role := c.Query("role")
	section := c.Query("section")
	page := c.Query("page")

	question, err := rc.ucReflection.GetQuestion(role, section, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success get question", question))
}

func (rc *ReflectionController) EvaluationGet(c *gin.Context) {
	evaluation, err := rc.ucReflection.GetEvaluation(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success get evaluation", evaluation))
}

func (rc *ReflectionController) Question(c *gin.Context) {
	evaluation, err := rc.ucReflection.GetQuestion(c.Query("role"), c.Query("section"), c.Query("page"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success get evaluation", evaluation))
}
