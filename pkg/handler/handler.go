package handler

import (
	"github.com/atadzan/AdvertAPI/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()

	advert := router.Group("/advert")
	{
		advert.POST("/", h.addAdvert)
		advert.GET("/", h.getAdverts)
		advert.GET("/:id", h.getAdvertById)
		advert.GET("/image/:id", h.getImage)
	}
	return router
}