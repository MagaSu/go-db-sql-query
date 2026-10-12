package main

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

type Client struct {
	ID       int
	FIO      string
	Login    string
	Birthday string
	Email    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Client.
// Теперь, если передать объект Client в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (c Client) String() string {
	return fmt.Sprintf("ID: %d FIO: %s Login: %s Birthday: %s Email: %s",
		c.ID, c.FIO, c.Login, c.Birthday, c.Email)
}

func main() {
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// добавление нового клиента
	newClient := Client{
		FIO:      "Сулейманов Мухаммад Сахибович", // укажите ФИО
		Login:    "Harkaro",                       // укажите логин
		Birthday: "20010703",                      // укажите день рождения
		Email:    "magamag268@gmail.com",          // укажите почту
	}

	id, err := insertClient(db, newClient)
	if err != nil {
		fmt.Println(err)
		return
	}

	// получение клиента по идентификатору и вывод на консоль
	client, err := selectClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client)

	// обновление логина клиента
	newLogin := "Cheberlo" // укажите новый логин
	err = updateClientLogin(db, newLogin, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	// получение клиента по идентификатору и вывод на консоль
	client, err = selectClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(client)

	// удаление клиента
	err = deleteClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}

	// получение клиента по идентификатору и вывод на консоль
	_, err = selectClient(db, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func insertClient(db *sql.DB, client Client) (int64, error) {
	query := `INSERT INTO clients (fio, login, birthday, email) VALUES (:fio, :login, :birthday, :email)`
	req, err := db.Exec(query,
		sql.Named("fio", client.FIO),
		sql.Named("login", client.Login),
		sql.Named("birthday", client.Birthday),
		sql.Named("email", client.Email),
	)
	if err != nil {
		return 0, err
	}
	id, err := req.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func updateClientLogin(db *sql.DB, login string, id int64) error {
	query := `UPDATE clients SET login = :login WHERE id = :id`
	_, err := db.Exec(query,
		sql.Named("login", login),
		sql.Named("id", id),
	)
	if err != nil {
		return err
	}

	return nil
}

func deleteClient(db *sql.DB, id int64) error {
	query := `DELETE FROM clients WHERE id = :id`

	_, err := db.Exec(query, sql.Named("id", id))
	if err != nil {
		return err
	}

	return nil
}

func selectClient(db *sql.DB, id int64) (Client, error) {
	client := Client{}

	row := db.QueryRow("SELECT id, fio, login, birthday, email FROM clients WHERE id = :id", sql.Named("id", id))
	err := row.Scan(&client.ID, &client.FIO, &client.Login, &client.Birthday, &client.Email)

	return client, err
}
