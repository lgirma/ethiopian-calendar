package ethiopian_calendar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNineveh(t *testing.T) {
	assert.Equal(t, NewDay(6, 7), GetNineveh(2014))
}

func TestGetFeastFromNineveh(t *testing.T) {
	runTestForYear(t, 2014)
	runTestForYear(t, 2015)
	runTestForYear(t, 2022)
	runTestForYear(t, 2024)
}

var preComputedTableFrom2014 = [][]DayOfYear {
	// Nenewe, Abiy Tsom, Debrezeit, Hosaena, Siqlet, Tinsae, Erget, PeraqliTos, Tsome Hawariat
	{
		NewDay(6, 7), NewDay(6, 21), NewDay(7, 18), NewDay(8, 9),
		NewDay(8, 14), NewDay(8, 16), NewDay(9, 25), NewDay(10, 5),
		NewDay(10, 6),
	},
	{
		NewDay(5, 29), NewDay(6, 13), NewDay(7, 10), NewDay(8,1),
		NewDay(8, 6), NewDay(8, 8), NewDay(9, 17), NewDay(9, 27),
		NewDay(9, 28),
	},
	{}, {}, {}, {}, {}, {},
	{
		NewDay(6, 11), NewDay(6, 25), NewDay(7, 22), NewDay(8, 13),
		NewDay(8, 18), NewDay(8, 20), NewDay(9, 29), NewDay(10, 9),
		NewDay(10, 10),
	},
	{},
	{
		NewDay(6, 15), NewDay(6, 29), NewDay(7, 26), NewDay(8, 17),
		NewDay(8, 22), NewDay(8, 24), NewDay(10, 13), NewDay(10, 14),
		NewDay(10, 10),
	},
}

func runTestForYear(t *testing.T, year int) {
	nv := GetNineveh(year)
	i := year - 2014
	assert.Equal(t, preComputedTableFrom2014[i][0], nv)
	assert.Equal(t, preComputedTableFrom2014[i][1], GetFeastFromNineveh(FEAST_ABIY_TSOM, nv))
	assert.Equal(t, preComputedTableFrom2014[i][2], GetFeastFromNineveh(FEAST_DEBREZEIT, nv))
	assert.Equal(t, preComputedTableFrom2014[i][3], GetFeastFromNineveh(FEAST_HOSAENA, nv))
	assert.Equal(t, preComputedTableFrom2014[i][4], GetFeastFromNineveh(FEAST_SIQLET, nv))
	assert.Equal(t, preComputedTableFrom2014[i][5], GetFeastFromNineveh(FEAST_TINSAE, nv))
	assert.Equal(t, preComputedTableFrom2014[i][6], GetFeastFromNineveh(FEAST_ERGET, nv))
	assert.Equal(t, preComputedTableFrom2014[i][7], GetFeastFromNineveh(FEAST_PERAQLITOS, nv))
}