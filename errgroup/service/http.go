package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Service struct {
	c   chan struct{}
	srv *http.Server
}

func NewService(c chan struct{}) (s *Service, err error) {
	return &Service{
		c: c,
	}, nil
}

func (s *Service) Start() (err error) {
	/* start http server */
	g := gin.Default()
	g.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	g.POST("/api/v5/device/cmd/getrootca", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code_msg": "",
		})
	})
	g.POST("/shutdown", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code_msg": "shut down http server",
		})
		s.c <- struct{}{}
	})
	s.srv = &http.Server{
		Addr:    ":8000",
		Handler: g,
	}
	return s.srv.ListenAndServe()
}

func (s *Service) Stop() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
