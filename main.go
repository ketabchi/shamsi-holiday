package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/ketabchi/util"
	"github.com/mostafah/go-jalali/jalali"
)

const YEAR = 1403

var DOM = [12]int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 30}

var (
	client *http.Client

	dateFormat = "%d/%02d/%02d"
)

func main() {
	taghvimcomThisYearHds := taghvimcomHolidays()
	timeirThisYearHds := timeirHolidays(YEAR)
	timeirLastYearHds := timeirHolidays(YEAR - 1)

	diffs := diffHolidays(taghvimcomThisYearHds, timeirThisYearHds)
	for _, diff := range diffs {
		log.Printf("%s isn't in both", diff)
	}

	holidays := append(timeirLastYearHds, timeirThisYearHds...)
	sort.Strings(holidays)
	data, err := json.MarshalIndent(holidays, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("holidays.json", data, 0644)
}

func diffHolidays(hds1, hds2 []string) []string {
	diffs := make([]string, 0)

	for _, hd := range hds1 {
		if !util.SliceContains(diffs, hd) && !util.SliceContains(hds2, hd) {
			diffs = append(diffs, hd)
		}
	}
	for _, hd := range hds2 {
		if !util.SliceContains(diffs, hd) && !util.SliceContains(hds1, hd) {
			diffs = append(diffs, hd)
		}
	}

	return diffs
}

func isFriday(y, m, d int) bool {
	t := jalali.Jtog(y, m, d)
	return t.Weekday() == time.Friday
}

func toDateFormat(y, m, d int) (string, string) {
	t := jalali.Jtog(y, m, d)
	return fmt.Sprintf(dateFormat, y, m, d), t.Format("2006/01/02")
}

func addHoliday(hds *[]string, y, m, d int) {
	jd, md := toDateFormat(y, m, d)
	*hds = append(*hds, jd, md)
}

func init() {
	client = new(http.Client)
}
