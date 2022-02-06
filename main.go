package main

import (
	uirestapi "golang-rest-data/UI_RESTAPI"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.LoadHTMLFiles("templates/index.html")
	r.GET("/", func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{

			"title": "JSON DATA LİNK",
		})

	})

	r.GET("/movie/ımdb9", uirestapi.GetIMDB9)
	r.GET("/movie/ımdb8", uirestapi.GetIMDB8)
	r.GET("/movie/ımdb7", uirestapi.GetIMDB7)
	r.GET("/movie/ımdb6", uirestapi.GetIMDB6)
	r.GET("/movie/ımdb5", uirestapi.GetIMDB5)

	r.Run(":8080")

}
