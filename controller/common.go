package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
	"qiyu/dao"
	logKit "qiyu/logger"
	"qiyu/models"
	"qiyu/models/param"
	"qiyu/util"
	"sort"
	"strings"
	"time"
)

func CommonGetEnum(c *gin.Context) {
	util.ResponseKit(c, http.StatusOK, "请求成功", gin.H{"time_enum": TimeEnum})
}

// 传入场地信息、日期、时间标签，计算价格
func SitePriceCalculate(site models.Site, time time.Time, timeEnum int) int {
	var busyTimeEnum []int
	_ = json.Unmarshal([]byte(site.BusyTimeEnum), &busyTimeEnum)
	for _, v := range busyTimeEnum {
		if v == timeEnum {
			return site.BusyPrice
		}
	}
	//判断是不是周末
	if time.Weekday() == 0 || time.Weekday() == 6 {
		if site.WeekendBusy == 1 {
			return site.BusyPrice
		}
	}
	//判断是不是忙碌日期
	var DataBusy []string
	json.Unmarshal([]byte(site.DataBusy), &DataBusy)
	for _, v := range DataBusy {
		if v == time.Format("2006-01-02") {
			return site.BusyPrice
		}

	}

	return site.FreePrice

}

// OrderSiteMoneyCalculate 根据选择的场地进行金额计算
// 不用这个了，原来写的真垃圾，呸，稍微改改就很好了
func OrderSiteMoneyCalculate1(siteDetail []param.SiteDetail, shopId int, time time.Time) (money int, siteOrderDetail []models.SiteOrderDetail) {
	siteList, err := dao.SiteAll(shopId)
	if err != nil {
		panic(err.Error())
	}
	for _, v := range siteDetail {
	A:
		for _, vv := range siteList {
			siteMoney := 0
			if v.SiteId == vv.Id {
				//先判断特定忙碌时期是不是有今天
				var DataBusy []string
				if len(vv.DataBusy) > 3 {
					json.Unmarshal([]byte(vv.DataBusy), &DataBusy)
					format := util.Now().Format("2006-01-02")
					for _, vvv := range DataBusy {
						if vvv == format {
							money += vv.BusyPrice * len(v.TimeEnum)
							//将钱写入场地金额
							siteMoney += vv.BusyPrice * len(v.TimeEnum)
							siteOrderDetail = append(siteOrderDetail, models.SiteOrderDetail{
								SiteId:   v.SiteId,
								TimeEnum: v.TimeEnum,
								Money:    siteMoney,
							})
							logKit.Log.Println("准备跳过A循环")
							continue A
						}
						logKit.Log.Println("如果你看见我了，代表上面的没有跳过")
					}
				}
				//如果是周末，就按照忙时价格算
				if time.Weekday() == 0 || time.Weekday() == 6 {
					//这是周末
					money += vv.BusyPrice * len(v.TimeEnum)
					siteMoney += vv.BusyPrice * len(v.TimeEnum)
					siteOrderDetail = append(siteOrderDetail, models.SiteOrderDetail{
						SiteId:   v.SiteId,
						TimeEnum: v.TimeEnum,
						Money:    siteMoney,
					})
					continue A
				}
				//找对对应的场地，然后看这个场地有多少个场次
				var busyTimeEnum []int
				json.Unmarshal([]byte(vv.BusyTimeEnum), &busyTimeEnum)
			B:
				for _, vvv := range v.TimeEnum {
					for _, busyTime := range busyTimeEnum {
						if vvv == busyTime {
							money += vv.BusyPrice
							siteMoney += vv.BusyPrice
							continue B
						}
					}
					money += vv.FreePrice
					siteMoney += vv.FreePrice
				}
			}
			siteOrderDetail = append(siteOrderDetail, models.SiteOrderDetail{
				SiteId:   v.SiteId,
				TimeEnum: v.TimeEnum,
				Money:    siteMoney,
			})
		}
	}
	return money, siteOrderDetail
}
func OrderSiteMoneyCalculate(siteDetail []param.SiteDetail, time time.Time) (totalMoney int, siteOrderDetail []models.SiteOrderDetail) {
	for _, v := range siteDetail {
		siteMoney := 0
		site, err := dao.SiteGet(v.SiteId)
		if err != nil {
			panic(err.Error())
		}
		for _, vv := range v.TimeEnum {
			calculate := SitePriceCalculate(*site, time, vv)
			totalMoney += calculate
			siteMoney += calculate
		}
		siteOrderDetail = append(siteOrderDetail, models.SiteOrderDetail{
			SiteName: site.Name,
			SiteId:   v.SiteId,
			TimeEnum: v.TimeEnum,
			Money:    siteMoney,
		})
	}
	return totalMoney, siteOrderDetail
}

// GetWorkTimeEnum 根据开始时间和结束时间获取对应的时间枚举
func GetWorkTimeEnum(startTime string, endTime string) ([]int, error) {
	var enum []int
	//toTime, err := util.StringToTime("2006-01-02 15:01", startTime)
	//fmt.Println("领导看见", toTime, err)
	startTimeIndex := 0
	endTimeIndex := 0
	for k, v := range TimeEnum {
		split := strings.Split(v, "~")

		if strings.Contains(repairTime(startTime), repairTime(split[0])) {
			startTimeIndex = k
		}
		if strings.Contains(repairTime(endTime), repairTime(split[1])) {
			endTimeIndex = k
		}
	}
	for k, _ := range TimeEnum {
		if k >= startTimeIndex && k <= endTimeIndex {
			enum = append(enum, k)
		}
	}
	if len(enum) == 0 {
		return nil, errors.New("有错误")
	}
	sort.Ints(enum)
	return enum, nil
}
func repairTime(time string) string {
	if len(time) == 4 {
		return "0" + time
	}
	return time

}
