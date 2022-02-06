package uirestapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DataApı struct {
	Title string `json:"title"`
	Imdb  int    `json:"imdb"`
}

func GetIMDB9(c *gin.Context) {

	k := RestIMDB9()
	var data []DataApı
	for _, v := range k {
		data = append(data, DataApı{Title: v.Title, Imdb: v.Imdb})

	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
