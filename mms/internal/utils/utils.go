package utils

import "regexp"

func ValidatePassword(password string) bool {
	passwordRegex := `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d).{8,}$`
	re := regexp.MustCompile(passwordRegex)
	return re.MatchString(password)
}
