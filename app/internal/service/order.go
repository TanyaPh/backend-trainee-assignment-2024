package service

import "api/internal/repository"

type OrderService struct {
	repo     repository.Order
}

func newOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create() (int, error) {
	return s.repo.Create()
}

// func (s *OrderService) GetAll() (error) {
// 	return s.repo.GetAll()
// }

// func (s *OrderService) GetById() (error) {
// 	return s.repo.GetById()
// }

// func (s *OrderService) Update() error {
// 	return s.repo.Update()
// }

// func (s *OrderService) Delete() error {
// 	return s.repo.Delete()
// }
