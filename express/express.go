package express

import (
	"fmt"
	"net/http"
)

type Express struct{}

func (e Express) Get(route string, callback func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(route, callback)
}

func (e Express) Listen(port string) {
	fmt.Println("Server listening on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}

type BodyParser struct {
	StatusCode int
	Status     string
	Message    string
}

func (b BodyParser) Json(res http.ResponseWriter, statusCode int, status string, message string) {
	body := BodyParser{statusCode, status, message}
	obj := "{ statusCode: " + string(body.StatusCode) + ", status: " + body.Status + ", message: " + body.Message + " }"

	fmt.Fprint(res, obj)
}
