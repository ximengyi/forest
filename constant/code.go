package constant



//1000以下为通用码，1000以上为用户自定义码
const (
	SuccessCode  = iota
	UndefErrorCode
	ValidErrorCode
	InternalErrorCode

	InvalidRequestErrorCode  = 401
	CustomizeCode            = 1000
	GROUPALL_SAVE_FLOWERROR  = 2001
)