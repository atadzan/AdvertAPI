package handler

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func(h *Handler) addCategory(c *gin.Context){
	var categoryInput AdvertAPI.CategoryInput
	if err := c.BindJSON(&categoryInput); err != nil {
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

// @Summary     Get Main Categories
// @Tags        category
// @Description Get main categories
// @ID          get_category
// @Produce     json
// @Success     200     {array} AdvertAPI.CategoryOutput "status"
// @Failure     400     error   http.StatusBadRequest
// @Failure     500     error   http.StatusInternalServerError
// @Failure     default error   http.StatusBadRequest
// @Router      /api/category/main [get]
func(h *Handler) getMainCategory(c *gin.Context){
	mainCategories, err := h.services.Category.GetMain()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, mainCategories)
}

// @Summary     Get Nested Categories
// @Tags        category
// @Description Get Nested Categories of main category
// @ID          get_nested
// @Accept      json
// @Produce     json
// @Param       id      path    int                      true "main category_id"
// @Success     200     {array} AdvertAPI.CategoryOutput "Nested category"
// @Failure     400     error   http.StatusBadRequest
// @Failure     500     error   http.StatusInternalServerError
// @Failure     default error   http.StatusBadRequest
// @Router      /api/categories/{id} [get]
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

// @Summary     Get Category Adverts
// @Tags        category
// @Description Get Adverts by Category id
// @ID          get_category_adverts
// @Accept      json
// @Produce     json
// @Param       id      path    int                    true "category_id"
// @Success     200     {array} AdvertAPI.AdvertOutput "adverts"
// @Failure     400     error   http.StatusBadRequest
// @Failure     500     error   http.StatusInternalServerError
// @Failure     default error   http.StatusBadRequest
// @Router      /api/category/{id} [get]
func(h *Handler) getCategoryAdverts(c *gin.Context){
	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid category id param")
		return
	}
	categoryAdverts, ok := h.services.Category.GetCategoryAdverts(categoryId)
	if ok != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, categoryAdverts)
}
