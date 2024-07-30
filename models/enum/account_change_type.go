package enum

import (
	"database/sql/driver"
)

type AccountChangeType uint8

const (
	AccountChangeTypeAdminRecharge AccountChangeType = 1
	AccountChangeTypeUserRecharge  AccountChangeType = 2
	AccountChangeTypePay           AccountChangeType = 3
	AccountChangeTypePayRefund     AccountChangeType = 4
)

func (t AccountChangeType) Value() (driver.Value, error) {
	return int64(t), nil
}

func (t AccountChangeType) MarshalJSON() ([]byte, error) {
	switch t {
	case AccountChangeTypeAdminRecharge:
		return []byte("\"后台充值\""), nil
	case AccountChangeTypeUserRecharge:
		return []byte("\"用户充值\""), nil
	case AccountChangeTypePay:
		return []byte("\"储值支付\""), nil
	case AccountChangeTypePayRefund:
		return []byte("\"储值支付退款\""), nil
	}
	panic("账户变更类型错误")
}
