package service

import (
	"github.com/vlad-marlo/example/go/modules/src/internal/model"
	"sync"
)

type Service struct {
	data []model.StoredData
	mu   sync.RWMutex
}

func New() *Service {
	return &Service{
		data: make([]model.StoredData, 0, 10),
		mu:   sync.RWMutex{},
	}
}
