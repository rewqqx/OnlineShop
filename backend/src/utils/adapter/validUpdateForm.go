package adapter

import (
	"errors"
	"strings"
)

func UpdateUserValid(user *User) (err error) {
	if len(user.Name) < 4 {
		return errors.New("name must contains more than 3 symbols")
	} else if len(user.Surname) < 4 {
		return errors.New("surname must contains more than 3 symbols")
	} else if len(user.Phone) < 6 || !strings.Contains(user.Phone, "+") {
		return errors.New("phone must contains more than 5 symbols or have +")
	} else if len(user.Mail) < 4 || !strings.Contains(user.Mail, "@") || !strings.Contains(user.Mail, ".") {
		return errors.New("mail must contains more than 3 symbols or have @ or .*")
	}

	return nil
}

func UpdateUserWithPasswordValid(user *ChangePassword) (err error) {
	if len(user.Password) < 8 {
		return errors.New("password must contains more than 7 symbols")
	}

	if user.Password != user.PasswordConfirmation {
		return errors.New("passwords must match")
	}

	return nil
}

func IsPasswordChangeRequest(updatePassword *ChangePassword) bool {
	if updatePassword.Password != "" && updatePassword.PasswordConfirmation != "" {
		return true
	}

	return false
}

func IsUpdateDataUserChangeRequest(user *User) bool {
	if user.Name == "" {
		return true
	} else if user.Surname == "" {
		return true
	} else if user.Mail == "" {
		return true
	} else if user.Phone == "" {
		return true
	}

	return false
}
