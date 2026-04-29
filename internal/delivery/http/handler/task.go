package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MaksimCpp/TaskManager/internal/delivery/http/middleware"
	"github.com/MaksimCpp/TaskManager/internal/usecase/task"
)

type TaskHandler struct {
	createTaskUseCase task.CreateTaskUseCase
	deleteTaskUseCase task.DeleteTaskUseCase
	getTasksUseCase task.GetTasksUseCase
}

func NewTaskHandler(
	createTaskUseCase task.CreateTaskUseCase,
	deleteTaskUseCase task.DeleteTaskUseCase,
	getTasksUseCase task.GetTasksUseCase,
) *TaskHandler {
	return &TaskHandler{
		createTaskUseCase: createTaskUseCase,
		deleteTaskUseCase: deleteTaskUseCase,
		getTasksUseCase: getTasksUseCase,
	}
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	var request struct {
		Title string `json:"title"`
		Description string `json:"description"`
		Completed bool `json:"completed"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request.", http.StatusBadRequest)
		return
	}

	var input task.CreateTaskInput = task.CreateTaskInput{
		Title: request.Title,
		Description: request.Description,
		Completed: request.Completed,
		UserID: userID,
	}
	err = h.createTaskUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *TaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}
	var request struct {
		TaskID string `json:"task_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request.", http.StatusBadRequest)
		return
	}

	var input task.DeleteTaskInput = task.DeleteTaskInput{
		TaskID: request.TaskID,
		UserID: userID,
	}
	err = h.deleteTaskUseCase.Execute(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *TaskHandler) Get(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "Unauthorized.", http.StatusUnauthorized)
		return
	}

	tasks, err := h.getTasksUseCase.Execute(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
