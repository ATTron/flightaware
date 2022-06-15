package flightaware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ATTron/flightaware/util"
	// autoload the .env file
	_ "github.com/joho/godotenv/autoload"
)

var apiKey = os.Getenv("API_KEY")
var attempts = 0

// GetFlight -- return data about flight
func GetFlight(flightNum string) Flight {
	flightEP := util.Join("flights/", flightNum)
	returnFlight := cleanData(&flightEP)
	return returnFlight
}

// fetchData - go and get the latest train information
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
func cleanData(flightNum *string) Flight {
	content, err := fetchData(flightNum)
	util.Check(err)

	returnFlight := Flight{}

	err = json.Unmarshal([]byte(content), &returnFlight)
	util.Check(err)
	return returnFlight
}

// WriteJSON - write out to file called 'av.json'
func writeJSON(content string) {
	f, err := os.Create("av.json")
	util.Check(err)

	defer f.Close()

	f.WriteString(content)

}
