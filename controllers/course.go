package controllers

import (
	"net/http"

	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/usecase"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type CourseController struct {
	ucCourse *usecase.CourseUsecase
}

func NewCourseController(ucCourse *usecase.CourseUsecase) *CourseController {
	return &CourseController{ucCourse: ucCourse}
}

func (cc *CourseController) GetAllMeta(c *gin.Context) {
	role := c.Query("role")
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

func (cc *CourseController) CreateMeta(c *gin.Context) {
	input := new(models.MetaLearningPathInput)
	if err := c.Bind(input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}
	meta := new(models.MetaLearningPath)
	if err := copier.Copy(meta, input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	if err := cc.ucCourse.Create(meta); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success add meta course", nil))
}

func (cc *CourseController) CreateData(c *gin.Context) {
	input := new(models.DataLearningPathInput)
	if err := c.Bind(input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, utils.ResponseWhenFail(err.Error()))
		return
	}
	data := new(models.DataLearningPath)
	if err := copier.Copy(data, input); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	if err := cc.ucCourse.Create(data); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseWhenFail(err.Error()))
		return
	}
	c.JSON(http.StatusOK, utils.ResponseWhenSuccess("success add data course", nil))
}
