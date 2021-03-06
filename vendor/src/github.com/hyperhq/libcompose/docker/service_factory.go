package docker

import (
	"github.com/hyperhq/libcompose/config"
	"github.com/hyperhq/libcompose/project"
)

// ServiceFactory is an implementation of project.ServiceFactory.
type ServiceFactory struct {
	context *Context
}

// Create creates a Service based on the specified project, name and service configuration.
func (s *ServiceFactory) Create(project *project.Project, name string, serviceConfig *config.ServiceConfig) (project.Service, error) {
	return NewService(name, serviceConfig, s.context), nil
}
