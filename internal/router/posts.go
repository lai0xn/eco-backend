package router

import (
	"github.com/labstack/echo/v4"
	"github.com/lai0xn/squid-tech/internal/handlers"
)

func postRoutes(e *echo.Group) {
	h := handlers.NewPostsHandler()
	g := e.Group("/posts")
	g.Use(jwtMiddelware)
	g.GET("/post/get/:id", h.Get)
	g.GET("", h.GetPage)
	g.GET("/post/search", h.Search)
	g.POST("/create", h.Create)
	g.POST("/comment", h.Comment)
	g.POST("/post/:id/image", h.UploadImage)
	g.PATCH("/post/:id/update", h.Update)
	g.DELETE("/post/:id/delete", h.Delete)
	g.DELETE("/post/comments/:id/delete", h.DeleteComment)

}
