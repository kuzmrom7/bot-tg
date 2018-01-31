package store

import (
	"log"
	"strings"
)

func CheckUser(id int64) {
	row := db.QueryRow("SELECT id FROM users WHERE id = $1 ", id)

	us := new(Users)

	_ = row.Scan(&us.Id)

	log.Println("Это id из бд ", us.Id, " а это просто ", id)

	if (id == us.Id) {
		DeleteNote(id)
	}

	CreateUser(id)
}

func DeleteNote(id int64) {

	result, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	_, err = result.RowsAffected()
	log.Println("--store--->delete note!")
}

func CreateUser(id int64) {
	_, err := db.Exec("INSERT INTO users (id) VALUES ($1)", id)
	if err != nil {
		log.Println()
	}
	log.Println("--store--->create note!")
}

func AddFrom(id int64, message string) {
	result, err := db.Exec("UPDATE users SET fromfrom = $1  WHERE id = $2", strings.ToUpper(message), id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()

	log.Println("--store--->create from!")

}
func AddTo(id int64, message string) {
	result, err := db.Exec("UPDATE users SET toto = $1  WHERE id = $2", strings.ToUpper(message), id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()

	log.Println("--store--->create to!")
}
func AddDate(id int64, message string) {
	result, err := db.Exec("UPDATE users SET data = $1  WHERE id = $2", strings.ToUpper(message), id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()

	log.Println("--store--->create date!")
}

func GetData(id int64) (from, to, date string) {

	row := db.QueryRow("SELECT * FROM users WHERE id = $1 ", id)

	us := new(Users)

	_ = row.Scan(&us.Id, &us.From, &us.To, &us.Data)

	log.Println("--store--->get date!")

	return us.From, us.To, us.Data
}
