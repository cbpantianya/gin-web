package errx

const (
	Success        = 0
	UnknownError   = -1
	NetWorkError   = 40400
	MySqlError     = 10000
	RedisError     = 20000
	GatewayDeny    = 40301
	FrequencyError = 40302
	QueryError     = 42200
)

var errMX = map[int]string{
	Success:        "成功",
	NetWorkError:   "网络错误",
	UnknownError:   "未知错误",
	MySqlError:     "数据库错误",
	RedisError:     "缓存错误",
	QueryError:     "请求参数错误",
	GatewayDeny:    "网关拒绝了你的请求",
	FrequencyError: "请求频率过高",
}

func GetErrMsg(code int) string {
	if msg, ok := errMX[code]; ok {
		return msg
	}
	return errMX[UnknownError]
}
