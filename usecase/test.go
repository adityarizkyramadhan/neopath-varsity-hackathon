package usecase

import (
	"fmt"

	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/models"
	"github.com/adityarizkyramadhan/neopath-varsity-hackathon/repositories"
)

type ApptitudeTestUsecase struct {
	repo *repositories.ApptitudeRepository
}

func NewApptitudeTestUsecase(repo *repositories.ApptitudeRepository) *ApptitudeTestUsecase {
	return &ApptitudeTestUsecase{
		repo: repo,
	}
}

// Get ApptitudeTest by Section
func (uc *ApptitudeTestUsecase) GetApptitudeTestBySection(section int) (*[]*models.ApptitudeTest, error) {
	return uc.repo.GetQuestionBySection(section)
}

// Save answer
func (uc *ApptitudeTestUsecase) SaveAnswer(answers []*models.ApptitudeAnswer) error {
	for _, answer := range answers {
		// Check if the corresponding test exists
		test, err := uc.repo.GetTestByID(answer.ApptitudeTestID)
		if err != nil {
			return err
		}
		if test == nil {
			return fmt.Errorf("test not found for answer with ID %d", answer.ID)
		}

		// Save the answer
		if err := uc.repo.CreateAnswer(answer); err != nil {
			return err
		}
	}
	return nil
}
