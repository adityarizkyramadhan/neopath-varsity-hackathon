package usecase

import (
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/repositories"
)

type MentorUsecase struct {
	repoGeneral *repositories.GeneralRepositoryImpl
}

func NewMentorUsecase(repoGeneral *repositories.GeneralRepositoryImpl) *MentorUsecase {
	return &MentorUsecase{repoGeneral: repoGeneral}
}

func (mu *MentorUsecase) GetAll() ([]*models.Mentor, error) {
	var mentors []*models.Mentor

	if err := mu.repoGeneral.FindAll(&mentors); err != nil {
		return mentors, err
	}

	return mentors, nil
}

func (mu *MentorUsecase) GetById(id uint) (*models.Mentor, error) {
	mentor := new(models.Mentor)
	if err := mu.repoGeneral.FindById(id, mentor); err != nil {
		return nil, err
	}
	return mentor, nil
}

func (mu *MentorUsecase) Register(arg models.MentorRegister) error {
	mentor := new(models.Mentor)
	mentor.Name = arg.Name
	mentor.Occupation = arg.Name
	mentor.Rating = arg.Rating
	return mu.repoGeneral.Create(mentor)
}
