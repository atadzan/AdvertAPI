package handler

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func(h *Handler) addCategory(c *gin.Context){
	_, err := getUserId(c)
	if err != nil {
		return
	}
	var categoryInput AdvertAPI.CategoryInput
	if err = c.BindJSON(&categoryInput); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	categoryId, ok := h.services.Category.Add(categoryInput)
	if ok != nil {
		newErrorResponse(c, http.StatusInternalServerError, ok.Error())
		return
	}
	c.JSON(http.StatusOK,  categoryId)
}

func(h *Handler) getMainCategory(c *gin.Context){
	mainCategories, err := h.services.Category.GetMain()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, mainCategories)
}

func(h *Handler) getNestedCategories(c *gin.Context){
	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid category id param")
		return
	}
	nestedCategories, ok := h.services.Category.GetNested(categoryId)
	if ok != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nestedCategories)
}
