package timestamp

import(
"fmt"
"time"
)
// TimeStp struct with date, hour and timezone
type TimeStp struct{
	Year int
	Month int
	Day int
	Hour int
	Minute int
	Second int
	Timezone int
}
// Settimestamp : It sets function's timestamp values igual to computer's timestamp.
func Settimestamp() {
	isnow :=time.Now()
	fmt.Println(isnow)
}
func changeyear(tms *TimeStp, year int) (err error) {
	if year <0 || year > 10000{
//		return errors.New(fmt.Sprintf("can't change year to %d",year))
		return fmt.Errorf("can't change year to %d",year)
	}
	tms.Year = year
	return
}
func changemonth(tms *TimeStp, month int) (err error) {
	if month <1 || month > 12{
//		return errors.New(fmt.Sprintf("can't change month to %d",month))
		return fmt.Errorf("can't change month to %d",month)
	}
	tms.Month = month
	return
}
// Changeday : Testing porpose
func Changeday(tms *TimeStp, day int) (err error) {
	daylimit := 28
	switch tms.Month {
		case 1,3,5,7,8,10,12:
			daylimit += 3
		case 4,6,9,11,13:
			daylimit += 2
		case 2:
			daylimit += (int(tms.Year%4)^1)
	}
	if day <0 || day > daylimit {
		return fmt.Errorf("can't change day to %d",day)
	}
	tms.Day = day
	return
}
func changehour(tms *TimeStp, hour int) (err error) {
	if hour <0 || hour > 23{
		return fmt.Errorf("can't change hour to %d",hour)
	}
	tms.Hour = hour
	return
}
func changeminute(tms *TimeStp, minute int) (err error) {
	if minute <0 || minute > 59{
		return fmt.Errorf("can't change minute to %d",minute)
	}
	tms.Minute = minute
	return
}
func changesecond(tms *TimeStp, second int) (err error) {
	if second <0 || second > 59{
		return fmt.Errorf("can't change second to %d",second)
	}
	tms.Second = second
	return
}
