package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getProfileData(c *gin.Context) {
	user_id, _ := c.Get(userCtx)
	profile_data, err := h.services.GetProfileData(user_id.(int))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, profile_data)
}

func (h *Handler) addBook(c *gin.Context) {
	var book_id int
	if err := c.BindJSON(&book_id); err != nil {
		fmt.Println(err)
	}
	user_id, _ := c.Get(userCtx)
	err := h.services.AddBook(user_id.(int), book_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "the book has been added successfully")
}

func (h *Handler) deleteBook(c *gin.Context) {
	var book_id int
	if err := c.BindJSON(&book_id); err != nil {
		fmt.Println(err)
	}
	user_id, _ := c.Get(userCtx)
	err := h.services.DeleteBook(user_id.(int), book_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "the book has been deleted successfully")
}