package model

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

type (
	User struct {
		Login     string
		Password  string
		Name      string
		Surname   string
		Age       int
		Sex       string
		City      string
		Interests string
	}

	UserView struct {
		Id        int
		Name      string
		Surname   string
		Age       int
		Sex       string
		City      string
		Interests string
	}

	DuplicateRecordError struct {
		Message string
	}
)

const (
	passwordSalt = "Z#$@df9gfd423"
)

func (e *DuplicateRecordError) Error() string {
	return fmt.Sprintf("Error %s", e.Message)
}

func GetUsers(limit int) ([]UserView, error) {
	var list []UserView
	query, err := db.Prepare("SELECT name, surname, age, sex, city FROM `users` LIMIT ?")
	if err != nil {
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query(limit)
	if err != nil {
		return nil, err
	}

	n := 0
	for rows.Next() {
		var p UserView
		err := rows.Scan(&p.Name, &p.Surname, &p.Age, &p.Sex, &p.City)
		if err != nil {
			return nil, err
		}
		list = append(list, p)
		n += 1
	}
	return list, nil
}

func GetUser(id int) (*UserView, error) {
	query, err := db.Prepare("SELECT id, name, surname, age FROM `users` WHERE id = ?")

	if err != nil {
		return nil, err
	}
	defer query.Close()
	user := new(UserView)
	err = query.QueryRow(id).Scan(&user.Id, &user.Name, &user.Surname, &user.Age, &user.City)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func AuthUser(login string, password string) (bool, error) {
	query, err := db.Prepare("SELECT name, login FROM `users` WHERE login = ? AND `password` = ?")
	if err != nil {
		return false, err
	}

	defer query.Close()

	var user User

	err = query.QueryRow(login, getPasswordHash(password)).Scan(&user.Name, &user.Login)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func AddUser(user User) error {
	_, err := db.Exec("INSERT INTO users (login, `password`, name, surname, city, age, sex, interests) "+
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?)",
		user.Login, getPasswordHash(user.Password), user.Name, user.Surname, user.City, user.Age, user.Sex,
		user.Interests)

	if err != nil {
		if myError, ok := err.(*mysql.MySQLError); ok {
			if myError.Number == 1062 {
				dupError := DuplicateRecordError{Message: fmt.Sprintf("user with email %s already exists", user.Login)}
				return &dupError
			}
		}
		return err
	}

	return nil
}

func getPasswordHash(password string) string {
	data := []byte(password + passwordSalt)
	return fmt.Sprintf("%x", sha1.Sum(data))
}
