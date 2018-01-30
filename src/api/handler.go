package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)

type Resp struct {
	Result    string `json:"result"`
	Timestamp string `json:"timestamp"`
	Rid       string `json:"rid"`
}

type Trains struct {
	Result  string `json:"result"`
	Message string `json:"message"`
	Tp      Train  `json:"tp"`
}

type Train struct {
	From  string `json:"from"`
	Where string `json:"where"`
}

var myClient = &http.Client{}
var cooks []*http.Cookie

func main() {
	foo1 := new(Resp)
	getRID("https://pass.rzd.ru/timetable/public/ru?STRUCTURE_ID=735&layer_id=5371&dir=0&tfl=3&checkSeats=1&code0=2041603&dt0=30.01.2018&code1=2000000", foo1)
	log.Println("get RID", foo1)

	foo2 := new(Trains)
	getTrains("http://pass.rzd.ru/timetable/public/ru?STRUCTURE_ID=735&layer_id=5371&rid="+foo1.Rid, foo2)
	log.Println("get Trains", foo2)

}

func getRID(url string, target interface{}) (error) {
	fmt.Println("URL идет", url)
	r, err := myClient.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	cooks = r.Cookies()

	return json.NewDecoder(r.Body).Decode(target)
}

func getTrains(url string, target interface{}) (error) {
	fmt.Println("URL идет", url)

	req, _ := http.NewRequest("GET", url, nil)

	for _, cookie := range cooks {
		req.AddCookie(cookie)
	}

	var client = &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}
