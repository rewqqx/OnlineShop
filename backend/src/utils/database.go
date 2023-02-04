package utils

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

var sshTunnel SSH

type DBConnect struct {
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`

	Connection *sqlx.DB
}

func InitConnection(tunnel SSH) {
	sshTunnel = tunnel
}

func (client *DBConnect) Open() error {

	driver := "postgres"

	if sshTunnel.client != nil {
		driver = "postgres+ssh"
		sql.Register(driver, &ViaSSHDialer{sshTunnel.client})
	}

	db, err := sqlx.Open(driver, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", client.User, client.Password, client.Ip, client.Port, client.Database))
	if err != nil {
		return err
	}

	err = db.Ping()

	if err != nil {
		return err
	}

	client.Connection = db
	return nil
}

func (client *DBConnect) Close() {
	client.Connection.Close()
}

func (client *DBConnect) GetTables(schema string) ([]string, error) {
	var res []string
	rows, err := client.Connection.Query("SELECT table_name FROM information_schema.tables WHERE table_schema = $1;", schema)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var el string
		err = rows.Scan(&el)

		if err != nil {
			fmt.Println(err)
		}

		res = append(res, el)
	}

	return res, nil
}
func (client *DBConnect) GetSchemas() ([]string, error) {
	var res []string
	rows, err := client.Connection.Query("SELECT schema_name FROM information_schema.schemata;")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var el string
		err = rows.Scan(&el)

		if err != nil {
			fmt.Println(err)
		}

		res = append(res, el)
	}

	return res, nil
}
