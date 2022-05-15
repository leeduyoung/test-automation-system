package main

import (
	"fmt"
	"time"

	"test-automation-system/src/lib"
)

func main() {
	fmt.Println("===================================")
	fmt.Println("==== 테스트 자동화 시스템 구축 ====")
	fmt.Println("=== CODECLIMATE + GITHUB ACTION ===")
	fmt.Println("===================================")

	a := time.Now()
	b := time.Now().AddDate(0, 0, 10)

	days := lib.CalcDaysBetween(a, b)
	fmt.Println("first day:", a)
	fmt.Println("second day:", b)
	fmt.Println("days:", days)
}
