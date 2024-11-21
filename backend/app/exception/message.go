package exception

// 500
const (
	ORMError = ""
)

// signup
const (
	AccountRegisterErrorMsg   = "this account have been registered"
	AccountOrPasswordErrorMsg = "account or Password are invalid"
	ValidateCodeMsg           = "validate code is error or expire"
)

// sign in
const (
	PasswordErrorMsg = "too many times with an incorrect password, please try again later"
	PasswordErrorCode = "too_many_times_error"
)

//
const (
	WhiteListErrorMsg = "the host is not allow to access"
)
