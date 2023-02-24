package handler

import (

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/katakuxiko/clean_go/package/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	confgisGin := cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000","https://silver-capybara-b0969e.netlify.app","https://avxm.netlify.app"},
		AllowMethods:     []string{"POST","GET","PUT", "PATCH","Delete"},
		AllowHeaders:     []string{"Origin","Authorization","Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
	})
	router.Use(confgisGin)
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/refresh", h.RefreshToken)
	}
	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)

			}
		}
		userList := api.Group("/userList")
			{
				userList.GET("/", h.getUserBooksAll)
			}
		items := api.Group("/items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}
	
	

	return router
}
