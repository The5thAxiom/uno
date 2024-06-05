package server

import (
	"fmt"
	"net/http"
)

func (s Server) AddRoutes() {
	s.Mux.HandleFunc("/ping", s.ping)
}

func (s *Server) ping(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusAccepted)
	fmt.Fprint(res, "{\"status\": \"running\"}")
}