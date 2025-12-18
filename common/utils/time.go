package utils

import "time"

const (
	DateFormat     = "2006-01-02"
	TimeFormat     = "15:04:05"
	DateTimeFormat = "2006-01-02 15:04:05"
)

// FormatDate 格式化日期
func FormatDate(t time.Time) string {
	return t.Format(DateFormat)
}

// FormatTime 格式化时间
func FormatTime(t time.Time) string {
	return t.Format(TimeFormat)
}

// FormatDateTime 格式化日期时间
func FormatDateTime(t time.Time) string {
	return t.Format(DateTimeFormat)
}

// ParseDate 解析日期
func ParseDate(s string) (time.Time, error) {
	return time.ParseInLocation(DateFormat, s, time.Local)
}

// ParseDateTime 解析日期时间
func ParseDateTime(s string) (time.Time, error) {
	return time.ParseInLocation(DateTimeFormat, s, time.Local)
}

// GetStartOfDay 获取当天开始时间
func GetStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// GetEndOfDay 获取当天结束时间
func GetEndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
}

// DaysBetween 计算两个日期之间的天数
func DaysBetween(start, end time.Time) int {
	return int(end.Sub(start).Hours() / 24)
}
