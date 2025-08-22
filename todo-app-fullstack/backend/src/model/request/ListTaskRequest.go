package request

type ListTask struct {
	NameTask string `json:"name_task" binding:"required,min=3"`
	Status string `json:"status" binding:"required"`
	TaskId uint `json:"task_id" binding:"required"`
	Model
}
