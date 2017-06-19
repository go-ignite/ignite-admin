package controllers

import (
	"ignite/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

func (router *MainRouter) PanelStatusHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "status.html", nil)
}

func (router *MainRouter) PanelIndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (router *MainRouter) PanelLoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	pwd := c.PostForm("password")

	if username == authPass && pwd == authPass {
		// Create the token
		token := jwt.New(jwt.GetSigningMethod("HS256"))
		// Set some claims
		token.Claims = jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		}
		// Sign and get the complete encoded token as a string
		tokenString, err := token.SignedString([]byte(secret))
		resp := models.Response{}

		if err != nil {
			resp.Success = false
			resp.Message = "Could not generate token"
			c.JSON(http.StatusInternalServerError, &resp)
		}

		resp.Success = true
		resp.Message = "success"
		resp.Data = tokenString
		c.JSON(http.StatusOK, &resp)
	}
}
