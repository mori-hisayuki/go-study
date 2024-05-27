package presenter

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mori-hisayuki/go-study/api/internal/presenter/controller/system"
)

const latest = "v1"

type Server struct{}

func (s *Server) Run(ctx context.Context) error {
	r := gin.Default()
	v1 := r.Group(latest)

	// 死活監視用
	{
		systemHandler := system.NewSystemHandler()
		v1.GET("/health", systemHandler.Health)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run()
	if err != nil {
		return err
	}
	return nil
}

func NewServer() *Server {
	return &Server{}
}
