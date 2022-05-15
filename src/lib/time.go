package lib

import "time"

// CalcElapsed 입력받은날부터 오늘까지 경과된 일수
func CalcElapsed(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

// CalcDaysBetween 두 날짜 사이의 일자 계산
func CalcDaysBetween(t1 time.Time, t2 time.Time) int {
	return int(t2.Sub(t1).Hours() / 24)
}

// CalcWeeksBetween 두 날짜 사이의 주 계산
func CalcWeeksBetween(t1 time.Time, t2 time.Time) int {
	return int(t2.Sub(t1).Hours() / (24 * 7))
}
