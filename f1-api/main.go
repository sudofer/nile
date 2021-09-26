package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

// TODO: Figure out if this struct is really necessary
type Response struct {
	MrData MRData `json:"MRData"`
}

type MRData struct {
	Xmlns          string           `json:"xmlns"`
	Series         string           `json:"series"`
	URL            string           `json:"url"`
	Limit          string           `json:"limit"`
	Offset         string           `json:"offset"`
	Total          string           `json:"total"`
	StandingsTable []StandingsTable `json:"StandingsTable"`
}

type StandingsTable struct {
	Season    string      `json:"season"`
	Standings []Standings `json:"StandingsLists"`
}

type Standings struct {
	Season          string           `json:"season"`
	Round           string           `json:"round"`
	DriverStandings []DriverStanding `json:"DriverStandings"`
}

type DriverStanding struct {
	Position     string        `json:"position"`
	PositionText string        `json:"positionText"`
	Points       string        `json:"points"`
	Wins         string        `json:"wins"`
	Driver       Driver        `json:"Driver"`
	Constructors []Constructor `json:"Constructors"`
}

type Driver struct {
	DriverID        string `json:"driverId"`
	PermanentNumber string `json:"permanentNumber"`
	Code            string `json:"code"`
	URL             string `json:"url"`
	GivenName       string `json:"givenName"`
	FamilyName      string `json:"familyName"`
	DateOfBirth     string `json:"dateOfBirth"`
	Nationality     string `json:"nationality"`
}

type Constructor struct {
	ConstructorID string `json:"constructorId"`
	URL           string `json:"url"`
	Name          string `json:"name"`
	Nationality   string `json:"nationality"`
}

func main() {
	mrData := new(Response)

	getJson("http://ergast.com/api/f1/current/driverStandings.json", mrData)

	fmt.Printf("%#v", mrData)
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
