package handler

import (
	"net/http"
	"online_lib_api/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://192.168.0.123:5173"},
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodHead, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.POST("/profile-data", h.getProfileData)
		api.POST("/add-book", h.addBook)
		api.POST("/delete-book", h.deleteBook)
	}

	api_free := router.Group("/api-free")
	{
		api_free.GET("/book-content/:book_name_id", h.getContentBook)
		api_free.GET("/all-books", h.getAllBooks)
		api_free.GET("/about-book/:book_name_id", h.getAboutBook)
		api_free.GET("/popular-genres", h.getPopularGenres)
	}

	// router.Static("/assets", "./assets")

	static := router.Group("/static")
	{
		static.Static("", "./static")
	}

	return router
}
