package flightaware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/ATTron/flightaware/util"
	// autoload the .env file
	_ "github.com/joho/godotenv/autoload"
)

var apiKey = os.Getenv("API_KEY")
var attempts = 0
var layout = "2006-01-02T15:04:05Z"

// GetFlight return all data about certain flight number
func GetFlight(flightNum string) []Flight {
	flightEP := util.Join("flights/", flightNum)
	returnFlight := cleanData(&flightEP)

	return returnFlight.Flights
}

// GetFlightWithTime return a (hopefully) single flight for a timeframe
func GetFlightWithTime(flightNum string, flightTime string) []Flight {
	convertedFlightTime, endTime := convertTime(flightTime)
	fTime := util.Join("flights/", flightNum, "?start=", convertedFlightTime, "&end=", endTime)
	returnFlight := cleanData(&fTime)

	return returnFlight.Flights
}

func convertTime(timestamp string) (string, string) {
	ts, err := time.Parse(layout, timestamp)
	if err != nil {
		return "ERROR CONVERTING TIMESTAMP", "ERROR"
	}
	loc, _ := time.LoadLocation("UTC")
	cft := ts.In(loc)
	eft := cft.AddDate(0, 0, 1)

	return cft.Format("01-02-2006"), eft.Format("01-02-2006")
}

// fetchData - go and get the latest flight information
func fetchData(flightNum *string) (string, error) {
	client := http.Client{}
	content := ""
attempt:
	for attempts < 3 {
		reqEP := util.Join(util.EndPoint, *flightNum)
		req, err := http.NewRequest("GET", reqEP, nil)
		util.Check(err)

		req.Header.Add("x-apikey", apiKey)
		resp, err := client.Do(req)
		util.Check(err)

		defer resp.Body.Close()
		switch resp.StatusCode {
		case 200:
			responseData, err := ioutil.ReadAll(resp.Body)
			util.Check(err)
			content = string(responseData[:])
			break attempt
		default:
			attempts++
		}
	}
	if attempts >= 3 {
		return "", util.ErrNotFound
	}
	return content, nil
}

// cleanData - massage the data out
func cleanData(flightNum *string) Response {
	content, err := fetchData(flightNum)
	util.Check(err)

	returnFlight := Response{}

	err = json.Unmarshal([]byte(content), &returnFlight)
	util.Check(err)
	return returnFlight
}

// writeJSON - write out to file called 'av.json'
func writeJSON(content string) {
	f, err := os.Create("av.json")
	util.Check(err)

	defer f.Close()

	f.WriteString(content)

}
