package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	taghvimcomURLFormat = "https://www.taghvim.com/get_events/?action=get_events&month=%d"

	taghvimcom [12]*TaghvimcomMonth
)

type (
	TaghvimcomMonth []TaghvimcomDay

	TaghvimcomDay struct {
		Day       int
		Month     string
		Name      string
		IsHoliday bool `json:"is_holiday"`
		Num       int  `json:"month_order"`
	}
)

func (tm TaghvimcomMonth) IsHoliday(d int) bool {
	for _, td := range tm {
		if td.Day == d && td.IsHoliday {
			return true
		}
	}
	return false
}

func taghvimcomHolidays() []string {
	holidays := make([]string, 0)

	for i := 1; i <= 12; i++ {
		tm, err := getTaghvimcomMonth(i)
		if err != nil {
			log.Fatal(err)
		}

		taghvimcom[i-1] = tm
	}

	for m := 1; m < 13; m++ {
		for d := 1; d <= DOM[m-1]; d++ {
			if isFriday(YEAR, m, d) || taghvimcom[m-1].IsHoliday(d) {
				addHoliday(&holidays, YEAR, m, d)
			}
		}
	}

	return holidays
}

func getTaghvimcomMonth(m int) (*TaghvimcomMonth, error) {
	url := fmt.Sprintf(taghvimcomURLFormat, m)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Requested-With", "XMLHttpRequest")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	tm := new(TaghvimcomMonth)

	err = json.NewDecoder(res.Body).Decode(tm)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	return tm, nil
}
