package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Application struct {
	router *mux.Router
}

type Response struct {
	Name string `json:"name"`
	Mail string `json:"mail"`
}

func New(name string) *Application {
	app := &Application{}
	app.router = mux.NewRouter()
	return app
}

func (a *Application) Run() {
	a.router.HandleFunc("/user/hello", Handler_Hello).Methods("GET")
	http.Handle("/", a.router)
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func Handler_Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world\n")
}
