package adapter

import (
	"backend/src/utils"
	"fmt"
)

type UserDatabase struct {
	database *utils.DBConnect
}

const USER_TABLE_NAME = "users"

type User struct {
	ID         int              `json:"id" db:"id"`
	Name       string           `json:"user_name" db:"user_name"`
	Surname    string           `json:"user_surname" db:"user_surname"`
	Patronymic string           `json:"user_patronymic" db:"user_patronymic"`
	Phone      string           `json:"phone" db:"phone"`
	Birthdate  *utils.Timestamp `json:"birthdate" db:"birthdate"`
	Password   string           `json:"password_hash" db:"password_hash"`
	Mail       string           `json:"mail" db:"mail"`
	RoleId     int              `json:"role_id" db:"role_id"`
	Token      string           `json:"token" db:"token"`
}

type AuthToken struct {
	ID    int
	Token string `json:"token" db:"token"`
}

type AuthData struct {
	Mail     string `json:"mail" db:"mail"`
	Password string `json:"password" db:"password_hash"`
}

func CreateUserDatabaseAdapter(database *utils.DBConnect) *UserDatabase {
	adapter := &UserDatabase{database: database}
	return adapter
}

func (adapter *UserDatabase) GetUser(id int) (user *User, err error) {
	user = &User{}
	err = adapter.database.Connection.Get(user, fmt.Sprintf("SELECT * FROM online_shop.%v WHERE id=$1", USER_TABLE_NAME), id)

	return
}

func (adapter *ItemDatabase) DeleteUser(id int) (err error) {
	_, err = adapter.database.Connection.Exec(fmt.Sprintf("DELETE FROM online_shop.%v WHERE id=$1", USER_TABLE_NAME), id)
	return
}

func (adapter *UserDatabase) CreateUser(user *User) (token AuthToken, err error) {
	_, err = adapter.database.Connection.Exec(fmt.Sprintf("INSERT INTO online_shop.%v (user_name, user_surname,user_patronymic, phone, birthdate, password_hash, mail, role_id, token) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", USER_TABLE_NAME), user.Name, user.Surname, user.Patronymic, user.Phone, user.Birthdate, utils.HashPassword(user.Password), user.Mail, user.RoleId, user.Token)
	return adapter.AuthUser(AuthData{Mail: user.Mail, Password: user.Password})
}

func (adapter *UserDatabase) CheckToken(token AuthToken) (ok bool, err error) {
	compareToken := AuthToken{}
	err = adapter.database.Connection.Get(&compareToken, fmt.Sprintf("SELECT token FROM online_shop.%v WHERE id=$1", USER_TABLE_NAME), token.ID)
	ok = compareToken.Token == token.Token
	return ok, nil
}

func (adapter *UserDatabase) AuthUser(data AuthData) (token AuthToken, err error) {
	err = adapter.database.Connection.Get(&token, fmt.Sprintf("SELECT id, token FROM online_shop.%v WHERE mail=$1 AND password_hash=$2", USER_TABLE_NAME), data.Mail, utils.HashPassword(data.Password))

	if err != nil {
		return
	}

	token.Token = utils.GenerateToken(32)

	_, err = adapter.database.Connection.Exec(fmt.Sprintf("UPDATE online_shop.%v SET token = $1 WHERE id = $2", USER_TABLE_NAME), token.Token, token.ID)

	return
}