package covid

import (
	"encoding/json"
	"net/http"
	"time"
)

type Full struct {
	Code int
	Data []Data
}

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

func Info(info *Full) error {
	resp, err := http.Get("https://www.trackcorona.live/api/countries")
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(info)
}
