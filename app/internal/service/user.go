package service

import "api/internal/repository"

type UserService struct {
	repo repository.User
}

func newUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) getBanner(tagId int, featureId int, lastRevision bool) {
	return s.repo.Get(tagId, featureId, lastRevision)
}
