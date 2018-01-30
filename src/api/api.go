package api

import (
	"log"
)

var (
	FromCode string
	ToCode   string
	Date     string
)

func App() {
	foo1 := new(Resp)
	FromCode = "2004001"
	ToCode = "2000000"
	Date = "02.02.2018"

	url := "https://pass.rzd.ru/timetable/public/ru?STRUCTURE_ID=735&layer_id=5371&dir=0&tfl=3&checkSeats=1&code0=" + FromCode + "&dt0=" + Date + "&code1=" + ToCode

	getRID(url, foo1)
	log.Println("get RID", foo1)

	foo2 := new(Trains)
	getTrains("http://pass.rzd.ru/timetable/public/ru?STRUCTURE_ID=735&layer_id=5371&rid="+foo1.Rid, foo2)
	log.Println("get Trains", foo2.Tp)

}
