package v1

import (
	v1 "github.com/bancodobrasil/featws-resolver-adapter-go/controllers/v1"
	"github.com/gin-gonic/gin"
)

// resolveRouter resolves a router group by adding a POST route to it.
func resolveRouter(router *gin.RouterGroup) {
	router.POST("/", v1.ResolveHandler)
}
