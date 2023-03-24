package usecase

import "github.com/adityarizkyramadhan/neopath-varsity-hackathon/repositories"

type MentorUsecase struct {
	repoGeneral *repositories.GeneralRepositoryImpl
}

func NewMentorUsecase(repoGeneral *repositories.GeneralRepositoryImpl) *MentorUsecase {
	return &MentorUsecase{repoGeneral: repoGeneral}
}

