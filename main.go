package main

import (
	"fmt"
	"net/http"

	"github.com/musjib999/expressServer/express"
)

func main() {
	app := express.Express{}
	bodyParser := express.BodyParser{}

	app.Get("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Welcome to our express app root endpont!!!")
	})

	app.Get("/user", func(res http.ResponseWriter, req *http.Request) {
		bodyParser.Json(res, 200, "Success", "All Users Fetched ")
	})

	app.Listen(":3000")
}
