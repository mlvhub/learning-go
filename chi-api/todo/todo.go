package todo

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Todo struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{todoId}", GetTodo)
	router.Delete("/{todoId}", DeleteTodo)
	router.Post("/", CreateTodo)
	router.Get("/", GetAllTodos)
	return router
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	todoId := chi.URLParam(r, "todoId")
	todos := Todo{
		Slug:  todoId,
		Title: "Hello world",
		Body:  "Hello world from planet Earth",
	}
	render.JSON(w, r, todos)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted TODO successfully"
	render.JSON(w, r, response)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created TODO successfully"
	render.JSON(w, r, response)
}
func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{
		{
			Slug:  "slug",
			Title: "Hello world",
			Body:  "Hello world from planet Earth",
		},
	}
	render.JSON(w, r, todos)
}
