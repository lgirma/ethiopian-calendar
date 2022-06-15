package ethiopian_calendar

import (
	"math"
)

// GREG_MONTHS_STARTING_SEPT Number of days in gregorian months
// starting with September (index 1)
// Index 0 is reserved for leap years switches.
// Index 4 is December, the final month of the year.
var GREG_MONTHS_STARTING_SEPT = []int{0, 30, 31, 30, 31, 31, 28, 31, 30, 31, 30, 31, 31, 30}

// GREG_ORDER Gregorian months ordered according to Ethiopian
var GREG_ORDER = []int{8, 9, 10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// GREG_MONTHS_STARTING_JAN Number of days in gregorian months
// starting with January (index 1)
// Index 0 is reserved for leap years switches.
var GREG_MONTHS_STARTING_JAN = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// ETH_MONTHS Number of days in ethiopian months
// starting with January (index 1)
// Index 0 is reserved for leap years switches.
// Index 10 is month 13, the final month of the year
var ETH_MONTHS = []int{0, 30, 30, 30, 30, 30, 30, 30, 30, 30, 5, 30, 30, 30, 30}

// ETH_ORDER Ethiopian months ordered according to Gregorian
var ETH_ORDER = []int{0, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 1, 2, 3, 4}



func startDayOfEc(year int) int {
	newYearDay := math.Floor(float64(year) / 100) - math.Floor(float64(year) / 400) - 4
	if (year - 1) % 4 == 3 {
		return int(newYearDay) + 1
	}
	return int(newYearDay)
}

func ToGregorian(ecDate Date) Date {
	gregorianMonths := make([]int, len(GREG_MONTHS_STARTING_SEPT))
	copy(gregorianMonths, GREG_MONTHS_STARTING_SEPT)

	newYearDay := startDayOfEc(ecDate.Year)
	gregorianYear := ecDate.Year + 7

	nextYear := gregorianYear + 1
	if (nextYear %4 == 0 && nextYear % 100 != 0) || nextYear % 400 == 0 {
		gregorianMonths[6] = 29
	}

	// calculate number of days up to that date
	until := ((ecDate.Month - 1) * 30) + ecDate.Day
	if until <=  37 && ecDate.Year <= 1575 {
		until += 28
		gregorianMonths[0] = 31
	} else {
		until += newYearDay - 1
	}

	// if ethiopian year is leap year, pagumen has six days
	if ecDate.Year - 1 % 4 == 3 {
		until += 1
	}

	// calculate month and date incrementally
	m := 0
	gregorianDate := 0
	for i := 0; i < len(gregorianMonths); i++ {
		if until <= gregorianMonths[i] {
			m = i
			gregorianDate = until
			break
		} else {
			m = i
			until -= gregorianMonths[i]
		}
	}

	// if m > 4, we're already on next Gregorian year
	if m > 4 {
		gregorianYear += 1
	}

	return Date{
		Day:   gregorianDate,
		Month: GREG_ORDER[m],
		Year:  gregorianYear,
	}
}

func ToEthiopian(gcDate Date) Date {
	// dates between 5 and 14 of May 1582 are invalid
	if gcDate.Month == 10 && gcDate.Day >= 5 && gcDate.Day <= 14 && gcDate.Year == 1582 {
		return Date{}
	}

	gregorianMonths := make([]int, len(GREG_MONTHS_STARTING_JAN))
	copy(gregorianMonths, GREG_MONTHS_STARTING_JAN)
	ethMonths := make([]int, len(ETH_MONTHS))
	copy(ethMonths, ETH_MONTHS)

	if (gcDate.Year %4 == 0 && gcDate.Year % 100 != 0) || gcDate.Year % 400 == 0 {
		gregorianMonths[6] = 29
	}

	ethYear := gcDate.Year - 8
	if ethYear % 4 == 3 {
		ethMonths[10] = 6
	}

	// Ethiopian new year in Gregorian calendar
	newYearDay := startDayOfEc(gcDate.Year - 8)

	// calculate number of days up to that date
	until := 0
	for i := 1; i < gcDate.Month; i++ {
		until += gregorianMonths[i]
	}
	until += gcDate.Day

	// update Tahissas (December) to match January 1st
	tahisas := 25
	if ethYear % 4 == 0 {
		tahisas = 26
	}

	// take into account the 1582 change
	if gcDate.Year < 1582 {
		ethMonths[1] = 0
		ethMonths[2] = tahisas
	} else if until <= 277 && gcDate.Year == 1582 {
		ethMonths[1] = 0
		ethMonths[2] = tahisas
	} else {
		tahisas = newYearDay - 3
		ethMonths[1] = tahisas
	}

	// calculate month and date incrementally
	m := 0
	ethDate := 0

	for i := 1; i < len(ethMonths); i++ {
		m = i
		if until <= ethMonths[m] {
			ethDate = until
			if m == 1 || ethMonths[m] == 0 {
				ethDate = until + 30 - tahisas
			}
			break
		} else {
			until -= ethMonths[m]
		}
	}

	// if m > 10, we're already on next Ethiopian year
	if m > 10 {
		ethYear += 1
	}

	return Date{
		Day:   ethDate,
		Month: ETH_ORDER[m],
		Year:  ethYear,
	}
}
