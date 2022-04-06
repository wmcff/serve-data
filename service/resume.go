package service

import (
	"github.com/wmcff/serve-data/container"
	"github.com/wmcff/serve-data/model"
)

// ResumeService is a service for managing master data such as format and resume.
type ResumeService interface {
	FindAllResumes() *[]model.Resume
}

type resumeService struct {
	container container.Container
}

// NewResumeService is constructor.
func NewResumeService(container container.Container) ResumeService {
	return &resumeService{container: container}
}

// FindAllResumes returns the list of all categories.
func (m *resumeService) FindAllResumes() *[]model.Resume {
	rep := m.container.GetRepository()
	resume := model.Resume{}
	result, err := resume.FindAll(rep)
	if err != nil {
		m.container.GetLogger().GetZapLogger().Errorf(err.Error())
		return nil
	}
	return result
}
