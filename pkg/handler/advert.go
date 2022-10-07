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

func(h *Handler) addAdvert(c *gin.Context){
	form, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	advert, err := InputProcess(c, form)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
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

func InputProcess(c *gin.Context, form *multipart.Form, )(AdvertAPI.AdvertInput, error){
	var advert AdvertAPI.AdvertInput
	advert.Title = c.PostForm("title")
	advert.Description = c.PostForm("description")
	advert.Category = c.PostForm("category")
	advert.Location = c.PostForm("location")
	advert.PhoneNumber = c.PostForm("phone_number")
	advert.Price, _ = strconv.Atoi(c.PostForm("price"))
	images := form.File["images"]
	var file AdvertAPI.AdvertImage
	if images != nil {
		for i := range images{
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

//func(h *Handler) updateAdvert(c *gin.Context){
//	id, err := strconv.Atoi(c.Param("id"))
//	if err != nil {
//		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
//		return
//	}
//	form, err := c.MultipartForm()
//	if err != nil {
//		newErrorResponse(c, http.StatusBadRequest, err.Error())
//		return
//	}
//	advert, err := InputProcess(c, form)
//	if err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err.Error())
//	}
//	if err := h.services.Advert.Update(id, advert); err != nil {
//		newErrorResponse(c, http.StatusInternalServerError, err)
//		return
//	}
//	c.JSON(http.StatusOK,"OK")
//}
