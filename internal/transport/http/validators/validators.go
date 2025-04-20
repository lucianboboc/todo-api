package validators

import "regexp"

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

func ValidateEmail(email string) bool {
	return EmailRX.MatchString(email)
}

func ValidateNotEmpty(value string) bool {
	return value != ""
}

func ValidateUserId(id int) bool {
	return id > 0
}

func ValidatePassword(password string) bool {
	if password == "" {
		return false
	}
	if len(password) < 6 {
		return false
	}
	return true
}
