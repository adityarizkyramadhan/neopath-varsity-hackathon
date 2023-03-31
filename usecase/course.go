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

func (cu *CourseUsecase) GetAllMetaCourse(role string) (*[]*models.MetaLearningPath, error) {
	var data *[]*models.MetaLearningPath
	if err := cu.repoGeneral.FindByColumn("role", role, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (cu *CourseUsecase) GetAllDataCourse(metaId string) (*[]*models.MetaLearningPath, error) {
	var data *[]*models.MetaLearningPath
	metaIdInt, err := strconv.Atoi(metaId)
	if err != nil {
		return nil, err
	}
	if err := cu.repoGeneral.FindByColumn("meta_id", uint(metaIdInt), data); err != nil {
		return nil, err
	}
	return data, nil
}

func (cu *CourseUsecase) MakeDone(metaId string) error {
	return cu.repoGeneral.DB.Update("is_done", true).Error
}
