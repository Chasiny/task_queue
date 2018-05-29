package task_manager

type Tasks []*Task

func (t Tasks) Len() int {
	return len(t)
}
func (t Tasks) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
	t[i].index = i
	t[j].index = j
}
func (t Tasks) Less(i, j int) bool {
	return t[i].NextTime < t[j].NextTime
}
func (t *Tasks) Push(x interface{}) {
	n := len(*t)
	item := x.(*Task)
	item.index = n
	*t = append(*t, item)
}
func (js *Tasks) Pop() (interface{}) {
	old := *js
	n := len(old)
	x := old[n-1]
	x.index = -1
	*js = old[0 : n-1]
	return x
}

type Task struct {
	ID       string
	NextTime int64
	Interval int64
	Cmd      string
	Args     []string

	IsActive bool

	index    int
}
