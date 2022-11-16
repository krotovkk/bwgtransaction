package clienthandler

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/krotovkk/bwgtransaction/internal/core/ports"
)

type HttpHandler struct {
	clientService ports.ClientService
}

func NewHttpHandler(clientService ports.ClientService) *HttpHandler {
	return &HttpHandler{
		clientService: clientService,
	}
}

func (h *HttpHandler) GetBalance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	balance, err := h.clientService.GetBalance(id)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, struct {
		balance float64
	}{balance: balance})
}

func (h *HttpHandler) ChangeBalance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	diff, err := strconv.ParseFloat(c.Param("diff"), 64)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	client, err := h.clientService.ChangeBalance(id, diff)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(200, client)
}
