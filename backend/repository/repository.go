package repository

import (
	"awesomeProject/models"
)

type RequestRepo interface {
	CreateRequest(req *models.Request) error
	GetAllRequests() ([]*models.Request, error)
	DeleteRequest(req *models.Request) error
	EditRequest(req *models.Request) error
}