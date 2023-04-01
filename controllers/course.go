package controllers

import (
	"net/http"

	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/usecase"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/utils"
	"github.com/gin-gonic/gin"
)

type CourseController struct {
	ucCourse *usecase.CourseUsecase
}

func NewCourseController(ucCourse *usecase.CourseUsecase) *CourseController {
	return &CourseController{ucCourse: ucCourse}
}

func (cc *CourseController) GetAllMeta(c *gin.Context) {
	role := c.Param("role")
	data, err := cc.ucCourse.GetAllMetaCourse(role)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success get all meta course", data))
}

func (cc *CourseController) GetAllData(c *gin.Context) {
	metaId := c.Param("meta_id")
	data, err := cc.ucCourse.GetAllDataCourse(metaId)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success get all data course", data))
}

func (cc *CourseController) DoneMeta(c *gin.Context) {
	metaId := c.Param("meta_id")
	studentID := c.MustGet("id").(uint)
	err := cc.ucCourse.MakeDone(metaId, studentID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success done meta course", nil))
}
