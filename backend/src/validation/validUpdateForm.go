package validation

import (
	"errors"
	"strings"
)

func UpdateUserValid(name, surname, phone, mail string) (err error) {
	if len(name) < 4 {
		return errors.New("name must contains more than 3 symbols")
	} else if len(surname) < 4 {
		return errors.New("surname must contains more than 3 symbols")
	} else if len(phone) < 6 || !strings.Contains(phone, "+") {
		return errors.New("phone must contains more than 5 symbols or have +")
	} else if len(mail) < 4 {
		return errors.New("mail must contains more than 3 symbols")
	} else if !strings.Contains(mail, "@") {
		return errors.New("mail must contains @")
	} else if !strings.Contains(mail, ".") {
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
	} else if surname == "" {
		return true
	} else if mail == "" {
		return true
	} else if phone == "" {
		return true
	}

	return false
}
