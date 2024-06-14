package request

type RequestCreateTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
