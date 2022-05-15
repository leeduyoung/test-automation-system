package lib

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalcElapsed(t *testing.T) {
	a := time.Now().AddDate(0, 0, -10)

	days := CalcElapsed(a)
	assert.Equal(t, days, 10)
}

func TestCalcDaysBetween(t *testing.T) {
	a := time.Now()
	b := time.Now().AddDate(0, 0, 10)

	days := CalcDaysBetween(a, b)
	assert.Equal(t, days, 10)
}

func TestCalcWeeksBetween(t *testing.T) {
	a := time.Now()
	b := time.Now().AddDate(0, 0, 10)

	weeks := CalcWeeksBetween(a, b)
	assert.Equal(t, weeks, 1)
}
