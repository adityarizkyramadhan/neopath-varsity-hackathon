package usecase

import (
	"strconv"

	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/repositories"
)

type CourseUsecase struct {
	repoGeneral *repositories.GeneralRepositoryImpl
}

func NewCourseUsecase(repoGeneral *repositories.GeneralRepositoryImpl) *CourseUsecase {
	return &CourseUsecase{repoGeneral: repoGeneral}
}

func (cu *CourseUsecase) GetAllMetaCourse(role string) ([]*models.MetaLearningPath, error) {
	var data []*models.MetaLearningPath
	if err := cu.repoGeneral.DB.Where("role = ?", role).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (cu *CourseUsecase) GetAllDataCourse(metaId string) (*[]*models.DataLearningPath, error) {
	var data []*models.DataLearningPath
	metaIdInt, err := strconv.Atoi(metaId)
	if err != nil {
		return nil, err
	}
	if err := cu.repoGeneral.DB.Preload("MetaLP").Where("meta_id = ?", uint(metaIdInt)).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (cu *CourseUsecase) MakeDone(metaId string, studentId uint) error {
	return cu.repoGeneral.DB.
		Model(&models.StudentProgress{}).
		Update("is_done", true).
		Where("meta_id = ?", metaId).
		Where("student_id = ?", studentId).
		Error
}

func (cu *CourseUsecase) Create(data interface{}) error {
	return cu.repoGeneral.Create(data)
}
