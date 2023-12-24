package models

type Task struct {
	Model
	Name   string `json:"name"`
	Point  int    `json:"point"`
	GoalID uint   `json:"-"`
}

type TaskResponse struct {
	Name  string `json:"name"`
	Point int    `json:"point"`
}

func ToTaskResponse(task Task) TaskResponse {
	return TaskResponse{
		Name:  task.Name,
		Point: task.Point,
	}
}

func ConvertTaskIDToTaskResponse(taskID uint, tasks []Task) []TaskResponse {
	var taskResponse []TaskResponse

	for _, task := range tasks {
		if task.ID == taskID {
			taskResponse = append(taskResponse, ToTaskResponse(task))
		}
	}

	return taskResponse
}
