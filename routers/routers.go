package routers

import (
	"github.com/aifuxi/aifuxi_cool_api/middleware"
	"github.com/aifuxi/aifuxi_cool_api/settings"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.New()

	gin.SetMode(settings.AppConfig.Mode)

	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	initAdminAuthRouterGroup(r)
	initAdminPublicRouterGroup(r)

	return r
}
