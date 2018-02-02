package bot

import (
	"github.com/bot-tg/src/store"
	"github.com/bot-tg/src/api"
	"log"
)

type Msg struct {
	message string
}

func AddData(id int64) (Message string) {
	from, to, date := store.GetData(id)
	data := api.App(from, to, date)

	var (
		message  = ""
		trainsss = ""
	)

	var (
		From      string
		To        string
		Date      string
		TrainNum  string
		Time0     string
		Time1     string
		Date1     string
		Date0     string
		Route0    string
		Route1    string
		TimeInWay string
	)

	for _, train := range data.Tp {
		From = train.From

		To = train.Where

		Date = train.Date

		for _, list := range train.List {

			TrainNum = list.Number
			Time0 = list.Time0
			Time1 = list.Time1
			Date0 = list.Date0
			Date1 = list.Date1
			Route0 = list.Route0
			Route1 = list.Route1
			TimeInWay = list.TimeInWay

			msg := handleMsgTrain(TrainNum, Time0, Time1, Date0, Date1, Route0, Route1, TimeInWay)

			trainsss = trainsss + msg
		}

		message = message + TraintTo + "*" + From + "\t - " + To + "\t" + Date + "*" + "\n\n" + trainsss + "\n\n"
	}
	log.Println("data all ready")
	return message
}

func handleMsgTrain(number, time0, time1, date0, date1, route0, route1, timeInway string) string {
	var msg string

	str1 := Train2 + " *" + number + " " + route0 + " -> " + route1 + "* \n"
	str2 := Clock + "Отпраление " + date0 + " " + time0 + "\n"
	str3 := Clock2 + "Прибытие " + date1 + " " + time1 + "\n"
	str4 := Clock3 + "В пути " + timeInway + "\n"
	wavy := Wavy+Wavy+Wavy+Wavy+Wavy+Wavy+Wavy+Wavy+Wavy+Wavy+Wavy

	msg = str1 + str2 + str3 + str4 + wavy+ "\n\n"

	return msg
}
