package service

import (
	"github.com/wmcff/zoogeek/container"
	"github.com/wmcff/zoogeek/model"
)

// ForestService is a service for managing master data such as format and forest.
type ForestService interface {
	FindAllForests() *[]model.Forest
}

type forestService struct {
	container container.Container
}

// NewForestService is constructor.
func NewForestService(container container.Container) ForestService {
	return &forestService{container: container}
}

// FindAllForests returns the list of all categories.
func (m *forestService) FindAllForests() *[]model.Forest {
	rep := m.container.GetRepository()
	forest := model.Forest{}
	result, err := forest.FindAll(rep)
	if err != nil {
		m.container.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}
