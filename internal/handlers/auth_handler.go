package handlers

import (
	"github.com/gin-gonic/gin"
	"jk.com/pets/internal/services"
	"jk.com/pets/pkg/response"
)

// 注册路由
func RegisterAuthRoutes(r *gin.Engine) {
	// 创建路由组，前缀为 /api/auth
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/register", register)
		authGroup.POST("/login", login)
		authGroup.GET("/get", getAllOwners)
	}
}

func register(c *gin.Context) {
	user := map[string]string{
		"username": "jiaokang",
		"email":    "@gmail.com",
	}
	response.Success(c, user)
}

func login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "login",
	})
}

func getAllOwners(c *gin.Context) {
	owners, err := services.GetAllOwners()
	if err != nil {
		response.Fail(c, 1000, err.Error())
		return
	}
	response.Success(c, owners)
}
