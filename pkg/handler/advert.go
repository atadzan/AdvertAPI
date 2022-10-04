package handler

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func(h *Handler) addAdvert(c *gin.Context){
	var advert AdvertAPI.AdvertInput
	if err := c.BindJSON(&advert); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	//fmt.Printf("Title: %s\n Description: %s\n Category: %s\n Location: %s\n Number: %s\n Price: %d\n",
	//	advert.Title, advert.Description, advert.Category, advert.Location, advert.PhoneNumber, advert.Price)
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
	adverts, err := h.services.Advert.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getAllAdvertResponse{
		Data: adverts,
	})
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
