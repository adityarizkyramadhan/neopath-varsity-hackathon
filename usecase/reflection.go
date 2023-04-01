package usecase

import (
	"strconv"

	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type ReflectionUsecase struct {
	repoGeneral    *repositories.GeneralRepositoryImpl
	repoReflection *repositories.ReflectionRepository
}

func NewReflectionUsecase(repoGeneral *repositories.GeneralRepositoryImpl, repoReflection *repositories.ReflectionRepository) *ReflectionUsecase {
	return &ReflectionUsecase{repoGeneral: repoGeneral, repoReflection: repoReflection}
}

func (ru *ReflectionUsecase) Create(arg *models.QuestionInput) error {
	question := new(models.Question)
	if err := copier.Copy(question, arg); err != nil {
		return err
	}
	return ru.repoGeneral.Create(question)
}

func (ru *ReflectionUsecase) GetQuestion(role string, section string, page string) (*models.Question, error) {
	secInt, err := strconv.Atoi(section)
	if err != nil {
		return nil, err
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	return ru.repoReflection.GetQuestion(role, secInt, pageInt)
}

func (ru *ReflectionUsecase) PostAnswer(arg []*models.AnswerInput, c *gin.Context) error {
	for _, v := range arg {
		err := ru.repoReflection.PostAnswer(c.MustGet("id").(uint), uint(v.QuestionID), v.Answer, v.SectionID)
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
	eval := new(models.Evaluation)
	for i := 1; i <= 5; i++ {
		avg, err := ru.repoReflection.GetEvaluation(studentID, i)
		avg *= 2.5
		if err != nil {
			return nil, err
		}
		if i == 1 {
			eval.Empathetic = avg
		} else if i == 2 {
			eval.Analytical = avg
		} else if i == 3 {
			eval.Adaptive = avg
		} else if i == 4 {
			eval.Collaborative = avg
		} else if i == 5 {
			eval.DesignSensitive = avg
		}
	}
	return eval, nil
}
