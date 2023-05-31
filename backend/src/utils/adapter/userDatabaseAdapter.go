package adapter

import (
	"backend/src/utils/crypto"
	"backend/src/utils/database"
	"backend/src/utils/timestamp"
	"backend/src/validation"
	"database/sql"
	"fmt"
)

type UserDatabase struct {
	database *database.DBConnect
}

const USER_TABLE_NAME = "users"

type User struct {
	ID         int                  `json:"id" db:"id"`
	Name       string               `json:"user_name" db:"user_name"`
	Surname    string               `json:"user_surname" db:"user_surname"`
	Patronymic string               `json:"user_patronymic" db:"user_patronymic"`
	Phone      string               `json:"phone" db:"phone"`
	Birthdate  *timestamp.Timestamp `json:"birthdate" db:"birthdate"`
	Password   string               `json:"password_hash" db:"password_hash"`
	Mail       string               `json:"mail" db:"mail"`
	RoleId     int                  `json:"role_id" db:"role_id"`
	Token      string               `json:"token" db:"token"`
	Sex        int                  `json:"sex" db:"sex"`
}

type ChangePassword struct {
	Password             string `json:"password_hash_change" `
	PasswordConfirmation string `json:"password_confirmation_hash"`
}

type AuthToken struct {
	ID    int
	Token string `json:"token" db:"token"`
}

type AuthData struct {
	Mail     string `json:"mail" db:"mail"`
	Password string `json:"password" db:"password_hash"`
}

func CreateUserDatabaseAdapter(database *database.DBConnect) *UserDatabase {
	adapter := &UserDatabase{database: database}
	return adapter
}

func (adapter *UserDatabase) GetUser(id int) (user *User, err error) {
	user = &User{}
	err = adapter.database.Connection.Get(user, fmt.Sprintf("SELECT * FROM online_shop.%v WHERE id=$1", USER_TABLE_NAME), id)

	user.Password = ""

	return
}

func (adapter *UserDatabase) GetUsers() (users []*User, err error) {
	rows, err := adapter.database.Connection.Query(fmt.Sprintf("SELECT * FROM online_shop.%v", USER_TABLE_NAME))
	if err != nil {
		return nil, err
	}

	return parseUsersFromRows(rows)
}

func parseUsersFromRows(rows *sql.Rows) (users []*User, err error) {
	for rows.Next() {
		var id int
		var name string
		var surname string
		var patronymic string
		var phone string
		var birthdate *timestamp.Timestamp
		var sex int
		var password string
		var mail string
		var roleId int
		var token string

		err = rows.Scan(&id, &name, &surname, &patronymic, &phone, &birthdate, &sex, &password, &mail, &roleId, &token)

		if err != nil {
			return
		}

		user := &User{id, name, surname, patronymic, phone, birthdate, password, mail, roleId, token, sex}
		users = append(users, user)
	}

	return
}

func (adapter *ItemDatabase) DeleteUser(id int) (err error) {
	_, err = adapter.database.Connection.Exec(fmt.Sprintf("DELETE FROM online_shop.%v WHERE id=$1", USER_TABLE_NAME), id)
	return
}

func (adapter *UserDatabase) CreateUser(user *User) (token AuthToken, err error) {
	_, err = adapter.database.Connection.Exec(fmt.Sprintf("INSERT INTO online_shop.%v (user_name, user_surname,user_patronymic, phone, birthdate, sex, password_hash, mail, role_id, token) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", USER_TABLE_NAME), user.Name, user.Surname, user.Patronymic, user.Phone, user.Birthdate, user.Sex, crypto.HashPassword(user.Password), user.Mail, user.RoleId, user.Token)
	return adapter.AuthUser(AuthData{Mail: user.Mail, Password: user.Password})
}

func (adapter *UserDatabase) UpdateUser(user *User, token AuthToken) (r AuthToken, err error) {
	err = validation.UpdateUserValid(user.Name, user.Surname, user.Phone, user.Mail)
	if err != nil {
		return
	}

	_, err = adapter.database.Connection.Exec(fmt.Sprintf("UPDATE online_shop.%v SET user_name = $1, user_surname = $2, phone = $3, mail = $4 WHERE id = $5", USER_TABLE_NAME), user.Name, user.Surname, user.Phone, user.Mail, token.ID)

	return token, nil
}

func (adapter *UserDatabase) UpdateUserWithPassword(user *ChangePassword, token AuthToken) (r AuthToken, err error) {
	err = validation.UpdateUserWithPasswordValid(user.Password, user.PasswordConfirmation)
	if err != nil {
		return token, err
	}

	token.Token = crypto.GenerateToken(32)

	_, err = adapter.database.Connection.Exec(fmt.Sprintf("UPDATE online_shop.%v SET password_hash = $1, token = $2 WHERE id = $3", USER_TABLE_NAME), crypto.HashPassword(user.Password), token.Token, token.ID)

	return token, err
}

func (adapter *UserDatabase) UpdatePassword(user *User) (token AuthToken, err error) {
	_, err = adapter.database.Connection.Exec(fmt.Sprintf("UPDATE online_shop.%v SET password_hash = $1 WHERE id = $2", USER_TABLE_NAME), crypto.HashPassword(user.Password), user.ID)
	return adapter.AuthUser(AuthData{Mail: user.Mail, Password: user.Password})
}

func (adapter *UserDatabase) CheckToken(token AuthToken) (ok bool, err error) {
	compareToken := AuthToken{}
	err = adapter.database.Connection.Get(&compareToken, fmt.Sprintf("SELECT token FROM online_shop.%v WHERE id=$1", USER_TABLE_NAME), token.ID)
	ok = compareToken.Token == token.Token
	return ok, nil
}

func (adapter *UserDatabase) CheckTokenAndRole(token AuthToken) (ok bool, err error) {
	var role_id int64
	err = adapter.database.Connection.Get(&role_id, fmt.Sprintf("SELECT role_id FROM online_shop.%v WHERE token=$1", USER_TABLE_NAME), token.Token)
	ok = role_id == 1
	return true, nil
}

func (adapter *UserDatabase) AuthUser(data AuthData) (token AuthToken, err error) {
	err = adapter.database.Connection.Get(&token, fmt.Sprintf("SELECT id, token FROM online_shop.%v WHERE mail=$1 AND password_hash=$2", USER_TABLE_NAME), data.Mail, crypto.HashPassword(data.Password))

	if err != nil {
		return
	}

	token.Token = crypto.GenerateToken(32)

	_, err = adapter.database.Connection.Exec(fmt.Sprintf("UPDATE online_shop.%v SET token = $1 WHERE id = $2", USER_TABLE_NAME), token.Token, token.ID)

	return
}
