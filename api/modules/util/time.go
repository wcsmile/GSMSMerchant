package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func GetNextTimeFromNow(baseTime string, period string) (next string, err error) {
	unit := period[len(period)-1 : len(period)]
	//周期间隔（0:永远,5m:5分钟，1h:小时，2d天，3w:3周，4n:4个月）
	fmtVal := "20060102150405"
	baseVal, err := time.Parse(fmtVal, baseTime)
	if err != nil {
		return
	}
	var dest time.Time
	switch unit {
	case "s", "m", "h":
		tmpval, err1 := time.ParseDuration(period)
		if err1 != nil {
			err = err1
			return
		}
		dest = baseVal.Add(tmpval)
	case "d":
		period = strings.Replace(period, "d", "", -1)
		times, err1 := strconv.Atoi(period)
		if err1 != nil {
			err = err1
			return
		}
		dest = baseVal.AddDate(0, 0, times)
	case "w":
		period = strings.Replace(period, "d", "", -1)
		times, err1 := strconv.Atoi(period)
		if err1 != nil {
			err = err1
			return
		}
		dest = baseVal.AddDate(0, 0, times*7) //一周7天
	case "n":
		period = strings.Replace(period, "n", "", -1)
		times, err1 := strconv.Atoi(period)
		if err1 != nil {
			err = err1
			return
		}
		dest = baseVal.AddDate(0, times, 0)
	default:
		dest = baseVal.AddDate(100, 0, 0)
	}
	next = dest.Format(fmtVal)
	return
}
func GetNextTimeFromBase(baseTime string, period string) (next string, err error) {
	unit := period[len(period)-1 : len(period)]
	//周期间隔（0:永远,5m:5分钟，1h:小时，2d天，3w:3周，4n:4个月）
	fmtVal := "20060102150405"
	fmtDest := ""
	var dest time.Time
	baseVal, _ := time.Parse(fmtVal, baseTime)
	switch unit {
	case "s":
		tmpval, err1 := time.ParseDuration(period)
		if err1 != nil {
			err = err1
			return
		}
		dest = baseVal.Add(tmpval)
		fmtDest = "20060102150405"
	case "m":
		tmpval, err1 := time.ParseDuration(period)
		if err1 != nil {
			err = err1
			return
		}
		dest = baseVal.Add(tmpval)
		fmtDest = "200601021504"
	case "h":
		tmpval, err1 := time.ParseDuration(period)
		if err1 != nil {
			err = err1
			return
		}
		dest = baseVal.Add(tmpval)
		fmtDest = "2006010215"
	case "d":
		period = strings.Replace(period, "d", "", -1)
		times, err1 := strconv.Atoi(period)
		if err1 != nil {
			err = err1
			return
		}
		dest = baseVal.AddDate(0, 0, times)
		fmtDest = "20060102"
	case "w":
		period = strings.Replace(period, "w", "", -1)
		times, err1 := strconv.Atoi(period)
		if err1 != nil {
			err = err1
			return
		}
		dest = baseVal.AddDate(0, 0, times*7) //一周7天
		fmtDest = "20060102"
	case "n":
		period = strings.Replace(period, "n", "", -1)
		times, err1 := strconv.Atoi(period)
		if err1 != nil {
			err = err1
			return
		}
		dest = baseVal.AddDate(0, times, 0)
		fmtDest = "200601"
	default:
		err = fmt.Errorf("不支持的类型：%s", unit)
		return
	}
	next = dest.Format(fmtDest)
	next = PadRight(next, len(fmtVal), "0")
	return
}

func GetCurrentFromBase(baseTime string, period string) (val string, err error) {
	unit := period[len(period)-1 : len(period)]
	fmtVal := "20060102150405"
	//周期间隔（0:永远,5m:5分钟，1h:小时，2d天，3w:3周，4n:4个月）
	now, _ := time.Parse(fmtVal, baseTime)
	var fmtStr string
	switch unit {
	case "m":
		fmtStr = "200601021504"
	case "h":
		fmtStr = "2006010215"
	case "d":
		fmtStr = "20060102"
	case "w":
		fmtStr = "20060102"
	case "n":
		fmtStr = "20060102"
	default:
		err = fmt.Errorf("不支持的类型：%s", unit)
		return
	}
	val = now.Format(fmtStr)
	val = PadRight(val, len(fmtVal), "0")
	return
}

func PadRight(origin string, lenval int, fillchar string) (res string) {
	if len(origin) >= lenval {
		return origin
	}
	return origin + strings.Repeat(fillchar, lenval-len(origin))
}
