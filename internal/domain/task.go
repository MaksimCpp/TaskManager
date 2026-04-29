package domain


type Task struct {
	ID string
	Title string
	Description string
	Completed bool
    UserID string
}

func NewTask(
	id string,
	title string,
	description string,
	completed bool,
    userID string,
) *Task {
	return &Task{
		ID: id,
		Title: title,
		Description: description,
		Completed: completed,
		UserID: userID,
	}
}
