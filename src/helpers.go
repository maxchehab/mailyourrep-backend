package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	strings "strings"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
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
func GetReps(address RequestData, selectedReps []string) (Reps, error) {
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
		if Contains(selectedReps, official.Name) {
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

//Payment charges a provided StripeToken $1.50 for each count provided
func Payment(stripeToken string, count int) error {
	fmt.Println(count)
	stripe.Key = STRIPE
	params := &stripe.ChargeParams{
		Amount:   150 * uint64(count),
		Currency: "usd",
		Desc:     "mailyourrep.org",
	}
	params.SetSource(stripeToken)

	_, err := charge.New(params)

	return err
}

//Contains checks if a string slice contanis an element
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

//Lob sends the letter to the selected representatives.
func Lob(requestData RequestData, rep Rep) error {

	return nil
}
