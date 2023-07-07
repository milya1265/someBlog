package server

import "github.com/gin-gonic/gin"

type Server struct {
	Srv gin.Engine
}

func (s *Server) Start(port string) {
	s.Srv.Run(port)
}
