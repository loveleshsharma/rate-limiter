package handlers

import "github.com/gin-gonic/gin"

type Controller struct {
}

func (c Controller) RateLimit(ctx *gin.Context) {

}

func NewController() Controller {
	return Controller{}
}
