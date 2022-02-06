package uirestapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestApı5 struct {
	Title string `json:"title"`
	Imdb  int    `json:"imdb"`
}

func GetIMDB5(c *gin.Context) {

	k := RestIMDB5()
	var data []RestApı5
	for _, v := range k {
		data = append(data, RestApı5{Title: v.Title, Imdb: v.Imdb})

	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
