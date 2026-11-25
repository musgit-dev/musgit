package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/musgit-dev/musgit/internal/musgit"

	"github.com/gin-gonic/gin"
)

func lessonRoutes(mg *musgit.Musgit, rg *gin.RouterGroup) {
	lessons := rg.Group("/lessons")
	lessons.GET("/", func(c *gin.Context) {
		res := mg.Lesson.GetAll()
		c.JSON(http.StatusOK, res)
	})

	lessons.POST("/", func(c *gin.Context) {
		res, err := mg.Lesson.Start()
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, fmt.Sprintf("Started new lesson: %d", res.ID))
	})

	lessons.POST("/stop", func(c *gin.Context) {
		res := mg.Lesson.StopCurrent()
		c.JSON(http.StatusOK, res)
	})
	lessons.POST("/pause", func(c *gin.Context) {
		res := mg.Lesson.PauseCurrent()
		c.JSON(http.StatusOK, res)
	})
	lessons.POST("/resume", func(c *gin.Context) {
		res := mg.Lesson.PauseCurrent()
		c.JSON(http.StatusOK, res)
	})
}
