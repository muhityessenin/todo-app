package todo

type Task struct {
	ID       string `json:"id"`
	Title    string `json:"title" binding:"required"`
	ActiveAt string `json:"activeAt" time_format:"YYYY-MM-DD" binding:"required"`
	Status   string `json:"status"`
}
