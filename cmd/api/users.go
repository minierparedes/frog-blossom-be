package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type createUsersRequest struct {
	Username    string `json:"username" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Role        string `json:"role" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	UserUrl     string `json:"user_url"`
	Description string `json:"description"`
}

func (s *Server) createUsers(ctx *gin.Context) {
	var req createUsersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorRosponse(err))
		return
	}
}
