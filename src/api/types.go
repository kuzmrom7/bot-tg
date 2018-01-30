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
	Time0  string `json:"time0"`
	Number string `json:"number"`
}
