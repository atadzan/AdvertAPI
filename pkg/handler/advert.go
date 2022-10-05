package handler

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func(h *Handler) addAdvert(c *gin.Context){
	var advert AdvertAPI.AdvertInput
	if err := c.BindJSON(&advert); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Advert.Add(advert)
	if err != nil{
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Advert ID": id,
	})
}

type getAllAdvertResponse struct {
	Data []AdvertAPI.AdvertInfo
}

func(h *Handler) getAdverts(c *gin.Context){
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	advertCount, err := h.services.Advert.CountAdverts()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	const advertPerPage = 5
	pageCount := int(math.Ceil(float64(advertCount) / float64(advertPerPage)))
	if pageCount == 0 {
		pageCount = 1
	}
	if page < 1 || page > pageCount{
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	offset := (page - 1) * advertPerPage
	adverts, err := h.services.Advert.GetAll(advertPerPage, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllAdvertResponse{
		Data: adverts,
	})




	//adverts, err := h.services.Advert.GetAll()
	//if err != nil {
	//	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	//	return
	//}
	//c.JSON(http.StatusOK, getAllAdvertResponse{
	//	Data: adverts,
	//})
}

func(h *Handler) getAdvertById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	advert, err := h.services.Advert.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, advert)
}
