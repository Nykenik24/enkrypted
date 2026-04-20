package server

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/middleware"
)

func (s *Server) AddMiddleware(midware middleware.Middleware) error {
	name := midware.Info().Name
	if _, ok := s.midware[name]; ok {
		return fmt.Errorf("middleware '%s' already in server", name)
	}
	s.midware[name] = midware

	return nil
}

func (s *Server) RemoveMiddleware(midware middleware.Middleware) {
	delete(s.midware, midware.Info().Name)
}

func (s *Server) RemoveMidwareByName(name string) {
	delete(s.midware, name)
}
