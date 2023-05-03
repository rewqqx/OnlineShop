package validation

import (
	"errors"
	"strings"
)

func UpdateUserValid(name, surname, phone, mail string) (err error) {
	if len(name) < 4 {
		return errors.New("name must contains more than 3 symbols")
	}

	if len(surname) < 4 {
		return errors.New("surname must contains more than 3 symbols")
	}

	if len(phone) < 6 {
		return errors.New("phone must contains more than 5 symbols")
	}

	if !strings.Contains(phone, "+") {
		return errors.New("phone must contains +")
	}

	if len(mail) < 4 {
		return errors.New("mail must contains more than 3 symbols")
	}

	if !strings.Contains(mail, "@") {
		return errors.New("mail must contains @")
	}

	if !strings.Contains(mail, ".") {
		return errors.New("mail must contains .*")
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
	if name == "" {
		return true
	}

	if surname == "" {
		return true
	}

	if mail == "" {
		return true
	}

	if phone == "" {
		return true
	}

	return false
}
