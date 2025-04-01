package database

import "database/sql"

func DbConnectionn() (*sql.DB, error) {
	connectionString := "root:senhas@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, err
}
