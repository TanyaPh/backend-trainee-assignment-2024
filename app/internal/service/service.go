package service

import "api/internal/repository"

type Auth interface {
	
}

type Admin interface {
	
}

type User interface {
	getBanner(tagId, featureId, lastRevision)
}

type Service struct {
	Authorization Auth
	Admin Admin
	User User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repo),
		Admin: newAdminService(repo),
		User: newUserService(repo),
	}
}