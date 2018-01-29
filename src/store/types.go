package store

type Users struct {
	Id   string `json:id`
	From string `json:"from"`
	To   string `json:"to"`
	Data string `json:"data"`
}
