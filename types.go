package logRush

type Log struct {
	Message   string `json:"message"`
	TimeStamp int    `json:"timestamp"`
	Stream    string `json:"stream"`
}

type LogStream struct {
	ID    string `json:"id"`
	Alias string `json:"alias"`
}
