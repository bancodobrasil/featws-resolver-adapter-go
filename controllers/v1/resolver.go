package v1

import (
	"net/http"

	"github.com/bancodobrasil/featws-resolver-adapter-go/services"
	payloads "github.com/bancodobrasil/featws-resolver-adapter-go/types"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// ResolveHandler godoc
// @Summary 		Resolve the JSON
// @Description 	Receive the params to create a resolver from JSON
// @Tags 			resolve
// @Accept  		json
// @Produce  		json
// @Param			context path string true "context"
// @Param 			load path string true "load"
// @Param  			parameters body payloads.ResolveInput true "Parameters"
// @Success 		200 {string} string "ok"
// @Failure 		400,404 {object} string
// @Failure 		500 {object} string
// @Failure 		default {object} string
// @Router 			/resolve [post]
func ResolveHandler(c *gin.Context) {
	var input payloads.ResolveInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Errorf("error occurs on biding JSON to input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := c.Request.Context()
	resolverOutput := services.Resolve(ctx, input)
	c.JSON(http.StatusOK, resolverOutput)
}
