package error

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	ERROR_ACCESS_GOOGLE_API_FAIL:        "googleAPI access error",
	ERROR_AUTH_TOKEN_INVALID:            "invalid token error",
	ERROR_READ_GOOGLE_API_RESPONSE_FAIL: "googleAPI response read error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
