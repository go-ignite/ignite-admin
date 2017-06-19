package controllers

import (
	"fmt"
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
	loginEntity := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BindJSON(&loginEntity); err != nil {
		resp := models.Response{Success: false, Message: "Could not parse username & password..."}
		c.JSON(http.StatusInternalServerError, &resp)
		return
	}

	fmt.Println("username:", loginEntity.Username)
	fmt.Println("pwd:", loginEntity.Password)

	fmt.Println("authUser:", authUser)
	fmt.Println("authPass:", authPass)

	if loginEntity.Username == authUser && loginEntity.Password == authPass {
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
			return
		}

		resp.Success = true
		resp.Message = "success"
		resp.Data = tokenString
		c.JSON(http.StatusOK, &resp)
		return
	}

	resp := models.Response{Success: false, Message: "Username of password is wrong!"}
	c.JSON(http.StatusOK, &resp)
}
