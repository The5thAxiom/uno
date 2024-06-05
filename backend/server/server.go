package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	Mux *http.ServeMux
	Port int
	Host string
} 

func (s Server) Run() error{
	s.AddRoutes()
	return http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.Mux)
}