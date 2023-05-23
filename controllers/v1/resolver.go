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
// @Description	Para conseguir utilizar o endpoint é necessário colocar no body, dentro do *context*, contexto, da requisição a agencia do cliente, *branch*, como também, a conta do mesmo, *account*, como o exemplo a seguir:
// @Description
// @Description	```
// @Description	{
// @Description	"context": {
// @Description		"account": "7894",
// @Description		"branch": "4024",
// @Description	},
// @Description	"load":[]
// @Description	}
// @Description	```
// @Description
// @Description	Com esse input o body de resposta trará todos os parâmetros da conta. O *load* é opcional na requisição. Nele é possível passar parâmetros que você deseja buscar no banco de dados ao invés de buscar todos os parâmetros, podendo buscar quantos sejam necessários, como o exemplo a seguir:
// @Description
// @Description	```
// @Description	{
// @Description	"context": {
// @Description		"account": "7894",
// @Description		"branch": "4024",
// @Description	},
// @Description	"load":["age","gender","holder"]
// @Description	}
// @Description	```
// @Description	Com esse input o body de resposta será trago dentro do parâmetro *context*, contexto.
// @Description
// @Description	```
// @Description	{
// @Description	"context": {
// @Description	"age": "36",
// @Description	"gender": "M",
// @Description	"holder": "1"
// @Description	},
// @Description	"errors": {}
// @Description	}
// @Description	```
// @Tags 			resolve
// @Accept  		json
// @Produce  		json
// @Param  			parameters body payloads.ResolveInput true "Parameters"
// @Success 		200 {string} string "ok"
// @Failure 		400,404 {object} string
// @Failure 		500 {object} string
// @Failure 		default {object} string
// @Security		Authentication Api Key
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
