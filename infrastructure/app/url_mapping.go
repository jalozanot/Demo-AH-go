package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jalozanot/demoCeiba/infrastructure/controllers"
)

func mapUrls(handler controllers.RedirectMovieHandler) *gin.Engine {

	ping := router.Group("/")
	{
		ping.GET("/ping", controllers.Ping)
	}

	movie := router.Group("/")
	{
		movie.POST("/peliculas", handler.Create)
		movie.GET("/peliculas/:id", handler.Get)
		movie.PUT("/peliculas/:id", handler.Update)
		movie.DELETE("/peliculas/:id", handler.Delete)
	}

	//cliente := ...

	return router
}
