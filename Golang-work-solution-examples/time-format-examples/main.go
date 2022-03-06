package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type UserTimeZone struct {
	DSTOffset    int    `json:"dstOffset"`
	RawOffset    int    `json:"rawOffset"`
	Status       string `json:"status"`
	TimeZoneID   string `json:"timeZoneId"`
	TimeZoneName string `json:"timeZoneName"`
}

func main() {
	// var createDate int64 = 1598346081269
	var createDate int64 = 1638997085000
	timezone := UserTimeZone{} 
	// timezone := {"dstOffset":3600,"rawOffset":-18000,"status":"OK","timeZoneId":"America/New_York","timeZoneName":"Eastern Daylight Time"}
	// timezone = UserTimeZone{
	// 	DSTOffset:    3600,
	// 	RawOffset:    -18000,    
	// 	Status:       "OK",
	// 	TimeZoneID:   "America/New_York",
	// 	TimeZoneName: "Eastern Daylight Time",
	// }
	sb, _ := json.Marshal(timezone)
	getTimeZoneAndFormatDate(createDate, string(sb))
}

func getTimeZoneAndFormatDate(createDate int64, timezone string) string {
	log.Println("getTimeZoneAndFormatDate func createDate ==>", createDate)
	tzID := ""
	var userTimeZone UserTimeZone
	if timezone != "" {
		err := json.Unmarshal([]byte(timezone), &userTimeZone)
		if err != nil {
			return err.Error()
		}
		tzID = userTimeZone.TimeZoneID

	}
	log.Println("TimeZone of user is ", tzID)
	return getFormattedDate(createDate, tzID)
}

func getFormattedDate(createDate int64, timeZoneID string) string {
	log.Println("getFormattedDate func. createDate ==>", createDate)
	ts := msToTime(createDate)
	fmt.Println(ts)
	log.Println("getFormattedDate func. ts ==>", ts.Format("02 January, 2006 15:04:05"))
	loc, _ := time.LoadLocation(timeZoneID)
	log.Println("getFormattedDate func. loc ==>", loc)
	timeStamp := ts.In(loc)
	log.Println("getFormattedDate func. timeStamp ==>", timeStamp)
	fmtdate := timeStamp.Format("02 Jan, 2006 15:04 PM") + " " + fmt.Sprint(timeStamp.Zone())
	log.Println("getFormattedDate func. fmtdate ==>", fmtdate)
	log.Println("-----------------------------------------------")
	zoneName, offset := timeStamp.Zone()
	log.Println("zonename:", zoneName, "offset:", offset)
	fmtdate1 := timeStamp.Format("02 Jan, 2006 03:04 PM") + " " + zoneName
	log.Println("getFormattedDate func. fmtdate ==>", fmtdate1)
	return fmtdate
}

func msToTime(ms int64) time.Time {
	return time.Unix(0, ms*int64(time.Millisecond))
}
