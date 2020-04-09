package covid

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Full struct {
	Code int    `json:"code"`
	Data []Data `json:"data"`
}

type Data struct {
	Location    string  `json:"location"`
	CountryCode string  `json:"country_code"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Confirmed   int     `json:"confirmed"`
	Dead        int     `json:"dead"`
	Recovered   int     `json:"recovered"`
	Updated     string  `json:"updated"`
}

func Info(countryName string) int {
	resp, err := http.Get("https://www.trackcorona.live/api/countries")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var info Full
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &info)
	if err != nil {
		panic(err)
	}
	i := 0
	var confirmed int
	for _, v := range info.Data {
		if countryName == v.Location {
			confirmed = v.Confirmed
			i++
		}
	}
	if i == 0 {
		return -1
	}
	return confirmed
}
