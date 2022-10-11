package handler

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary     Register
// @Tags        auth
// @Description Create account in app
// @ID          create-account
// @Accept      json
// @Produce     json
// @Param       input   body      AdvertAPI.SignUpInput true "account info"
// @Success     200     {integer} integer               1
// @Failure     400     error     http.StatusBadRequest
// @Failure     500     error     http.StatusInternalServerError
// @Failure     default error     http.StatusBadRequest
// @Router      /auth/sign-up [post]
func(h *Handler) signUp(c *gin.Context){
	var input AdvertAPI.SignUpInput
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, id)
}

// @Summary     Login
// @Tags        auth
// @Description Sign in app
// @ID          login
// @Accept      json
// @Produce     json
// @Param       input   body     AdvertAPI.SignInInput true "credentials"
// @Success     200     {string} string                "token"
// @Failure     400     error    http.StatusBadRequest
// @Failure     500     error    http.StatusInternalServerError
// @Failure     default error    http.StatusBadRequest
// @Router      /auth/sign-in [post]
func(h *Handler) signIn(c *gin.Context){
	var input AdvertAPI.SignInInput
	if err := c.BindJSON(&input); err != nil{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}
