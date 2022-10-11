package handler

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func(h *Handler) addComment(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	advertId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid advert id param")
		return
	}
	var comment AdvertAPI.InputComm
	if err := c.BindJSON(&comment); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Comment.AddCom(comment, userId, advertId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError ,err.Error())
		return
	}
	c.JSON(http.StatusOK, "Success")
}

func(h *Handler) getComment(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	advertId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid advert id param")
		return
	}
	comments, err := h.services.Comment.GetCom(advertId, userId)
	if err != nil {
		newErrorResponse(c ,http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, comments)
}

func(h *Handler) delComment(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	advertId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid advert id param")
		return
	}
	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid comment id param")
	}
	err = h.services.Comment.DelCom(advertId, userId, commentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully deleted")
}
func(h *Handler) updComment(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	advertId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid advert id param")
		return
	}
	commentId, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid comment id param")
		return
	}
	var comment AdvertAPI.InputComm
	if err := c.BindJSON(&comment); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.Comment.UpdCom(comment, userId, advertId, commentId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully updated")
}

