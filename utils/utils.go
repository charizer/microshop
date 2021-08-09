package utils

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"strconv"
	"time"
)

func String2Int(val string) int {
	number, err := strconv.Atoi(val)
	if err != nil {
		return -1
	} else {
		return number
	}
}

func Int2String(val int) string {
	return strconv.Itoa(val)
}

func VerifyMobile(mobile string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobile)
}

func GetIntParam(c *gin.Context, name string) int {
	str := c.DefaultQuery(name,"")
	return String2Int(str)
}

func FormatTimestamp(timestamp int64, format string) string {
	tm := time.Unix(timestamp/1e3, 0)
	return tm.Format(format)
}
