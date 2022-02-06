package uirestapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestApı7 struct {
	Title string `json:"title"`
	Imdb  int    `json:"imdb"`
}

func GetIMDB7(c *gin.Context) {

	k := RestIMDB7()
	var data []RestApı7
	for _, v := range k {
		data = append(data, RestApı7{Title: v.Title, Imdb: v.Imdb})

	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
