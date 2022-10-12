package handler

import (
	_ "github.com/atadzan/AdvertAPI/docs"
	"github.com/atadzan/AdvertAPI/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine{
	router := gin.New()

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api", h.userIdentity)
	{
		advert := api.Group("/advert")
		{
			advert.POST("/", h.addAdvert)
			advert.DELETE("/:id", h.deleteAdvert)
			advert.PUT("/edit/:id", h.updateAdvert)
			advert.PUT("/:id", h.addFavList)
			advert.GET("fav", h.getFavList)
			advert.DELETE("fav/:id", h.deleteFav)
		}
		comment := api.Group(":id/comment")
		{
			comment.POST("/", h.addComment)
			comment.DELETE("/:comment_id", h.delComment)
			comment.PUT("/:comment_id", h.updComment)
		}
	}
	advert := router.Group("api/advert")
	{
		advert.GET("/", h.getAdverts)
		advert.GET("/:id", h.getAdvertById)
		advert.GET("/search", h.searchTitle)
	}
	auth := router.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	return router
}