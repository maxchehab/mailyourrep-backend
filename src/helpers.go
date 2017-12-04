package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	strings "strings"

	"github.com/ttacon/chalk"
)

//Success styles success messages with a green background
func Success(s interface{}) {
	fmt.Println(
		chalk.White.NewStyle().WithBackground(chalk.Green).Style(s.(string)))
}

//Error styles error messages with a red background
func Error(s interface{}) {
	fmt.Println(
		chalk.White.NewStyle().WithBackground(chalk.Red).Style(s.(string)))
}

//CreateRequestData creates a RequestData structure from an http.Request
func CreateRequestData(r *http.Request) (RequestData, error) {
	decoder := json.NewDecoder(r.Body)
	var requestData RequestData
	err := decoder.Decode(&requestData)
	if err != nil {
		return requestData, err
	}
	defer r.Body.Close()
	return requestData, err
}

//GetReps uses Google's civil api and retrieves representatives relating to a provided address
func GetReps(address RequestData) (Reps, error) {
	var url = CreateRepURL(address)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	b := []byte(body)
	var googleCivic GoogleCivic

	err = json.Unmarshal(b, &googleCivic)
	if err != nil {
		panic(err)
	}

	var indices []int
	for _, office := range googleCivic.Offices {
		if strings.Contains(office.Name, "Representatives") || strings.Contains(office.Name, "Senate") {
			indices = append(indices, office.OfficialIndices...)
		}
	}

	var reps Reps

	for _, index := range indices {
		var official = googleCivic.Officials[index]
		reps = append(reps, Rep{
			official.Name,
			official.Address[0].Zip,
			official.Address[0].Line1,
			official.Address[0].Line2,
			official.Address[0].City,
			official.Address[0].State,
			"USA",
		})
	}

	return reps, nil
}

//CreateRepURL constructs a usable Google api query with a given address.
func CreateRepURL(address RequestData) string {
	var query = address.Line1 + " " + " " + address.City + " " + address.State + " " + address.Zipcode

	url := &url.URL{Path: query}
	query = CIVICINFO + url.String()
	return query
}
