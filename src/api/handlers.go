package api

import (
	"fmt"
	"log"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var myClient = &http.Client{}
var cooks []*http.Cookie

func getRID(url string, target interface{}) (error) {
	log.Println("URL request", url)
	r, err := myClient.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()

	cooks = r.Cookies()
	time.Sleep(1500 * time.Millisecond)
	return json.NewDecoder(r.Body).Decode(target)
}

func getTrains(url string, target interface{}) {
	log.Println("URL request", url)

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

	file, err := os.Create("one.json")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	file.Write(body)

	_ = json.Unmarshal(body, target)

}
