package server

import (
	"github.com/batariloa/StreamingService/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	streamHandler *handler.StreamHandler
}

func New(sh *handler.StreamHandler) *Server {
	return &Server{
		streamHandler: sh,
	}
}

func (s *Server) Start() {

	r := gin.Default()
	s.setupRoutes(r)
	r.Run()
}

func (s *Server) setupRoutes(router *gin.Engine) {

	v1 := router.Group("v1")
	{
		v1.GET("/stream/:songId",
			s.streamHandler.StreamFileToUserHandler)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
