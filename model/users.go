package model

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
)

type (
	User struct {
		Login     string
		Password  string
		Name      string
		SurName   string
		Age       int
		Sex       string
		City      string
		Interests string
	}

	UserView struct {
		name    string
		surname string
		age     int
		sex     int
	}
)

const (
	passwordSalt = "Z#$@df9gfd423"
)

func GetUsers(limit int) ([]UserView, error) {
	var list = make([]UserView, limit)
	query, err := db.Prepare("SELECT name, surname FROM `users` LIMIT ?")
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
		err := rows.Scan(&p.name, &p.surname)
		if err != nil {
			return nil, err
		}
		list[n] = p
		n += 1
	}
	return list, nil
}

func GetUser(id int) (*UserView, error) {
	query, err := db.Prepare("SELECT name, surname FROM `users` WHERE id = ?")

	if err != nil {
		return nil, err
	}
	defer query.Close()
	user := new(UserView)
	err = query.QueryRow(id).Scan(&user.name, &user.surname)
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
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?), ",
		user.Login, getPasswordHash(user.Password), user.Name, user.SurName, user.City, user.Age, user.Sex,
		user.Interests)

	if err != nil {
		return err
	}

	return nil
}

func getPasswordHash(password string) string {
	data := []byte(password + passwordSalt)
	return fmt.Sprintf("%x", sha1.Sum(data))
}
