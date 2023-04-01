package usecase

import (
	"strconv"

	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/repositories"
	"github.com/gin-gonic/gin"
)

type ReflectionUsecase struct {
	repoGeneral    *repositories.GeneralRepositoryImpl
	repoReflection *repositories.ReflectionRepository
}

func NewReflectionUsecase(repoGeneral *repositories.GeneralRepositoryImpl, repoReflection *repositories.ReflectionRepository) *ReflectionUsecase {
	return &ReflectionUsecase{repoGeneral: repoGeneral, repoReflection: repoReflection}
}

func (ru *ReflectionUsecase) GetQuestion(role string, section string, page string) (*models.Question, error) {
	secInt, err := strconv.Atoi(role)
	if err != nil {
		return nil, err
	}
	pageInt, err := strconv.Atoi(role)
	if err != nil {
		return nil, err
	}
	return ru.repoReflection.GetQuestion(role, secInt, pageInt)
}

func (ru *ReflectionUsecase) PostAnswer(arg []*models.AnswerInput, c *gin.Context) error {
	for _, v := range arg {
		err := ru.repoReflection.PostAnswer(c.MustGet("id").(uint), uint(v.QuestionID), v.Answer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ru *ReflectionUsecase) UpdateAnswer(arg []*models.AnswerInput, c *gin.Context) error {
	for _, v := range arg {
		err := ru.repoReflection.UpdateAnswer(c.MustGet("id").(uint), uint(v.QuestionID), v.Answer)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ru *ReflectionUsecase) GetEvaluation(c *gin.Context) (*models.Evaluation, error) {
	studentID := c.MustGet("id").(uint)
	questionID, err := strconv.Atoi(c.Param("question_id"))
	if err != nil {
		return nil, err
	}
	return ru.repoReflection.GetEvaluation(studentID, uint(questionID))
}
