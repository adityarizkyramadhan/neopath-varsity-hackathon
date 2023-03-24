package usecase

import (
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/middlewares"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/repositories"
	"golang.org/x/crypto/bcrypt"
)

type StudentUsecase struct {
	repoGeneral *repositories.GeneralRepositoryImpl
}

func NewStudentUsecase(repoGeneral *repositories.GeneralRepositoryImpl) *StudentUsecase {
	return &StudentUsecase{repoGeneral: repoGeneral}
}

func (su *StudentUsecase) Login(arg *models.StudentLogin) (string, error) {
	student := new(models.Student)

	if err := su.repoGeneral.FindByColumn("email", arg.Email, student); err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(arg.Password)); err != nil {
		return "", err
	}

	token, err := middlewares.GenerateJWToken(student.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (su *StudentUsecase) Register(arg *models.StudentLogin) error {
	student := new(models.Student)

	hashPass, err := bcrypt.GenerateFromPassword([]byte(arg.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	student.Email = arg.Email
	student.Password = string(hashPass)

	if err := su.repoGeneral.Create(student); err != nil {
		return err
	}

	return nil
}
