package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router *MainRouter) AboutHandler(c *gin.Context) {
	fmt.Println("come here...")
	c.HTML(http.StatusOK, "about.html", nil)
}
