package memberships

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/panjiasmoroart/goforum26/internal/middleware"
	"github.com/panjiasmoroart/goforum26/internal/model/memberships"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine

	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("memberships")
	route.GET("/ping", h.Ping)
	route.POST("/signup", h.SignUp)
	route.POST("/login", h.Login)

	// route yg membutuhkan login
	routeRefresh := h.Group("memberships")
	routeRefresh.Use(middleware.AuthRefreshMiddleware())
	routeRefresh.POST("/refresh", h.Refresh)
}
