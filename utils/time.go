package utils

import "time"

const (
	NOW_TIME = "2006-01-02 15:04:05"
	NOW_YEAR = "2006"
	NOW_DATE = "2006-01-02"
	NOW_MONTH = "2006-01"
	NOW_MONTH_SIGN = "200601"
)

// 当前时间
func NowTime() string {
	return time.Now().Format(NOW_TIME)
}

// 当前日期
func NowDate() string {
	return time.Now().Format(NOW_DATE)
}

//当前年份
func NowYear() string {
	return time.Now().Format(NOW_YEAR)
}

func NextMonth()  {
	//t, _ := time.Now().AddDate()
}
