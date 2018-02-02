package api

type Resp struct {
	Result    string `json:"result"`
	Timestamp string `json:"timestamp"`
	Rid       string `json:"rid"`
}

type Trains struct {
	Result             string `json:"result"`
	TransferSearchMode string `json:"TransferSearchMode"`
	Timestamp          string `json:"timestamp"`
	FlFPKRoundBonus    bool   `json:"flFPKRoundBonus"`
	Tp                 Train  `json:"tp"`
}

type Train []struct {
	State string     `json:"state"`
	From  string     `json:"from"`
	Where string     `json:"where"`
	Date  string     `json:"date"`
	List  ListTrains `json:"list"`
}

type ListTrains []struct {
	Number    string `json:"number"`
	Time0     string `json:"time0"`
	Date0     string `json:"date0"`
	Time1     string `json:"time1"`
	Date1     string `json:"date1"`
	Route0    string `json:"route0"`
	Route1    string `json:"route1"`
	TimeInWay string `json:"timeInWay"`
	Cars      Cars   `json:"cars"`
}

type Cars []struct {

}
