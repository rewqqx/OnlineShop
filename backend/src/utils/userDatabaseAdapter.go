package utils

import "time"

type UserDatabase struct {
	database *DBConnect
}

type User struct {
	ID         int       `json:"id" db:"id"`
	Name       string    `json:"user_name" db:"user_name"`
	Surname    string    `json:"user_surname" db:"user_surname"`
	Patronymic string    `json:"user_patronymic" db:"user_patronymic"`
	Phone      string    `json:"phone" db:"phone"`
	Birthdate  time.Time `json:"birthdate" db:"birthdate"`
	Password   string    `json:"password_hash" db:"password_hash"`
	Mail       string    `json:"mail" db:"mail"`
	RoleId     int       `json:"role_id" db:"role_id"`
	Token      string    `json:"token" db:"token"`
}

type AuthToken struct {
	ID    int
	Token string `json:"token"`
}

func (adapter *UserDatabase) getUser(id int) (user User, err error) {
	user = User{}
	res, err := adapter.database.Connection.Query("SELECT * FROM online_shop.users WHERE id=$1", id)

	if res.Next() {
		if err = res.Scan(&user); err != nil {
			return user, err
		}
	}

	return
}

func (adapter *UserDatabase) createUser(user User) (id int64, err error) {
	res, err := adapter.database.Connection.Exec("INSERT INTO online_shop.users (user_name, user_surname,user_patronymic, phone, birthdate, password_hash, mail, role_id) VALUES ($1, $2, $3, $4, $5, $6, $7, #8)", user.Name, user.Surname, user.Patronymic, user.Phone, user.Birthdate, user.Password, user.Mail, user.RoleId)
	id, err = res.LastInsertId()
	return
}

func (adapter *UserDatabase) checkToken(token AuthToken) (ok bool, err error) {
	compareToken := AuthToken{}
	res, err := adapter.database.Connection.Query("SELECT token FROM online_shop.users WHERE id=$1", token.ID)

	if err != nil {
		return false, err
	}

	if res.Next() {
		if err = res.Scan(&compareToken); err != nil {
			return false, err
		}
	}

	ok = compareToken.Token == token.Token
	return ok, nil
}
