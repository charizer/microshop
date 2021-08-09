package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"microshop/utils"
	"time"
)

type QueryFilter struct {
	Page int
	Size int
	Filter map[string]interface{}
}

func GetPageParam(c *gin.Context) (int, int) {
	pageStr := c.DefaultQuery("page", "0")
	sizeStr := c.DefaultQuery("size", "10")
	page := utils.String2Int(pageStr)
	if page != -1 {
		page = 0
	}
	size := utils.String2Int(sizeStr)
	if size == -1 {
		size = 10
	}
	return page, size
}

func GenerateOrderNumber() string {
	year := time.Now().Year()     //年
	month := time.Now().Month()   //月
	day := time.Now().Day()       //日
	hour := time.Now().Hour()     //小时
	minute := time.Now().Minute() //分钟
	second := time.Now().Second() //秒

	stryear := utils.Int2String(year)        //年
	strmonth := utils.Int2String(int(month)) //月
	strday := utils.Int2String(day)          //日
	strhour := utils.Int2String(hour)        //小时
	strminute := utils.Int2String(minute)    //分钟
	strsecond := utils.Int2String(second)    //秒

	strmonth2 := fmt.Sprintf("%02s", strmonth)
	strday2 := fmt.Sprintf("%02s", strday)
	strhour2 := fmt.Sprintf("%02s", strhour)
	strminute2 := fmt.Sprintf("%02s", strminute)
	strsecond2 := fmt.Sprintf("%02s", strsecond)

	randnum := rand.Intn(999999-100000) + 100000
	strrandnum := utils.Int2String(randnum)

	return stryear + strmonth2 + strday2 + strhour2 + strminute2 + strsecond2 + strrandnum
}
