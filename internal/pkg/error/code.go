package error

const (
	SUCCESS = 200
	ERROR   = 400

	ERROR_ACCESS_GOOGLE_API_FAIL        = 10001
	ERROR_AUTH_TOKEN_INVALID            = 10002
	ERROR_READ_GOOGLE_API_RESPONSE_FAIL = 10003

	ERROR_UPLOAD_SAVE_IMAGE_FAIL    = 30001
	ERROR_UPLOAD_CHECK_IMAGE_FAIL   = 30002
	ERROR_UPLOAD_CHECK_IMAGE_FORMAT = 30003
)
