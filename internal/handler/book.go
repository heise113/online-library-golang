package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getContentBook(c *gin.Context) {
	book_name_id := c.Param("book_name_id")
	book_conent, err := h.services.GetContentBook(book_name_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"book_content": book_conent,
	})
}

func (h *Handler) getAllBooks(c *gin.Context) {
	all_books, err := h.services.GetAllBooks()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, all_books)
}

func (h *Handler) getAboutBook(c *gin.Context) {
	book_name_id := c.Param("book_name_id")
	fmt.Println(book_name_id)
	about_book, err := h.services.GetAboutBook(book_name_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, about_book)
}

func (h *Handler) getPopularGenres(c *gin.Context) {
	popular_genres, err := h.services.GetPopularGenres()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, popular_genres)
}