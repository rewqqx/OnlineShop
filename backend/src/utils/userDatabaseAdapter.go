package utils

type UserDatabase struct {
	database *DBConnect
}

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Phone      string `json:"phone"`
	Birthdate  string `json:"birthdate"`
	Password   string `json:"password_hash"`
	Mail       string `json:"mail"`
	RoleId     int    `json:"role_id"`
}

type AuthToken struct {
	ID    int
	Token string `json:"token"`
}

func (adapter *UserDatabase) getUser(id int) (user User, err error) {
	user = User{}
	err = adapter.database.Connection.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	return
}

func (adapter *UserDatabase) createUser(user User) (id int64, err error) {
	res := adapter.database.Connection.MustExec("INSERT INTO users (user_name, user_surname,user_patronymic, phone, birthdate, password_hash, mail, role_id) VALUES ($1, $2, $3, $4, $5, $6, $7, #8)", user.Name, user.Surname, user.Patronymic, user.Phone, user.Birthdate, user.Password, user.Mail, user.RoleId)
	id, err = res.LastInsertId()
	return
}

func (adapter *UserDatabase) checkToken(token AuthToken) (ok bool, err error) {
	compareToken := AuthToken{}
	err = adapter.database.Connection.Get(&compareToken, "SELECT token FROM users WHERE id=$1", token.ID)

	if err != nil {
		return false, err
	}

	ok = compareToken.Token == token.Token
	return ok, nil
}
