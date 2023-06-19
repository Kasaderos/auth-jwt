package rest

import (
	"birthday-bot/internal/adapters/logger"
	"birthday-bot/internal/domain/usecases"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type St struct {
	lg      logger.Lite
	baseURL string
	ucs     *usecases.St
}

func GetHandler(lg logger.Lite, ucs *usecases.St, withCors bool, baseURL string) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// middlewares

	r.Use(MwRecovery(lg, nil))
	if withCors {
		r.Use(MwCors())
	}

	// handlers

	s := &St{lg: lg, ucs: ucs, baseURL: baseURL}

	// healthcheck
	r.GET("/healthcheck", func(c *gin.Context) { c.Status(http.StatusOK) })

	// users
	r.GET("/user/:id", s.hUserGet)
	r.POST("/user", s.hUserCreate)
	// TODO update password
	r.PUT("/user/:id", s.hUserUpdate)
	r.DELETE("/user/:id ", s.hUserDelete)

	r.POST("/auth/", s.hAuth)
	r.POST("/auth/by_refresh_token", s.hRefresh)
	r.GET("/profile", s.hProfile)
	r.POST("/logout", s.hLogout)
	r.POST("/register", s.hRegister)
	return r
}

func (o *St) getRequestContext(c *gin.Context) context.Context {
	return context.Background()
}
