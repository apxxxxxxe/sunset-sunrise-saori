package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kelvins/sunrisesunset"
)

const version = "0.5.1"
const format = "20060102"

func main() {

	// 経度緯度の既定値は兵庫県明石市
	latitude := 35.000
	longitude := 135.000
	utcOffset := 9.0

	if len(os.Args) < 2 {
		fmt.Println("error: not enough args")
		return
	}

	if len(os.Args) >= 4 {
		var l float64
		var err error
		l, err = strconv.ParseFloat(os.Args[2], 3)
		if err == nil {
			latitude = l
		}
		l, err = strconv.ParseFloat(os.Args[3], 3)
		if err == nil {
			longitude = l
		}
	}

	if len(os.Args) >= 5 {
		u, err := strconv.ParseFloat(os.Args[4], 3)
		if err == nil {
			utcOffset = u
		}
	}

	t, err := time.Parse(format, os.Args[1])
	if err != nil {
		fmt.Println("error: args should format like '20060102' (YYYYMMDD)")
		return
	}

	p := sunrisesunset.Parameters{
		Latitude:  latitude,
		Longitude: longitude,
		UtcOffset: utcOffset,
		Date:      t,
	}

	// Calculate the sunrise and sunset times
	sunrise, sunset, err := p.GetSunriseSunset()

	// If no error has occurred, print the results
	if err == nil {
		fmt.Println(sunrise.Format("15,04,05") + "," + sunset.Format("15,04,05"))
	} else {
		fmt.Println("error: failed to get sunrise/sunset time")
	}
}
