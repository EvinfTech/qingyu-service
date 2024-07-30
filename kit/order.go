package kit

import (
	"fmt"
	"os"
	"qiyu/dao"
	"qiyu/models"
	"qiyu/util"
	"sync/atomic"
	"time"
)

var db = dao.DB
var num int64

// GenerateOrderNo 获取订单号
// 自动生成订单号，根据用户id
func GenerateOrderNo(UserId int) string {
	t := util.Now()
	s := util.NewTime().Format(t, "YYYYMMDDHHmmss")
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := Sup(m, 3)
	p := os.Getpid() % 1000
	ps := Sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := Sup(r, 4)
	toInt64, _ := util.ToInt64(UserId)
	uid := Sup(toInt64, 5)
	n := fmt.Sprintf("%s%s%s%s%s", s, ms, ps, rs, uid)
	return n
}

// 对长度不足n的数字前面补0
func Sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}

func AutoSuperviseMoney() {
	for {
		lock.Lock()
		var orderList []models.Order
		db.Where("supervise_status = 'N' AND status = 'Y' AND gmt_create < ?", util.Now().AddDate(0, 0, -1)).Find(&orderList)
		if len(orderList) != 0 {
			var money int
			for _, v := range orderList {
				money += v.Money
				v.SuperviseStatus = "Y"
				v.GmtUpdate = util.Now()
				db.Updates(&v)

			}
			shop := models.Shop{}
			db.First(&shop)
			shop.Balance += money
			db.Updates(&shop)
			db.Create(&models.ShopBill{
				TransactionSerialNo: GenerateOrderNo(1),
				Type:                "J",
				Money:               money,
				Balance:             shop.Balance,
				GmtCreate:           util.Now(),
				GmtUpdate:           util.Now(),
			})
			lock.Unlock()
		}

		time.Sleep(time.Hour * 12)
	}

}
