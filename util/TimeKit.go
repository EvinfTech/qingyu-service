package util

/**
此工具类是用来做时间使用
Author：陈兆年
data：:2021年7月6日10:20:02
*/
import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// 常量
const (
	RFC3339 = "2006-01-02T15:04:05+08:00"
	TT      = "2006-01-02 15:04:05"
	YMD     = "2006-01-02"
	HMS     = "15:04:05"
)

func Time2String(t time.Time) string {
	return t.Format(TT)
}

func Date2String(t time.Time) string {
	return t.Format(YMD)
}

// 时区
var Location = time.FixedZone("Asia/Shanghai", 8*60*60)

type GoTime struct {
	Location time.Location
}

func Now() time.Time {
	return NewTime().NowTime()
}

// 实例
func NewTime() *GoTime {
	return &GoTime{}
}

// 获取当前时间戳
func (gt *GoTime) NowUnix() int64 {
	return gt.NowTime().Unix()
}

// 获取当前时间Time
func (gt *GoTime) NowTime() time.Time {
	return time.Now().In(Location)
}

// 获取年月日
func (gt *GoTime) GetYmd() string {
	return gt.NowTime().Format(YMD)
}

// 获取时分秒
func (gt *GoTime) GetHms() string {
	return gt.NowTime().Format(HMS)
}

// 获取当天的开始时间, eg: 2018-01-01 00:00:00
func (gt *GoTime) NowStart() string {
	now := gt.NowTime()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, Location)
	return tm.Format(TT)
}

// 获取当天的结束时间, eg: 2018-01-01 23:59:59
func (gt *GoTime) NowEnd() string {
	now := gt.NowTime()
	tm := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 1e9-1, Location)
	return tm.Format(TT)
}

// 当前时间 减去 多少秒
func (gt *GoTime) Before(beforeSecond int64) string {
	return time.Unix(gt.NowUnix()-beforeSecond, 0).Format(TT)
}

// 当前时间 加上 多少秒
func (gt *GoTime) Next(beforeSecond int64) string {
	return time.Unix(gt.NowUnix()+beforeSecond, 0).Format(TT)
}

// 2006-01-02T15:04:05Z07:00 转 时间戳
func (gt *GoTime) RfcToUnix(layout string) int64 { //转化所需模板
	tm, err := time.ParseInLocation(time.RFC3339, layout, Location) //使用模板在对应时区转化为time.time类型
	if err != nil {
		return int64(0)
	}
	return tm.Unix()
}

// 2006-01-02 15:04:05 转 时间戳
func (gt *GoTime) ToUnix(layout string) int64 {
	theTime, _ := time.ParseInLocation(TT, layout, Location)
	return theTime.Unix()
}

// 获取RFC3339格式
func (gt *GoTime) GetRFC3339() string {
	return gt.NowTime().Format(time.RFC3339)
}

// 转换成RFC3339格式
func (gt *GoTime) ToRFC3339(layout string) string {
	tm, err := time.ParseInLocation(TT, layout, Location)
	if err != nil {
		return ""
	}
	return tm.Format(time.RFC3339)
}

// Format time.Time struct to string
// MM - month - 01
// M - month - 1, single bit
// DD - day - 02
// D - day 2
// YYYY - year - 2006
// YY - year - 06
// HH - 24 hours - 03
// H - 24 hours - 3
// hh - 12 hours - 03
// h - 12 hours - 3
// mm - minute - 04
// m - minute - 4
// ss - second - 05
// s - second = 5
func (gt *GoTime) Format(t time.Time, format string) string {
	res := strings.Replace(format, "MM", t.Format("01"), -1)
	res = strings.Replace(res, "M", t.Format("1"), -1)
	res = strings.Replace(res, "DD", t.Format("02"), -1)
	res = strings.Replace(res, "D", t.Format("2"), -1)
	res = strings.Replace(res, "YYYY", t.Format("2006"), -1)
	res = strings.Replace(res, "YY", t.Format("06"), -1)
	res = strings.Replace(res, "HH", fmt.Sprintf("%02d", t.Hour()), -1)
	res = strings.Replace(res, "H", fmt.Sprintf("%d", t.Hour()), -1)
	res = strings.Replace(res, "hh", t.Format("03"), -1)
	res = strings.Replace(res, "h", t.Format("3"), -1)
	res = strings.Replace(res, "mm", t.Format("04"), -1)
	res = strings.Replace(res, "m", t.Format("4"), -1)
	res = strings.Replace(res, "ss", t.Format("05"), -1)
	res = strings.Replace(res, "s", t.Format("5"), -1)
	return res
}

// layout 是转成的格式，str是string类型的时间  	time, err := util.StringToTime("2006-01-02", param.GmtSiteUse)
func StringToTime(layout string, str string) (time.Time, error) {
	t, err := time.ParseInLocation(layout, str, Location)
	if nil == err && !t.IsZero() {
		return t, nil
	}
	return t, errors.New("格式不正确")
}

// 获取传入的时间所在月份的第一天，即某月第一天的0点。如传入time.Now(), 返回当前月份的第一天0点时间。
func GetFirstDateOfMonth(d time.Time) time.Time {
	d = d.AddDate(0, 0, -d.Day()+1)
	return GetZeroTime(d)
}

// 获取传入的时间所在月份的最后一天，即某月最后一天的0点。如传入time.Now(), 返回当前月份的最后一天0点时间。
func GetLastDateOfMonth(d time.Time) time.Time {
	return GetFirstDateOfMonth(d).AddDate(0, 1, -1)
}

// 获取某一天的0点时间
func GetZeroTime(d time.Time) time.Time {
	return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
}
func GetUpdateDate() int {
	return StringToInt(Now().Format("20060102"))
}
func GetUpdateTime() int {
	return int(Now().Unix())
}

// 获取两个时间差几天
func TimeSubDays(t1, t2 time.Time) int {
	//
	//if t1.Location().String() != t2.Location().String() {
	//  return -1
	//}
	hours := t1.Sub(t2).Hours()

	if hours <= 0 {
		return -1
	}
	// sub hours less than 24
	if hours < 24 {
		// may same day
		t1y, t1m, t1d := t1.Date()
		t2y, t2m, t2d := t2.Date()
		isSameDay := (t1y == t2y && t1m == t2m && t1d == t2d)

		if isSameDay {

			return 0
		} else {
			return 1
		}

	} else { // equal or more than 24

		if (hours/24)-float64(int(hours/24)) == 0 { // just 24's times
			return int(hours / 24)
		} else { // more than 24 hours
			return int(hours/24) + 1
		}
	}

}

// GetFirstDateOfWeek 获取本周周一的日期
func GetFirstDateOfWeek(t time.Time) time.Time {

	offset := int(time.Monday - t.Weekday())
	if offset > 0 {
		offset = -6
	}

	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local).
		AddDate(0, 0, offset)
}

// GetLastDateOfWeek 获取本周周日
func GetLastDateOfWeek(t time.Time) time.Time {

	return GetFirstDateOfWeek(t).
		AddDate(0, 0, 6)

}

// GetNextFirstDateOfWeek 获取下周周一
func GetNextFirstDateOfWeek(t time.Time) time.Time {

	return GetFirstDateOfWeek(t).
		AddDate(0, 0, 7)

}

// GetLastWeekFirstDate 获取上周的周一日期
func GetLastWeekFirstDate(t time.Time) time.Time {
	thisWeekMonday := GetFirstDateOfWeek(t)

	return thisWeekMonday.AddDate(0, 0, -7)
}

// GetLastWeekLastDate 获取下周周日
func GetLastWeekLastDate(t time.Time) time.Time {
	return GetLastDateOfWeek(t).AddDate(0, 0, -7)
}
