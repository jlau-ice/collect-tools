package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jlau-ice/collect/internal/types/interfaces"
)

type UserHandler struct {
	userService interfaces.UserService
}

func NewUserHandler(
	userService interfaces.UserService,
) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	ctx := c.Request.Context()
	fmt.Print(ctx)
	c.JSON(http.StatusCreated, nil)
}
