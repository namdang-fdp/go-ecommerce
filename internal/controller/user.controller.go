package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/namdang-fdp/go-ecommerce/internal/service"
	"github.com/namdang-fdp/go-ecommerce/pkg/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// controller --> service --> repo --> models --> dbs

func (uc *UserController) GetUserById(c *gin.Context) {
	response.SuccessResponse(c, 1000, []string{"namdang-fdp", "cr7"})
}
