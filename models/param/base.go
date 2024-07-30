package param

type BaseDetail struct {
	Id uint32 `json:"id" binding:"required"`
}
