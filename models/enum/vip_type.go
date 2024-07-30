package enum

import (
	"database/sql/driver"
)

type VipType uint8

const (
	VipTypeCommon   VipType = 1
	VipTypeRecharge VipType = 2
)

func (t VipType) Value() (driver.Value, error) {
	return int64(t), nil
}

func (t VipType) MarshalJSON() ([]byte, error) {
	switch t {
	case VipTypeCommon:
		return []byte("\"普通会员\""), nil
	case VipTypeRecharge:
		return []byte("\"储值会员\""), nil
	default:
		panic("会员类型错误")
	}
}
