package harvest

import (
	"encoding/json"
	"regexp"
	"strings"
	"time"
)

type Date struct {
	time.Time
}

var shortDateExp = regexp.MustCompile("^[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]$")

func (harvestDate *Date) UnmarshalJSON(input []byte) (err error) {
	dateString := string(input)
	dateString = strings.Trim(dateString, "\"")
	if shortDateExp.MatchString(dateString) {
		harvestDate.Time, err = time.Parse("2006-01-02", dateString)
	} else if dateString != "null" {
		err = json.Unmarshal(input, &harvestDate.Time)
	}
	return err
}
