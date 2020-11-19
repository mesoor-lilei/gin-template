package router

import (
	"github.com/gin-gonic/gin"
	"gin-template/service/user"
)

func UserRouter(group *gin.RouterGroup) {
	group.POST("user", user.Save)
	group.GET("user/:id", user.GetById)
}
