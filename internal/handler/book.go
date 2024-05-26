package handler

import (
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