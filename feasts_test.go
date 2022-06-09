package ethiopian_calendar

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNineveh(t *testing.T) {
	assert.Equal(t, NewDayOfYear(6, 7), GetNineveh(2014))
}

func TestGetFeastFromNineveh(t *testing.T) {
	assert.Equal(t, NewDayOfYear(6, 7),
		GetFeastFromNineveh(FEAST_ABIY_TSOM, NewDayOfYear(6, 7)),
		NewDayOfYear(6, 21))
}