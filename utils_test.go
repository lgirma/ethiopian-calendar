package ethiopian_calendar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEvangelist(t *testing.T) {
	assert.Equal(t, EVANGELIST_MATTHEW, GetEvangelist(2001))
	assert.Equal(t, EVANGELIST_JOHN, GetEvangelist(2000))
	assert.Equal(t, EVANGELIST_LUKE, GetEvangelist(1999))
	assert.Equal(t, EVANGELIST_MARK, GetEvangelist(1998))
}

func TestGetRabit(t *testing.T) {
	r := GetRabit(2001)
	assert.Equal(t, 1875, r)
}

func TestGetWenber(t *testing.T) {
	r := GetWenber(2001)
	assert.Equal(t, 14, r)
}

func TestGetMeTq(t *testing.T) {
	r := GetMeTq(2001)
	assert.Equal(t, 26, r)
}

func TestGetMebacha(t *testing.T) {
	r := GetMebacha(2010)
	assert.Equal(t, DAY_MON, r)

	r = GetMebacha(2011)
	assert.Equal(t, DAY_TUE, r)

	r = GetMebacha(2012)
	assert.Equal(t, DAY_THU, r)

	r = GetMebacha(2013)
	assert.Equal(t, DAY_FRI, r)
}
