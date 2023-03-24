package usecase

import (
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/middlewares"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/repositories"
	"github.com/jinzhu/copier"
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

func (su *StudentUsecase) Register(arg *models.StudentRegister) error {
	student := new(models.Student)

	hashPass, err := bcrypt.GenerateFromPassword([]byte(arg.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	student.Email = arg.Email
	student.Password = string(hashPass)
	student.SchoolID = arg.SchoolID

	if err := su.repoGeneral.Create(student); err != nil {
		return err
	}

	return nil
}

func (su *StudentUsecase) UpdateProfile(id uint, arg *models.StudentUpdate) error {
	student := new(models.Student)

	if err := su.repoGeneral.FindById(id, student); err != nil {
		return err
	}

	student.Name = arg.Name
	student.Age = arg.Age
	student.Gender = arg.Gender
	student.SelfDescription = arg.SelfDescription

	if err := su.repoGeneral.Update(id, student); err != nil {
		return err
	}

	return nil
}

func (su *StudentUsecase) GetProfile(id uint) (*models.StudentDTO, error) {
	student := new(models.Student)

	if err := su.repoGeneral.FindById(id, student); err != nil {
		return nil, err
	}

	studentDTO := new(models.StudentDTO)

	if err := copier.Copy(studentDTO, student); err != nil {
		return nil, err
	}

	return studentDTO, nil
}
