package usecase

import (
	"github.com/adityarizkyramadhan/neopath-versity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-versity-hackathon/repositories"
	"golang.org/x/crypto/bcrypt"
)

type StudentUsecase struct {
	repoGeneral *repositories.GeneralRepositoryImpl
}

func NewStudentUsecase(repoGeneral *repositories.GeneralRepositoryImpl) *StudentUsecase {
	return &StudentUsecase{repoGeneral: repoGeneral}
}

func (uu *StudentUsecase) Login(arg *models.StudentLogin) (string, error) {
	student := new(models.Student)

	if err := uu.repoGeneral.FindByColumn("username", arg.Email, student); err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(arg.Password)); err != nil {
		return "", err
	}

	return "", nil
}
