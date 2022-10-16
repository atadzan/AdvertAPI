package handler

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary     Add Comment
// @Security    ApiKeyAuth
// @Tags        comment
// @Description Add Comment to Advert
// @ID          add_comment
// @Accept      json
// @Produce     json
// @Param       id      path     int                 true "advert id"
// @Param       comment body     AdvertAPI.InputComm true "comment body"
// @Success     200     {string} string              "id"
// @Failure     400     error    http.StatusBadRequest
// @Failure     500     error    http.StatusInternalServerError
// @Failure     default error    http.StatusBadRequest
// @Router      /api/{id}/comment [post]
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
	if err = c.BindJSON(&comment); err != nil {
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

// @Summary     Delete Comment
// @Security    ApiKeyAuth
// @Tags        comment
// @Description Delete Advert Comment by ID
// @ID          delete_comment
// @Accept      json
// @Produce     json
// @Param       id      path     int true "advert ID"
// @Param       comment_id    path     int true "comment ID"
// @Success     200     {string} http.StatusOK
// @Failure     400     error    http.StatusBadRequest
// @Failure     500     error    http.StatusInternalServerError
// @Failure     default error    http.StatusBadRequest
// @Router      /api/{id}/comment/{comment_id} [delete]
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

// @Summary     Update Comment
// @Security    ApiKeyAuth
// @Tags        comment
// @Description Update Advert Comment By ID
// @ID          update comment
// @Accept      json
// @Produce     json
// @Param       id         path     int true "advert ID"
// @Param       comment_id path     int true "comment ID"
// @Param       comment body     AdvertAPI.InputComm true "comment body"
// @Success     200        {string} http.StatusOK
// @Failure     400        error    http.StatusBadRequest
// @Failure     500        error    http.StatusInternalServerError
// @Failure     default    error    http.StatusBadRequest
// @Router      /api/{id}/comment/{comment_id} [put]
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

