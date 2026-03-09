package main

import (
	"context"
	"net/http"
	"pet_api/internal/api/ToDo"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	http.Handler()
	r.Get("/{id}", ToDo.GetTask(ctx, id))
}
