package models

import "time"

type ShopBill struct {
	Id                  int       `json:"id"`                    //自增id
	TransactionSerialNo string    `json:"transaction_serial_no"` //交易流水号
	Type                string    `json:"type"`                  //流水类型 (T:提现 J:结算)
	Money               int       `json:"money"`                 //流水金额
	Balance             int       `json:"balance"`               //余额
	GmtCreate           time.Time `json:"gmt_create"`            //创建时间
	GmtUpdate           time.Time `json:"gmt_update"`            //更新时间
}
