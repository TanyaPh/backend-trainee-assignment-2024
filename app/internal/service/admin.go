package service

import "api/internal/repository"

type AdminService struct {
	repo repository.Auth
}

func newAdminService(repo repository.Auth) *AdminService {
	return &AdminService{repo: repo}
}
