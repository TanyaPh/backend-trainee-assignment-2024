package service

import "api/internal/repository"

type AuthService struct {
	repo repository.Admin
}

func newAuthService(repo repository.Admin) *AuthService {
	return &AuthService{repo: repo}
}