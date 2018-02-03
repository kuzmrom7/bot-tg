package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)

type Codes struct {
	Name string `json:"n"`
	Code int    `json:"c"`
}

func main() {
	searchCity("МОСКВА")
}

func searchCity(city string) {
	foo1 := new([]Codes)

	url := fmt.Sprintf("http://www.rzd.ru/suggester?compactMode=y&stationNamePart=%s&lang=ru", city)
	getCode(url, foo1)

	une := []rune(city)
	val := len(une)

	foo2 := make([]Codes, 0)
	for _, code := range *foo1 {
		runes := []rune(code.Name)
		safeSubstring := string(runes[0:val])
		if (safeSubstring == city) {
			foo2 = append(foo2, code)
			log.Println(safeSubstring, code.Code, code.Name)
		}

	}
	for _, i := range foo2 {
		log.Println(" >>>", i.Name)
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
