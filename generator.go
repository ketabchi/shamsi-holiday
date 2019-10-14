package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mostafah/go-jalali/jalali"
)

const y = 1398

var (
	holidays   []string
	wg         sync.WaitGroup
	dateFormat = strconv.Itoa(y) + "/%02d/%02d"
	urlFormat  = "https://www.time.ir/fa/event/list/0/" + strconv.Itoa(y) + "/%d/%d"
)

func main() {
	dom := []int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}

	for m := 1; m < 13; m++ {
		for d := 1; d < dom[m-1]; d++ {
			if isFriday(m, d) {
				addHoliday(m, d)
			} else {
				go func(m, d int) {
					wg.Add(1)
					defer wg.Done()

					if isShamsiHoliday(m, d) {
						addHoliday(m, d)
					}
				}(m, d)
			}
		}
		wg.Wait()
	}

	sort.Strings(holidays)

	data, err := json.MarshalIndent(holidays, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("holidays.json", data, 0644)
}

func isFriday(m, d int) bool {
	t := jalali.Jtog(y, m, d)
	return t.Weekday() == time.Friday
}

func isShamsiHoliday(m, d int) bool {
	url := fmt.Sprintf(urlFormat, m, d)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != 200 {
		log.Fatalf("Expected status code 200 got %d from %s", res.StatusCode, url)
	}

	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatal(err)
	}

	return doc.Find(".eventHoliday").Length() != 0
}

func addHoliday(m, d int) {
	holidays = append(holidays, fmt.Sprintf(dateFormat, m, d))
}
