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
		From     string
		To       string
		Date     string
		TrainNum string
		Time     string
	)

	for _, train := range data.Tp {
		From = train.From

		To = train.Where

		Date = train.Date

		for _, list := range train.List {

			TrainNum = list.Number

			Time = list.Time0

			trainsss = trainsss + Number + " Номер поезда *" + TrainNum + "*\n" + Clock + " Время отправления *" + Time + "*\n\n"
		}

		message = message + TraintTo + "*" + From + "\t -> " + To + "\t" + Date + "*" + "\n\n" + trainsss + "\n\n"
	}
	log.Println("data")
	return message
}
