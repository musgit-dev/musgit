package routes

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/musgit-dev/musgit/internal/musgit"
	"github.com/musgit-dev/musgit/models"
)

type Piece struct {
	Name       string `json:"name"  form:"name"       binding:"required"`
	Composer   string `json:"composer" form:"composer"   binding:"required"`
	Complexity int64  `json:"complexity" form:"complexity"`
}

func pieceRoutes(mg *musgit.Musgit, rg *gin.RouterGroup) {
	{
		pieces := rg.Group("/pieces")
		pieces.GET("/", func(c *gin.Context) {
			res := mg.Piece.GetAll()
			c.JSON(http.StatusOK, res)
		})
		pieces.GET("/:id", func(c *gin.Context) {
			id, _ := strconv.Atoi(c.Param("id"))
			res, err := mg.Piece.Get(int64(id))
			if err != nil {
				c.JSON(http.StatusNotFound, errors.New("Not found"))
			}
			c.JSON(http.StatusOK, res)
		})
		pieces.POST("/", func(c *gin.Context) {
			var newPiece Piece
			if err := c.ShouldBind(&newPiece); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			res, err := mg.Piece.Add(
				newPiece.Name,
				newPiece.Composer,
				models.PieceComplexity(newPiece.Complexity),
			)
			if err != nil {
				log.Fatal(err)
			}
			c.JSON(
				http.StatusOK,
				fmt.Sprintf(
					"Added new piece => [%s]: %s",
					res.Name,
					res.Composer.Name,
				),
			)
		})
	}

}
