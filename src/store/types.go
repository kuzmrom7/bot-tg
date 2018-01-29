package store

type Users struct {
	Id   int64  `json:id`
	From string `json:"from"`
	To   string `json:"to"`
	Data string `json:"data"`
}
