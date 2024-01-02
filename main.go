package main

import (
	"html/template"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	counter := Counter{}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("index.html")
		data := map[string]int{
			"CounterValue": counter.GetValue(),
		}

		tmpl.ExecuteTemplate(w, "index.html", data)
	})

	router.Post("/increase", func(w http.ResponseWriter, r *http.Request) {
		tmplStr := "<div id=\"counter\">{{.CounterValue}}</div>"
		tmpl := template.Must(template.New("counter").Parse(tmplStr))

		counter.Increase()

		data := map[string]int{
			"CounterValue": counter.GetValue(),
		}

		tmpl.ExecuteTemplate(w, "counter", data)
	})

	router.Post("/decrease", func(w http.ResponseWriter, r *http.Request) {
		tmplStr := "<div id=\"counter\">{{.CounterValue}}</div>"
		tmpl := template.Must(template.New("counter").Parse(tmplStr))

		counter.Decrease()

		data := map[string]int{
			"CounterValue": counter.GetValue(),
		}

		tmpl.ExecuteTemplate(w, "counter", data)
	})

	http.ListenAndServe("localhost:3000", router)
}

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increase() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func (c *Counter) Decrease() {
	c.mu.Lock()
	c.value--
	c.mu.Unlock()
}

func (c *Counter) GetValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.value
}
