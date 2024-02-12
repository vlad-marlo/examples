package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/vlad-marlo/example/go/modules/src/internal/model"
)

func (srv *Service) HandleGetByID(id string) (*model.StoredData, error) {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, storedItem := range srv.data {
		if storedItem.ID == id {
			return &storedItem, nil
		}
	}
	return nil, errors.New("not found")
}

func (srv *Service) HandleCreate(name string) model.StoredData {
	res := model.StoredData{
		ID:   uuid.NewString(),
		Name: name,
	}
	srv.mu.Lock()
	srv.data = append(srv.data, res)
	srv.mu.Unlock()
	return res
}

func (srv *Service) HandleGetAll() []model.StoredData {
	srv.mu.RLock()
	resp := make([]model.StoredData, len(srv.data), len(srv.data))
	copy(resp, srv.data)
	srv.mu.RUnlock()
	return resp
}
