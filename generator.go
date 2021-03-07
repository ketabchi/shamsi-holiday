package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/ketabchi/util"
	"github.com/mostafah/go-jalali/jalali"
)

const YEAR = 1399

var DOM = [12]int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 30}

var (
	client *http.Client

	dateFormat = strconv.Itoa(YEAR) + "/%02d/%02d"
)

func main() {
	taghvimcomHds := taghvimcomHolidays()
	timeirHds := timeirHolidays()

	diffs := diffHolidays(taghvimcomHds, timeirHds)
	for _, diff := range diffs {
		log.Printf("%s isn't in both", diff)
	}

	sort.Strings(timeirHds)

	data, err := json.MarshalIndent(timeirHds, "", "	")
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

func isFriday(m, d int) bool {
	t := jalali.Jtog(YEAR, m, d)
	return t.Weekday() == time.Friday
}

func toDateFormat(m, d int) (string, string) {
	t := jalali.Jtog(YEAR, m, d)
	return fmt.Sprintf(dateFormat, m, d), t.Format("2006/01/02")
}

func addHoliday(hds *[]string, m, d int) {
	jd, md := toDateFormat(m, d)
	*hds = append(*hds, jd, md)
}

func init() {
	client = new(http.Client)
}
