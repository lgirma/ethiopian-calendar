package ethiopian_calendar

import (
	"math"
)

// gregMonthsStartingSept Number of days in gregorian months
// starting with September (index 1)
// Index 0 is reserved for leap years switches.
// Index 4 is December, the final month of the year.
var gregMonthsStartingSept = []int{0, 30, 31, 30, 31, 31, 28, 31, 30, 31, 30, 31, 31, 30}

// gregOrder Gregorian months ordered according to Ethiopian
var gregOrder = []int{8, 9, 10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// gregMonthsStartingJan Number of days in gregorian months
// starting with January (index 1)
// Index 0 is reserved for leap years switches.
var gregMonthsStartingJan = []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// ethMonths Number of days in ethiopian months
// starting with January (index 1)
// Index 0 is reserved for leap years switches.
// Index 10 is month 13, the final month of the year
var ethMonths = []int{0, 30, 30, 30, 30, 30, 30, 30, 30, 30, 5, 30, 30, 30, 30}

// ethOrder Ethiopian months ordered according to Gregorian
var ethOrder = []int{0, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 1, 2, 3, 4}


func startDayOfEc(year int) int {
	newYearDay := math.Floor(float64(year) / 100) - math.Floor(float64(year) / 400) - 4
	if (year - 1) % 4 == 3 {
		return int(newYearDay) + 1
	}
	return int(newYearDay)
}

// ToGregorian Converts given date in Ethiopian calendar to Gregorian calendar
func ToGregorian(ecDate Date) Date {
	gregorianMonths := make([]int, len(gregMonthsStartingSept))
	copy(gregorianMonths, gregMonthsStartingSept)

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
		Month: gregOrder[m],
		Year:  gregorianYear,
	}
}

// ToEthiopian Converts given date in Gregorian calendar to Ethiopian calendar
func ToEthiopian(gcDate Date) Date {
	// dates between 5 and 14 of May 1582 are invalid
	if gcDate.Month == 10 && gcDate.Day >= 5 && gcDate.Day <= 14 && gcDate.Year == 1582 {
		return Date{}
	}

	gregorianMonths := make([]int, len(gregMonthsStartingJan))
	copy(gregorianMonths, gregMonthsStartingJan)
	ethiopianMonths := make([]int, len(ethMonths))
	copy(ethiopianMonths, ethMonths)

	if (gcDate.Year %4 == 0 && gcDate.Year % 100 != 0) || gcDate.Year % 400 == 0 {
		gregorianMonths[2] = 29
	}

	ethYear := gcDate.Year - 8
	if ethYear % 4 == 3 {
		ethiopianMonths[10] = 6
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
		ethiopianMonths[1] = 0
		ethiopianMonths[2] = tahisas
	} else if until <= 277 && gcDate.Year == 1582 {
		ethiopianMonths[1] = 0
		ethiopianMonths[2] = tahisas
	} else {
		tahisas = newYearDay - 3
		ethiopianMonths[1] = tahisas
	}

	// calculate month and date incrementally
	m := 0
	ethDate := 0

	for i := 1; i < len(ethiopianMonths); i++ {
		m = i
		if until <= ethiopianMonths[m] {
			ethDate = until
			if m == 1 || ethiopianMonths[m] == 0 {
				ethDate = until + 30 - tahisas
			}
			break
		} else {
			until -= ethiopianMonths[m]
		}
	}

	// if m > 10, we're already on next Ethiopian year
	if m > 10 {
		ethYear += 1
	}

	return Date{
		Day:   ethDate,
		Month: ethOrder[m],
		Year:  ethYear,
	}
}
