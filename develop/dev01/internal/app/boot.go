package app

import (
	"d-alejandro/training-level2/develop/dev01/internal/helpers"
	"fmt"
	"os"
	"time"
)

/*
Boot is the project loading function
*/
func Boot() {
	ntpAdapter := helpers.NewNTPAdapter()

	const NTPServerAddress = "0.beevik-ntp.pool.ntp.org"

	currentTime := time.Now()

	ntpTime, err := ntpAdapter.GetCurrentTime(NTPServerAddress)

	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	const TimeLayout = "2006-01-02 15:04:05.000000000"

	currentTimeFormatted := currentTime.Format(TimeLayout)
	ntpTimeFormatted := ntpTime.Format(TimeLayout)

	fmt.Println(" Текущее время:\n", currentTimeFormatted, "\n /\n Точное время NTP:\n", ntpTimeFormatted)
}

/*
 Текущее время:
 2024-05-25 19:36:02.289271367
 /
 Точное время NTP:
 2024-05-25 19:36:02.523600047
*/
