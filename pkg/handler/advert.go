package handler

import (
	"fmt"
	"github.com/atadzan/AdvertAPI"
	"github.com/gin-gonic/gin"
	"io"
	"math"
	"net/http"
	"os"
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

func(h *Handler) uploadImage(c *gin.Context){
	//name := c.PostForm("name")
	form, err := c.MultipartForm()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	files := form.File["files"]
	for i := range files {
		file, err := files[i].Open()
		if err != nil {
			newErrorResponse(c, http.StatusNoContent, err.Error())
			return
		}
		defer file.Close()
		fname := files[i].Filename
		fsize := files[i].Size
		//kilobytes := fsize / 1024
		ftype := files[i].Header.Get("Content-type")
		fmt.Println(fname)
		fmt.Printf("%t", fsize)

		fmt.Println(ftype)

		//	Create file
		tempFile, err := os.CreateTemp("assets/uploadImages", "upload-*.jpg")
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		filepath := tempFile.Name()

		//read all the contents of our uploaded file into a byte array
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		//Write this byte array to our temporary array
		_, err = tempFile.Write(fileBytes)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		//return upload file message
		insForm, err := h.services.AddDB(fname, ftype, filepath, fsize)
		//defer tempFile.Close()
		fmt.Println(insForm)
	}
}

func(h *Handler) getImage(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	image, err := h.services.Advert.GetImage(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, image)
}
