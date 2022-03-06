package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)

	tf := t.Format(time.Kitchen)
	fmt.Println(tf)

	/*
		which is Unix time 1136239445. Since MST is GMT-0700, the reference time can be thought of as

		01/02 03:04:05PM '06 -0700
	*/
	// 01-02-2006 --> 01/02 03:04:05PM '06 -0700
	// 01 used for month, where i want month put 01
	// 02 used for day, where i want day put 02
	// 2006 used for year, where i want year put 2006
	tf1 := t.Format("01-02-2006")
	fmt.Println(tf1)

	tf2 := t.Format("02-01-2006")
	fmt.Println(tf2)

}

// func (t Time) Format(layout string) string
// layout
/*
	const (
    ANSIC       = "Mon Jan _2 15:04:05 2006"
    UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
    RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
    RFC822      = "02 Jan 06 15:04 MST"
    RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
    RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
    RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
    RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
    RFC3339     = "2006-01-02T15:04:05Z07:00"
    RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
    Kitchen     = "3:04PM"
    // Handy time stamps.
    Stamp      = "Jan _2 15:04:05"
    StampMilli = "Jan _2 15:04:05.000"
    StampMicro = "Jan _2 15:04:05.000000"
    StampNano  = "Jan _2 15:04:05.000000000"
)
*/
