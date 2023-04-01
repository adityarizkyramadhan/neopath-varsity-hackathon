package repositories

import (
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"gorm.io/gorm"
)

type ReflectionRepository struct {
	DB *gorm.DB
}

func NewReflectionRepository(db *gorm.DB) *ReflectionRepository {
	return &ReflectionRepository{
		DB: db,
	}
}

func (rr *ReflectionRepository) GetQuestion(role string, section int, page int) (*models.Question, error) {
	question := new(models.Question)
	if err := rr.DB.Where("role = ?", role).Where("section = ?", section).Where("page = ?", page).First(question).Error; err != nil {
		return nil, err
	}
	return question, nil
}

func (rr *ReflectionRepository) PostAnswer(studentID uint, questionID uint, answer int, sectionID uint) error {
	ans := &models.Answer{
		StudentID:  studentID,
		QuestionID: questionID,
		Answer:     answer,
		SectionID:  sectionID,
	}
	if err := rr.DB.Create(ans).Error; err != nil {
		return err
	}
	return nil
}

func (rr *ReflectionRepository) UpdateAnswer(studentID uint, questionID uint, answer int) error {
	answerObj := new(models.Answer)
	if err := rr.DB.Where("student_id = ? AND question_id = ?", studentID, questionID).First(answerObj).Error; err != nil {
		return err
	}
	answerObj.Answer = answer
	if err := rr.DB.Save(answerObj).Error; err != nil {
		return err
	}

	return nil
}

func (rr *ReflectionRepository) GetEvaluation(studentID uint, section int) (float64, error) {

	var answer float64
	err := rr.DB.Model(&models.Answer{}).Where("student_id = ? AND section_id = ?", studentID, section).Select("AVG(answer)").Scan(&answer).Error
	if err != nil {
		return -1.0, err
	}
	return answer, nil
}
