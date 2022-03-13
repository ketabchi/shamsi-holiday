package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

var timeirURLFormat = "http://www.time.ir/fa/event/list/0/%d/%d/%d"

func timeirHolidays(y int) []string {
	var wg sync.WaitGroup
	holidays := make([]string, 0)

	for m := 1; m < 13; m++ {
		for d := 1; d <= DOM[m-1]; d++ {
			if isFriday(y, m, d) {
				addHoliday(&holidays, y, m, d)
				continue
			}
			wg.Add(1)
			go func(m, d int) {
				defer wg.Done()

				if isH, err := isTimeirHoliday(y, m, d); err != nil {
					log.Fatal(err)
				} else if isH {
					addHoliday(&holidays, y, m, d)
				}
			}(m, d)
		}
		wg.Wait()
	}

	return holidays
}

func isTimeirHoliday(y, m, d int) (bool, error) {
	url := fmt.Sprintf(timeirURLFormat, y, m, d)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
	req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)

	res, err := client.Do(req)
	if err != nil {
		return false, err
	}
	if res.StatusCode != 200 {
		return false, fmt.Errorf("expected status code 200 got %d from %s", res.StatusCode, url)
	}

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return false, err
	}

	return doc.Find(".eventHoliday").Length() != 0, nil
}
