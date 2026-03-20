package email

import "github.com/mlinder10/wcj/config"

func SendConfirmation(to, code string) error {
	return Send(
		to,
		"Register for Word Catching Journal - Confirmation Code",
		`Your confirmation code is: `+code,
	)
}

func SendResetPassword(to, code string) error {
	url := config.Env.ClientURL + "/reset-password/" + code
	return Send(
		to,
		"Reset Password - Confirmation Code",
		`Reset your password at: `+url,
	)
}
