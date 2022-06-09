package ethiopian_calendar

import "math"

func GetEvangelist(year int) int {
	aa := 5500 + year
	result := aa % 4
	if result == 0 {
		return EVANGELIST_JOHN
	}
	return result
}

func GetRabit(year int) int {
	aa := 5500 + year
	return int(math.Floor(float64(aa) / 4))
}

func GetWenber(year int) int {
	aa := 5500 + year
	return (aa % 19) - 1
}

func GetMeTq(year int) int {
	return (GetWenber(year) * 19) % 30
}

func GetMebacha(year int) int {
	aa := 5500 + year
	rem := (aa + GetRabit(year)) % 7
	return rem + 1
}

func GetMebajaHamer(year int) DayOfYear {
	mebacha := GetMebacha(year)
	meTq := GetMeTq(year)

	meTqMonth := MONTH_MESKEREM
	if meTq >= 2 && meTq <= 14 {
		meTqMonth = MONTH_TIKIMT
	}
	meTqDay := ((meTqMonth-1) * 30 + mebacha + meTq - 1) % 7
	if meTqDay == 0 {
		meTqDay = 7
	}
	tewsakTable := []int{0, 6, 5, 4, 3, 2, 8, 7}
	meTqAndDayTewsak := meTq + tewsakTable[meTqDay]
	month := 0
	if meTqAndDayTewsak > 30 {
		month = 1
	}
	return DayOfYear{
		Month: month,
		Day: meTqAndDayTewsak % 30,
	}
}