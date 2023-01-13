package services

import "github.com/go/crud/models"

type MovieService interface {
	CreateMovie(*models.Movie) error
	GetMovie(*string) (*models.Movie, error)
	GetAll() ([]*models.Movie, error)
	UpdateMovie(*models.Movie) error
	DeleteMovie(*string) error
}
