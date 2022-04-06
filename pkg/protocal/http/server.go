package http

import "github.com/gin-gonic/gin"

type RegisterController func(r *gin.Engine)

// Server is a http server with gin embeded
type Server struct {
	router *gin.Engine
}

func NewServer(register RegisterController) *Server {
	router := gin.New()
	register(router)
	return &Server{router: router}
}


func (s *Server) Run() error{
	return s.router.Run(":8080")
}
