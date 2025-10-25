package response

const (
	ErrCodeSuccess      = 1000 // Success
	ErrCodeParamInvalid = 1001 // Email invalid
)

var msg = map[int]string{
	ErrCodeSuccess:      "success",
	ErrCodeParamInvalid: "Email invalid",
}
