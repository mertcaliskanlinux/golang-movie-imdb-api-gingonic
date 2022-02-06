package uirestapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestApı6 struct {
	Title string `json:"title"`
	Imdb  int    `json:"imdb"`
}

func GetIMDB6(c *gin.Context) {

	k := RestIMDB6()
	var data []RestApı6
	for _, v := range k {
		data = append(data, RestApı6{Title: v.Title, Imdb: v.Imdb})

	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
