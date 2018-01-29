package store

import "log"

func CreateNote(id int64)  {
	_, err := db.Exec("INSERT INTO users (id) VALUES ($1)", id)
	if err != nil{
		log.Println()
	}
	log.Println("--store--->create note!")
}
