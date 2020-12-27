package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const sourceURL = "https://dashboard.saskatchewan.ca/api/health/indicator/detail/health-wellness%3Acovid-19%3Acases?legacyRegions=true"

type CovidData struct {
	TabTitles      []string    `json:"tabTitles"`
	RevisionDate   string      `json:"revisionDate"`
	ReleaseDate    string      `json:"releaseDate"`
	UpdateDate     interface{} `json:"updateDate"`
	GraphFootnotes interface{} `json:"graphFootnotes"`
	Highlights     string      `json:"highlights"`
	Tabs           []struct {
		Chart struct {
			ChartTitle string `json:"chartTitle"`
			ChartType  string `json:"chartType"`
			YAxis      string `json:"yAxis"`
			Data       []struct {
				SeriesTitle string `json:"seriesTitle"`
				SeriesID    string `json:"seriesId"`
				Color       string `json:"color"`
				Group       string `json:"group"`
				Data        []struct {
					Time  int `json:"time"`
					Value int `json:"value"`
				} `json:"data"`
			} `json:"data"`
		} `json:"chart"`
		Tables []struct {
			Title     string `json:"title"`
			IsVisible bool   `json:"isVisible"`
			Header    struct {
				Cells []struct {
					Value    string `json:"value"`
					Footnote string `json:"footnote"`
				} `json:"cells"`
			} `json:"header"`
			Body []struct {
				Cells []struct {
					Value    json.RawMessage `json:"value"`
					Footnote string          `json:"footnote"`
				} `json:"cells"`
			} `json:"body"`
		} `json:"tables"`
	} `json:"tabs"`
}

func main() {
	res, err := http.Get(sourceURL)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var raw CovidData
	err = json.Unmarshal(buf, &raw)

	if err != nil {
		panic(err)
	}
	for _, v := range raw.Tabs[0].Chart.Data[0].Data {
		fmt.Printf("%v: %v\n", time.Unix(int64(v.Time), 0), v.Value)
	}
}
