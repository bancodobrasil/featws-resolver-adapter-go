package v1

import (
	"net/http"

	"github.com/bancodobrasil/featws-resolver-adapter-go/services"
	"github.com/bancodobrasil/featws-resolver-adapter-go/types"
	"github.com/gin-gonic/gin"
)

// ResolveHandler ...
func ResolveHandler(c *gin.Context) {
	var input types.ResolveInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resolverOutput := services.Resolve(input)
	c.JSON(http.StatusOK, resolverOutput)
}
