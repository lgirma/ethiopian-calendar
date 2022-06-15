package ethiopian_calendar

import "math"

func GetNineveh(year int) DayOfYear {
	mebajaHamer := GetMebajaHamer(year)
	return NewDay(mebajaHamer.Month + 4, mebajaHamer.Day)
}

func GetFeastFromNineveh(feast int, nineveh DayOfYear) DayOfYear {
	day := nineveh.Day + feast
	result := DayOfYear{
		Month: nineveh.Month,
		Day: day % 30,
	}
	if day % 30 == 0 {
		result.Day = 30
		result.Month += day / 30 - 1
	} else {
		result.Month += int(math.Floor(float64(day) / 30))
	}
	return result
}

func GetFeast(feast int, year int) DayOfYear {
	return GetFeastFromNineveh(feast, GetNineveh(year))
}