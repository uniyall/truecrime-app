package data

import "prateekkuniyal.dev/desicrimepod/models"

type MovieStorage interface {
	GetTopCases() ([]models.Case, error)
	// GetRandomCases() ([]models.Case, error)
	// GetCaseById(id int) (*models.Case, error)
	// SearchCasesbyName(name string) ([]models.Case, error)
	// GetAllCategories() ([]models.Category, error)
}
