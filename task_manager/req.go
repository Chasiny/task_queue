package task_manager

type Request struct {
	Ok       bool     `json:"ok"`
	Error    string   `json:"error"`
	ID       string   `json:"id"`
	Cmd      string   `json:"cmd"`
	Args     []string `json:"args"`
	Interval int64    `json:"interval"`
}
