package usersHandlers

import (


	"users/usersModels"
	"users/usersServices"

	"github.com/gin-gonic/gin"
)

type IServices interface {
	PostRegisterHandler(c *gin.Context)
	

}

type handlers struct {
	s usersServices.IServices
}

func NewHandlers(s usersServices.IServices) IServices{
	return &handlers{s: s}
}

func (h *handlers) PostRegisterHandler(c *gin.Context) {
	var req usersModels.RequestRegister

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"status": "error", "message": err.Error()})
		return
	}
	err := h.s.PotsRegisterSer(req)
	if err != nil {
		c.JSON(500, gin.H{"status": "error", "message": "Failed to register user"})
		return
	}

	c.JSON(200, gin.H{"status": "OK", "message": "User registered successfully"})
}
