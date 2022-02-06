package utils

import (
	"encoding/csv"
	"golang-rest-data/config"
	"log"
	"os"
	"strconv"
)

type CsvDATA struct {
	Unnamed int
	Rank    int
	Title   string
	ImdbID  string
	Avg     float64
	Score   int
}

func readCSV() [][]string {
	csvFile, err := os.Open("movie.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	records, _ := reader.ReadAll()

	return records

}

func readAllCsv() map[string]float64 {

	records := readCSV()

	var csvEmp CsvDATA
	result := make(map[string]float64)

	for _, v := range records {
		csvEmp.Unnamed, _ = strconv.Atoi(v[0])
		csvEmp.Rank, _ = strconv.Atoi(v[1])
		csvEmp.Title = v[2]
		csvEmp.ImdbID = v[3]
		csvEmp.Avg, _ = strconv.ParseFloat(v[4], 64)
		csvEmp.Score, _ = strconv.Atoi(v[5])
		result[csvEmp.Title] = csvEmp.Avg

	}

	return result
}

type CsvTitleAndAVG struct {
	Title string
	Avg   float64
}

func TitleAndAvgCsv() {

	TitleAndAvgResult := readAllCsv()

	var csvtitleandavg CsvTitleAndAVG

	for title, avg := range TitleAndAvgResult {
		csvtitleandavg.Title = title
		csvtitleandavg.Avg = avg
		config.InsertDataBase(csvtitleandavg.Title, csvtitleandavg.Avg)

	}

}
