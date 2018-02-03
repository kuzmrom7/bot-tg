package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)

type Codes struct {
	N string `json:"n"`
	C int    `json:"c"`
}

func main() {
	searchCity()
}

func searchCity()  {
	foo1 := new([]Codes)

	city := "ОРСК"

	url := fmt.Sprintf("http://www.rzd.ru/suggester?compactMode=y&stationNamePart=%s&lang=ru",city )
	getCode(url, foo1)

	une := []rune(city)
	val := len(une)

	for _, code := range *foo1{
		runes := []rune(code.N)
		safeSubstring := string(runes[0:val])
		if (safeSubstring == city){
			log.Println(safeSubstring, code.C , code.N)
		}

	}
}

func getCode(url string, target interface{}) (error) {
	var myClient = &http.Client{}

	log.Println("--API-CITY---> URL request", url)
	r, err := myClient.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
