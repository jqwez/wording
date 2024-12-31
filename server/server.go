package server

import (
	"net/http"

	"github.com/jqwez/wording/server/routes"
)

type Server struct {
	Mux *http.ServeMux
}

func NewServer() *Server {
	server := &Server{}
	server.Mux = http.NewServeMux()
	server.RegisterStatic()
	_ = routes.NewBlossomRoutes(server.Mux)

	return server
}

func (s *Server) RegisterStatic() {
	s.Mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
}
