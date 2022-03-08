package v1

import (
	v1 "github.com/bancodobrasil/featws-resolver-adapter-go/controllers/v1"
	"github.com/gin-gonic/gin"
)

func resolveRouter(router *gin.RouterGroup) {
	router.POST("/", v1.ResolveHandler)
}
