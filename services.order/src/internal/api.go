package order

import (
	"github.com/aliaydins/oipattern/services.order/src/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "order service is up"})
}

func (h *Handler) createOrder(c *gin.Context) {
	req := new(entity.Order)
	c.BindJSON(req)

	err := h.service.CreateOrder(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order received successfully", "data": OrderMapper(req)})
}

func (h *Handler) getList(c *gin.Context) {
	outboxList, err := h.service.GetList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": outboxList})
}
