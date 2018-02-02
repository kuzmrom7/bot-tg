package store

import (
	"log"
	"strings"
)

func CheckUser(id int64) {
	row := db.QueryRow("SELECT id FROM users WHERE id = $1 ", id)

	us := new(Users)

	_ = row.Scan(&us.Id)

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
	log.Println("--STORE---> delete note!")
}

func CreateUser(id int64) {
	_, err := db.Exec("INSERT INTO users (id) VALUES ($1)", id)
	if err != nil {
		log.Println()
	}
	log.Println("--STORE---> create note!")
}

func AddFrom(id int64, message string) {
	result, err := db.Exec("UPDATE users SET fromfrom = $1  WHERE id = $2", strings.ToUpper(message), id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()

	log.Println("--STORE---> create from!")

}
func AddTo(id int64, message string) {
	result, err := db.Exec("UPDATE users SET toto = $1  WHERE id = $2", strings.ToUpper(message), id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()

	log.Println("--STORE---> create to!")
}
func AddDate(id int64, message string) {
	result, err := db.Exec("UPDATE users SET data = $1  WHERE id = $2", strings.ToUpper(message), id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()

	log.Println("--STORE---> create date!")
}

func GetData(id int64) (from, to, date string) {

	row := db.QueryRow("SELECT * FROM users WHERE id = $1 ", id)

	us := new(Users)

	_ = row.Scan(&us.Id, &us.From, &us.To, &us.Data, &us.FromQues, &us.ToQues, &us.DateQues)

	log.Println("--STORE---> get date!")

	return us.From, us.To, us.Data
}

//CHECK MARK IN THE DATABASE

func FromQuestions(id int64) bool {
	row := db.QueryRow("SELECT fromquestions FROM users WHERE id = $1 ", id)

	var fromquestions bool

	_ = row.Scan(&fromquestions)

	log.Println("--STORE---> get fromquestions!", fromquestions)

	return fromquestions
}

func ToQuestions(id int64) bool {
	row := db.QueryRow("SELECT toquestions FROM users WHERE id = $1 ", id)

	var toquestions bool

	_ = row.Scan(&toquestions)

	log.Println("--STORE---> get toquestions!")

	return toquestions
}

func DateQuestions(id int64) bool {
	row := db.QueryRow("SELECT datequestions FROM users WHERE id = $1 ", id)

	var datequestions bool

	_ = row.Scan(&datequestions)

	log.Println("--STORE---> get datequestions!")

	return datequestions
}

//WRITE MARK IN DATABASE

func WriteFromQuestions(id int64, val bool) {

	result, err := db.Exec("UPDATE users SET fromquestions = $1  WHERE id = $2", val, id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()

	log.Println("--STORE---> UPDATED fromquestions for ID = ", id)
}

func WriteToQuestions(id int64, val bool) {

	result, err := db.Exec("UPDATE users SET toquestions = $1  WHERE id = $2", val, id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()

	log.Println("--STORE---> UPDATED toquestions for ID = ", id)

}

func WriteDateQuestions(id int64, val bool) {

	result, err := db.Exec("UPDATE users SET datequestions = $1  WHERE id = $2", val, id)
	if err != nil {
		log.Fatal(err)
	}

	_, err = result.RowsAffected()

	log.Println("--STORE---> UPDATED datequestions for ID = ", id)

}
