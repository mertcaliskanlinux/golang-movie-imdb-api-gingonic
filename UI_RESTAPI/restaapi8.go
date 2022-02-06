package uirestapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestApı8 struct {
	Title string `json:"title"`
	Imdb  int    `json:"imdb"`
}

func GetIMDB8(c *gin.Context) {

	k := RestIMDB8()
	var data []RestApı8
	for _, v := range k {
		data = append(data, RestApı8{Title: v.Title, Imdb: v.Imdb})

	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
