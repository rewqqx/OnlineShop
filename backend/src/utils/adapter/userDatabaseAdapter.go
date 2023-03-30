package adapter

import "backend/src/utils"

type UserDatabase struct {
	database *utils.DBConnect
}

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
	err = adapter.database.Connection.Get(user, "SELECT * FROM online_shop.users WHERE id=$1", id)

	return
}

func (adapter *UserDatabase) CreateUser(user *User) (token AuthToken, err error) {
	_, err = adapter.database.Connection.Exec("INSERT INTO online_shop.users (user_name, user_surname,user_patronymic, phone, birthdate, password_hash, mail, role_id, token) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", user.Name, user.Surname, user.Patronymic, user.Phone, user.Birthdate, utils.HashPassword(user.Password), user.Mail, user.RoleId, user.Token)
	return adapter.AuthUser(AuthData{Mail: user.Mail, Password: user.Password})
}

func (adapter *UserDatabase) CheckToken(token AuthToken) (ok bool, err error) {
	compareToken := AuthToken{}
	err = adapter.database.Connection.Get(&compareToken, "SELECT token FROM online_shop.users WHERE id=$1", token.ID)
	ok = compareToken.Token == token.Token
	return ok, nil
}

func (adapter *UserDatabase) AuthUser(data AuthData) (token AuthToken, err error) {
	err = adapter.database.Connection.Get(&token, "SELECT id, token FROM online_shop.users WHERE mail=$1 AND password_hash=$2", data.Mail, utils.HashPassword(data.Password))

	if err != nil {
		return
	}

	token.Token = utils.GenerateToken(32)

	_, err = adapter.database.Connection.Exec("UPDATE online_shop.users SET token = $1 WHERE id = $2", token.Token, token.ID)

	return
}
