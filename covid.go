package covid

import (
	"encoding/json"
	"net/http"
	"time"
)

type Data struct {
	Location    string
	CountryCode string
	Latitude    float64
	Longtitude  float64
	Confirmed   int
	Dead        int
	Recovered   int
	Updated     time.Time
}

func Info(data *[]Data) error {
	resp, err := http.Get("https://www.trackcorona.live/api/countries")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(data)
}
