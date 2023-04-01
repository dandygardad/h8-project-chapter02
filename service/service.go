package service

import "project08/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	BookService
}

func NewService(repo repository.RepoInterface) *Service {
	return &Service{repo: repo}
}
