package uirestapi

import (
	"golang-rest-data/utils"
)

type RestAPIData struct {
	Title string
	Imdb  int
}

func RestIMDB9() []RestAPIData {
	IMDB9 := utils.Imdb9SQL()

	var restapıdata RestAPIData

	var data []RestAPIData

	for k, v := range IMDB9 {
		restapıdata.Title = k
		restapıdata.Imdb = v
		data = append(data, RestAPIData{Title: restapıdata.Title, Imdb: restapıdata.Imdb})

	}
	return data
}

func RestIMDB8() []RestAPIData {
	IMDB8 := utils.Imdb8SQL()

	var restapıdata RestAPIData

	var data []RestAPIData

	for k, v := range IMDB8 {
		restapıdata.Title = k
		restapıdata.Imdb = v
		data = append(data, RestAPIData{Title: restapıdata.Title, Imdb: restapıdata.Imdb})

	}
	return data

}

func RestIMDB7() []RestAPIData {
	IMDB7 := utils.Imdb7SQL()

	var restapıdata RestAPIData

	var data []RestAPIData

	for k, v := range IMDB7 {
		restapıdata.Title = k
		restapıdata.Imdb = v
		data = append(data, RestAPIData{Title: restapıdata.Title, Imdb: restapıdata.Imdb})

	}
	return data

}

func RestIMDB6() []RestAPIData {
	IMDB6 := utils.Imdb6SQL()

	var restapıdata RestAPIData

	var data []RestAPIData

	for k, v := range IMDB6 {
		restapıdata.Title = k
		restapıdata.Imdb = v
		data = append(data, RestAPIData{Title: restapıdata.Title, Imdb: restapıdata.Imdb})

	}
	return data

}

func RestIMDB5() []RestAPIData {
	IMDB6 := utils.Imdb5SQL()

	var restapıdata RestAPIData

	var data []RestAPIData

	for k, v := range IMDB6 {
		restapıdata.Title = k
		restapıdata.Imdb = v
		data = append(data, RestAPIData{Title: restapıdata.Title, Imdb: restapıdata.Imdb})

	}
	return data

}
