package timestamp

import(
"fmt"
"time"
"errors"
)

type timestamp struct{
	year int
	month int
	day int
	hour int
	minute int
	second int
	timezone int
}

func Settimestamp() {
	isnow :=time.Now()
	fmt.Println(isnow)
	
}
func changeyear(tms *timestamp, year int) (err error) {
	if year <0 || year > 10000{
		return errors.New(fmt.Sprintf("can't change year to %d",year))
	}
	tms.year = year
	return
}
func changemonth(tms *timestamp, month int) (err error) {
	if month <1 || month > 12{
		return errors.New(fmt.Sprintf("can't change month to %d",month))
	}
	tms.month = month
	return
}
func changeday(tms *timestamp, day int) (err error) {
	daylimit := 28
	switch tms.month {
		case 1,3,5,7,8,10,12:
			daylimit += 3
		case 4,6,9,11,13:
			daylimit += 2
		case 2:
			daylimit += (int(tms.year%4)^1)
	}
	if day <0 || day > daylimit {
		return errors.New(fmt.Sprintf("can't change day to %d",day))
	}
	tms.day = day
	return
}
func changehour(tms *timestamp, hour int) (err error) {
	if hour <0 || hour > 23{
		return errors.New(fmt.Sprintf("can't change hour to %d",hour))
	}
	tms.hour = hour
	return
}
func changeminute(tms *timestamp, minute int) (err error) {
	if minute <0 || minute > 59{
		return errors.New(fmt.Sprintf("can't change minute to %d",minute))
	}
	tms.minute = minute
	return
}
func changesecond(tms *timestamp, second int) (err error) {
	if second <0 || second > 59{
		return errors.New(fmt.Sprintf("can't change second to %d",second))
	}
	tms.second = second
	return
}
