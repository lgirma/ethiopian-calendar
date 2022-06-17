package ethiopian_calendar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToGregorian(t *testing.T) {
	// Ordinary date
	gc := ToGregorian(NewDate(2014, 10, 8))
	assert.Equal(t, 2022, gc.Year)
	assert.Equal(t, 6, gc.Month)
	assert.Equal(t, 15, gc.Day)

	// Test Feb 28 + 1
	gc = ToGregorian(NewDate(2014, 6, 22))
	assert.Equal(t, 2022, gc.Year)
	assert.Equal(t, 3, gc.Month)
	assert.Equal(t, 1, gc.Day)

	// Before January
	gc = ToGregorian(NewDate(2014, 4, 1))
	assert.Equal(t, 2021, gc.Year)
	assert.Equal(t, 12, gc.Month)
	assert.Equal(t, 10, gc.Day)
}

func TestToEthiopian(t *testing.T) {
	// Ordinary date
	ec := ToEthiopian(NewDate(2022, 6, 15))
	assert.Equal(t, 2014, ec.Year)
	assert.Equal(t, 10, ec.Month)
	assert.Equal(t, 8, ec.Day)

	// Test Feb 28 + 1
	ec = ToEthiopian(NewDate(2022, 3, 1))
	assert.Equal(t, 2014, ec.Year)
	assert.Equal(t, 6, ec.Month)
	assert.Equal(t, 22, ec.Day)

	// Before January
	ec = ToEthiopian(NewDate(2021, 12, 10))
	assert.Equal(t, 2014, ec.Year)
	assert.Equal(t, 4, ec.Month)
	assert.Equal(t, 1, ec.Day)

	// Pagumen
	ec = ToEthiopian(NewDate(2022, 9, 6))
	assert.Equal(t, 2014, ec.Year)
	assert.Equal(t, 13, ec.Month)
	assert.Equal(t, 1, ec.Day)

}

