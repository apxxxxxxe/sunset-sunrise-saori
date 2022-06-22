package main

import (
	"fmt"
	"github.com/kelvins/sunrisesunset"
	"os"
	"time"
)

const format = "20060102"

func main() {

	if len(os.Args) < 2 {
		fmt.Println("error: not enough args")
		return
	}

	t, err := time.Parse(format, os.Args[1])
	if err != nil {
		fmt.Println("error: args should format like '20060102' (YYYYMMDD)")
		return
	}

	// 経度緯度は兵庫県明石市を基準
	p := sunrisesunset.Parameters{
		Latitude:  35.000000,
		Longitude: 135.000000,
		UtcOffset: 9.0,
		Date:      t,
	}

	// Calculate the sunrise and sunset times
	sunrise, sunset, err := p.GetSunriseSunset()

	// If no error has occurred, print the results
	if err == nil {
		fmt.Println(sunrise.Format("15,04,05") + ";" + sunset.Format("15,04,05"))
		fmt.Println(sunrise.Format("15,04,05") + "," + sunset.Format("15,04,05"))
	} else {
		fmt.Println("error: failed to get sunrise/sunset time")
	}
}
