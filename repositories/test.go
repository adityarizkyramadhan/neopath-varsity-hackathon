package repositories

import (
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"gorm.io/gorm"
)

type ApptitudeRepository struct {
	DB *gorm.DB
}

func NewApptitudeRepository(db *gorm.DB) *ApptitudeRepository {
	return &ApptitudeRepository{
		DB: db,
	}
}

func (ar *ApptitudeRepository) CreateTest(test *models.ApptitudeTest) error {
	if err := ar.DB.Create(test).Error; err != nil {
		return err
	}
	return nil
}

func (ar *ApptitudeRepository) GetTestByID(id uint) (*models.ApptitudeTest, error) {
	test := new(models.ApptitudeTest)
	if err := ar.DB.Where("id = ?", id).First(test).Error; err != nil {
		return nil, err
	}
	return test, nil
}

func (ar *ApptitudeRepository) UpdateTest(test *models.ApptitudeTest) error {
	if err := ar.DB.Save(test).Error; err != nil {
		return err
	}
	return nil
}

func (ar *ApptitudeRepository) DeleteTest(id uint) error {
	if err := ar.DB.Where("id = ?", id).Delete(&models.ApptitudeTest{}).Error; err != nil {
		return err
	}
	return nil
}

func (ar *ApptitudeRepository) GetQuestionBySection(section int) (*[]*models.ApptitudeTest, error) {
	var data []*models.ApptitudeTest
	if err := ar.DB.Where("section = ?", section).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (ar *ApptitudeRepository) CreateAnswer(answer *models.ApptitudeAnswer) error {
	if err := ar.DB.Create(answer).Error; err != nil {
		return err
	}
	return nil
}

func (ar *ApptitudeRepository) GetAnswerByID(id uint) (*models.ApptitudeAnswer, error) {
	answer := new(models.ApptitudeAnswer)
	if err := ar.DB.Where("id = ?", id).First(answer).Error; err != nil {
		return nil, err
	}
	return answer, nil
}

func (ar *ApptitudeRepository) UpdateAnswer(answer *models.ApptitudeAnswer) error {
	if err := ar.DB.Save(answer).Error; err != nil {
		return err
	}
	return nil
}

func (ar *ApptitudeRepository) DeleteAnswer(id uint) error {
	if err := ar.DB.Where("id = ?", id).Delete(&models.ApptitudeAnswer{}).Error; err != nil {
		return err
	}
	return nil
}
