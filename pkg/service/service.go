package service

import "L0/pkg/repository"

type Service struct {
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
