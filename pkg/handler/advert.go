package handler

import (
	"github.com/atadzan/AdvertAPI"
	"github.com/gin-gonic/gin"
	"io"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

// @Summary     Add Advert
// @Security    ApiKeyAuth
// @Tags        advert
// @Description Add Advert to DB
// @ID          add_advert
// @Accept      mpfd
// @Produce     json
// @Param       form    formData AdvertAPI.AdvertInput true "advert info"
// @Success     200     {string} string                "id"
// @Failure     400     error    http.StatusBadRequest
// @Failure     500     error    http.StatusInternalServerError
// @Failure     default error    http.StatusBadRequest
// @Router      /api/advert [post]
func (h *Handler) addAdvert(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	advert, err := InputProcess(c, form, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	id, err := h.services.Advert.Add(advert)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Advert ID": id,
	})
}

func InputProcess(c *gin.Context, form *multipart.Form, userId int) (AdvertAPI.AdvertInput, error) {
	var advert AdvertAPI.AdvertInput
	advert.Title = c.PostForm("title")
	advert.Description = c.PostForm("description")
	advert.Category, _ = strconv.Atoi(c.PostForm("category"))
	advert.Location = c.PostForm("location")
	advert.PhoneNumber = c.PostForm("phone_number")
	advert.Price, _ = strconv.Atoi(c.PostForm("price"))
	advert.UserId = userId
	images := form.File["images"]
	var file AdvertAPI.AdvertImage
	if images != nil {
		for i := range images {
			image, err := images[i].Open()
			if err != nil {
				newErrorResponse(c, http.StatusNoContent, err.Error())
				return advert, err
			}
			defer image.Close()
			file.Fname = images[i].Filename
			file.Fsize = images[i].Size
			file.Ftype = images[i].Header.Get("Content-type")

			//	Create file
			tempFile, err := os.CreateTemp("assets/uploadImages", "*.jpg")
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return advert, err
			}
			filepath := tempFile.Name()
			file.Path = filepath

			//read all the contents of our uploaded file into a byte array
			fileBytes, err := io.ReadAll(image)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return advert, err
			}

			//Write this byte array to our temporary array
			_, err = tempFile.Write(fileBytes)
			if err != nil {
				newErrorResponse(c, http.StatusInternalServerError, err.Error())
				return advert, err
			}

			advert.Images = append(advert.Images, file)
		}
	}
	return advert, nil
}

// @Summary     Get Adverts
// @Tags        advert
// @Description Get adverts by page
// @ID          get_adverts
// @Accept      json
// @Produce     json
//@Param       page    query   string false "page info"
// @Success     200     {array} AdvertAPI.AdvertInfo
// @Failure     400     error   http.StatusBadRequest
// @Failure     500     error   http.StatusInternalServerError
// @Failure     default error   http.StatusBadRequest
// @Router      /api/advert [get]
func (h *Handler) getAdverts(c *gin.Context) {
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
	if page < 1 || page > pageCount {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	offset := (page - 1) * advertPerPage
	adverts, err := h.services.Advert.GetAll(advertPerPage, offset)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, adverts)
}

// @Summary     Get Advert by ID
// @Tags        advert
// @Description Get Advert by ID
// @ID          get_advert
// @Accept      json
// @Produce     json
// @Param       id      path     int true "advert ID"
// @Success     200     {object} AdvertAPI.AdvertInfo
// @Failure     400     error    http.StatusBadRequest
// @Failure     500     error    http.StatusInternalServerError
// @Failure     default error    http.StatusBadRequest
// @Router      /api/advert/{id} [get]
func (h *Handler) getAdvertById(c *gin.Context) {
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

func(h *Handler) getImage(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	object, err := h.services.Advert.GetImage(id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}
	for _, i := range object{
		fi, err := os.Open(i.Path)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		_, err = io.Copy(c.Writer, fi)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		return
	}
	return
}

// @Summary     Delete Advert
// @Security    ApiKeyAuth
// @Tags        advert
// @Description Delete Advert
// @ID          del_advert
// @Accept      json
// @Produce     json
// @Param       id      path     int    true "advert ID"
// @Success     200     {string} string "status"
// @Failure     400     error    http.StatusBadRequest
// @Failure     500     error    http.StatusInternalServerError
// @Failure     default error    http.StatusBadRequest
// @Router      /api/advert/{id} [delete]
func (h *Handler) deleteAdvert(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	err = h.services.Advert.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary     Update Advert
// @Security    ApiKeyAuth
// @Tags        advert
// @Description Update Advert
// @ID          update_advert
// @Accept      json
// @Produce     json
// @Param       form    formData AdvertAPI.AdvertInfo true "update advert info"
// @Success     200     {string} string               "status"
// @Failure     400     error    http.StatusBadRequest
// @Failure     500     error    http.StatusInternalServerError
// @Failure     default error    http.StatusBadRequest
// @Router      /api/advert/{id} [put]
func (h *Handler) updateAdvert(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	form, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	advert, err := InputProcess(c, form, userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if err := h.services.Advert.Update(id, advert); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "Successfully updated",
	})
}

// @Summary     Add Advert to Favourite List
// @Security    ApiKeyAuth
// @Tags        advert
// @Description Add Advert to Favourite List
// @ID          add_fav
// @Accept      json
// @Produce     json
// @Param       id      path     int    true "credentials"
// @Success     200     {string} string "status"
// @Failure     400     error    http.StatusBadRequest
// @Failure     500     error    http.StatusInternalServerError
// @Failure     default error    http.StatusBadRequest
// @Router      /api/advert/fav/{id} [post]
func (h *Handler) addFavList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Advert.AddFav(userId, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully added")
}

// @Summary     Get User Favourite List
// @Security    ApiKeyAuth
// @Tags        advert
// @Description Get User Favourite List
// @ID          get_fav
// @Produce     json
// @Success     200     {string} string "status"
// @Failure     400     error    http.StatusBadRequest
// @Failure     500     error    http.StatusInternalServerError
// @Failure     default error    http.StatusBadRequest
// @Router      /api/advert/fav [get]
func (h *Handler) getFavList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	favAdverts, err := h.services.Advert.GetFav(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, favAdverts)
}

// @Summary     Delete Advert from Favourite List
// @Security    ApiKeyAuth
// @Tags        advert
// @Description Delete Advert from Favourite List
// @ID          del_fav
// @Accept      json
// @Produce     json
// @Param       id      path     int    true "advert"
// @Success     200     {string} string "status"
// @Failure     400     error    http.StatusBadRequest
// @Failure     500     error    http.StatusInternalServerError
// @Failure     default error    http.StatusBadRequest
// @Router      /api/advert/fav/{id} [delete]
func (h *Handler) deleteFav(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Advert.DeleteFav(userId, id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully deleted")
}

// @Summary     Check Favourite List
// @Security    ApiKeyAuth
// @Tags        fav_list
// @Description Check Advert from Favourite List
// @ID          fav_list
// @Accept      json
// @Produce     json
// @Param       id      path     int    true "advert"
// @Success     200     {bool} bool "status"
// @Failure     400     error    http.StatusBadRequest
// @Failure     500     error    http.StatusInternalServerError
// @Failure     default error    http.StatusBadRequest
// @Router      /api/advert/fav/{id} [get]
func(h *Handler) checkFav(c *gin.Context){
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response, ok := h.services.Advert.CheckFavList(userId, id)
	if ok != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}

// @Summary     Search
// @Tags        advert
// @Description Search Adverts by Title
// @ID          search_adv
// @Accept      json
// @Produce     json
// @Param       title   query   string               true "title"
// @Success     200     {array} AdvertAPI.AdvertInfo "status"
// @Failure     400     error   http.StatusBadRequest
// @Failure     500     error   http.StatusInternalServerError
// @Failure     default error   http.StatusBadRequest
// @Router      /api/advert/search [get]
func(h *Handler) searchByTitle(c *gin.Context){
	title := c.DefaultQuery("title", "Advert")
	adverts, err := h.services.Advert.Search(title)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, adverts)
}
