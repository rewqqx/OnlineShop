package validation

import (
	"errors"
	"strings"
)

func UpdateUserValid(name, surname, phone, mail string) (err error) {
	if len(name) < 4 || len(surname) < 4 || len(phone) < 6 || !strings.Contains(phone, "+") || len(mail) < 4 || !strings.Contains(mail, "@") || !strings.Contains(mail, ".") {
		return errors.New("invalid data in fields")
	}

	return nil
}

func UpdateUserWithPasswordValid(password, passwordConfirmation string) (err error) {
	if len(password) < 8 {
		return errors.New("password must contains more than 7 symbols")
	}

	if password != passwordConfirmation {
		return errors.New("passwords must match")
	}

	return nil
}

func IsPasswordChangeRequest(password, passwordConfirmation string) bool {
	if password != "" && passwordConfirmation != "" {
		return true
	}

	return false
}

func IsUpdateDataUserChangeRequest(name, surname, phone, mail string) bool {
	if name == "" || surname == "" || mail == "" || phone == "" {
		return true
	}

	return false
}
